package tailf

import (
	"fmt"
	cm "learnGO/logagent/commonconf"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

const (
	//StatusNormal 正常状态
	StatusNormal = 1
	//StatusDelete 所有配置都被删除状态
	StatusDelete = 2
)

//TextMsg 定义消息结构体
type TextMsg struct {
	Msg   string //实际的信息
	Topic string //往kafka哪个topic写
}

//TailObj 将tail组件中的Tail对象与指定的CollectConf的关联对象
type TailObj struct {
	tail   *tail.Tail
	conf   cm.CollectConf
	status int
	//退出channel,会有一个select语句在判断是否退出
	exitChan chan int
}

//TailObjMgr 管理多个TailObj
type TailObjMgr struct {
	tails   []*TailObj
	msgChan chan *TextMsg
	//定义一个锁 保证同时只有一个goroutine更新配置
	lock sync.Mutex
}

//定义一个TailObjMgr对象
var tailObjMgr *TailObjMgr

//GetOneOnLine 拿到一行log
func GetOneOnLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
}

//InitTail 初始化tail组件
func InitTail(conf []cm.CollectConf, chanSize int) (err error) {
	if len(conf) == 0 {
		err = fmt.Errorf("invaild config for log collect,conf:%v", conf)
		return
	}
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}
	for _, v := range conf {
		createNewTask(v)
	}
	return
}

func readFromTail(tailObj *TailObj) {
	for {
		select {
		case msg, ok := <-tailObj.tail.Lines:
			if !ok {
				logs.Warn("tail file close reopen,filename:%s\n", tailObj.tail.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			textMsg := &TextMsg{
				Msg:   msg.Text,
				Topic: tailObj.conf.Topic,
			}
			tailObjMgr.msgChan <- textMsg
		case <-tailObj.exitChan:
			logs.Warn("tail obj will exited,conf:&v", tailObj.conf)
			return
		}
	}

}

//UpdateConfig 根据监听到的变化更新配置
func UpdateConfig(confs []cm.CollectConf) (err error) {
	tailObjMgr.lock.Lock()
	defer tailObjMgr.lock.Unlock()

	for _, oneConf := range confs {
		var isRunning = false
		for _, obj := range tailObjMgr.tails {
			//判断LogPath是否在运行
			//TODO fengjie 2019年11月25日 为什么没有判断Topic呢？是估计没写？
			if oneConf.LogPath == obj.conf.LogPath {
				isRunning = true
				break
			}
		}
		//本条配置未曾修改，进入下一个循环
		if isRunning {
			continue
		}

		createNewTask(oneConf)
	}

	var tailObjs []*TailObj
	for _, obj := range tailObjMgr.tails {
		// 将其TailObj属性设置为删除状态
		// 如果新配置中有此条配置则置为正常状态
		obj.status = StatusDelete
		for _, oneConf := range confs {
			if oneConf.LogPath == obj.conf.LogPath {
				obj.status = StatusNormal
				break
			}
		}
		// 如果是删除了，那么向此对象发送退出channel，此TailObj的readFromTail go routine退出
		if obj.status == StatusDelete {
			obj.exitChan <- 1
			continue
		}
		tailObjs = append(tailObjs, obj)
	}

	tailObjMgr.tails = tailObjs
	return
}
func createNewTask(conf cm.CollectConf) {

	obj := &TailObj{
		conf:     conf,
		exitChan: make(chan int, 1),
	}

	tails, errTail := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if errTail != nil {
		logs.Error("collect filename[%s] failed, err:%v", conf.LogPath, errTail)
		return
	}

	obj.tail = tails
	tailObjMgr.tails = append(tailObjMgr.tails, obj)

	go readFromTail(obj)

}
