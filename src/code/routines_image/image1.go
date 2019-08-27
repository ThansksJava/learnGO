package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/adonovan/gopl.io/ch8/thumbnail"
)

// 串行
func makeThumbnails1(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// 带错的并发 执行太快了应该是，没完成整个函数就返回了,把后台的工作线程都给kill了
// 试试让它睡一会儿
func makeThumbnails2(filenames []string) {
	for idx, f := range filenames {
		go func(file string, idx int) {
			_, err := thumbnail.ImageFile(file)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(idx, "执行完毕")
		}(f, idx)
	}

	time.Sleep(1000000000)
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(file string) {
			_, err := thumbnail.ImageFile(file)
			if err != nil {
				log.Println(err)
			}
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}

//当新的goroutine开始执行字面函数时，
//for循环可能已经更新了f并且开始了另一轮的迭代或者(更有可能的)已经结束了整个循环，
//所以当这些goroutine开始读取f的值时，它们所看到的值已经是slice的最后一个元素了
//到最后可能只生成最后一个图片的缩略图
func makeThumbnails3WithError(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func() {
			_, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
			}
			ch <- struct{}{}
		}()
	}

	for range filenames {
		<-ch
	}
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
func main() {
	images := make([]string, 4)
	images[0] = "E:/imageforgolang/1.jpg"
	images[1] = "E:/imageforgolang/2.jpg"
	images[2] = "E:/imageforgolang/3.jpg"
	images[3] = "E:/imageforgolang/4.jpg"
	// makeThumbnails1(images)
	// makeThumbnails2(images)
	// makeThumbnails3WithError(images)
	ch1 := make(chan string,4)
	for _, s := range images {
		ch1 <- s
	}
	fmt.Println(makeThumbnails6(ch1))
}
