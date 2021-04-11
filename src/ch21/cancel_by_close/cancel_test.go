package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// 获取取消通知
func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

// 发送取消消息
/* func cancel(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
} */

// 通过关闭 Channel 取消
func cancel(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	cancel(cancelChan)
	time.Sleep(time.Second * 1)
}
