package test

import (
	"fmt"
	"testing"
)

func TestSendToNilChannel(t *testing.T) {
	var c chan int
	if c == nil {
		fmt.Println("c is nil!")
	}
	c <- 1 // 实验结果证明，会阻塞
}
