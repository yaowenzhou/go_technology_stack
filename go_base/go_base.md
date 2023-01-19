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

## 2. go格式化时间
相关源码: [time_format_test](./demos/tests/time_format_test.go)

```go
func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02"))                    // 精确到日
	fmt.Println(now.Format("2006-01-02 15:04:05"))           // 精确到秒
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))       // 精确到毫秒
	fmt.Println(now.Format("2006-01-02 15:04:05.000000"))    // 精确到微秒
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000")) // 精确到纳秒
}
```

## 3. 空的select语句将永远阻塞

## 4. os.Exit 会直接让程序退出，不会执行defer