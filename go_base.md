# go 的基础知识

## 1. 向为nil的channel发送数据，会阻塞
相关源码: [send_to_nil_channle_test](./demos/tests/send_to_nil_channle_test.go)
```go
package test

import (
	"fmt"
	"testing"
)

func TestSendToNilChannel(t *testing.T) {
	var c chan int
	if c == nil {
		fmt.Println("c is nil!") // 此处会被打印
	}
	c <- 1 // 实验结果证明，会阻塞
}
```