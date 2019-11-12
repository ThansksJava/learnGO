package tailf

import (
	"fmt"
	cm "logagent/commonconf"
	"time"

	"github.com/hpcloud/tail"
)

//TextMsg 定义消息结构体
type TextMsg struct {
	Msg   string //实际的信息
	Topic string //往kafka哪个topic写
}

//TailObj 将tail组件中的Tail对象与指定的CollectConf的关联对象
type TailObj struct {
	tail *tail.Tail
	conf cm.CollectConf
}

//TailObjMgr 管理多个TailObj
type TailObjMgr struct {
	tails   []*TailObj
	msgChan chan *TextMsg
}

//定义一个TailObjMgr对象
var tailObjMgr *TailObjMgr

//GetOneOnLine 拿到一行log
func GetOneOnLine()(msg *TextMsg){
	msg = <- tailObjMgr.msgChan
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
		tailObj := &TailObj{
			conf: v,
		}
		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,
			Follow:    true,
			MustExist: false,
			Poll:      true,
		})

		if errTail != nil {
			err = errTail
			return
		}
		tailObj.tail = tails

		tailObjMgr.tails = append(tailObjMgr.tails, tailObj)

		go readFromTail(tailObj)
	}
	return
}

func readFromTail(tailObj *TailObj) {
	for {
		msg, ok := <-tailObj.tail.Lines
		if !ok {
			fmt.Printf("tail file close reopen,filename:%s\n", tailObj.tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		textMsg := &TextMsg{
			Msg:   msg.Text,
			Topic: tailObj.conf.Topic,
		}
		tailObjMgr.msgChan <- textMsg
	}

}
