# 本文是go进阶的内容

## 1. error 解析
### 1.1 error的定义:
error 是一个接口类型，包含了一个`Error()`的方法，所有实现了该方法的结构体都可以当做error类型
```go
type error interface {
	Error() string
}
```
#### 1.1.1 errors.New
下面是`Go/src/errors/errors.go`的源码
```go
// because the former will succeed if err wraps an *fs.PathError.
package errors

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```
&emsp;&emsp;New方法通过传入的text创建不同的errorString，这个结构体实现了Error()方法，这样便可以动态的创建不同的错误类型。
&emsp;&emsp;New返回了一个实现error接口的errorString类型的指针。 那为啥不是返回errorString的值类型？
&emsp;&emsp;我们在调用errors.New("")来返回一个错误时， 可以通过比较指针，来比较error是否相等， 实际上就是控制相同的错误我们只创建一个error对象。否则对象复制一下，在比较就是false了。
&emsp;&emsp;而且指针才会与nil相比较， 如果是字符串"" 虽然错误内容为空，但是还是有错误的。

### 1.2 实现自己的error类型
见代码: [define_yourself_error_test.go](error/define_yourself_error_test.go)

```go
package go_error

import (
	"fmt"
	"testing"
)

type MyError struct {
	s string
}

func (e MyError) Error() string {
	return e.s
}

func NewMyError(msg string) error {
	return MyError{s: msg}
}

func myErrorTest() error {
	me := NewMyError("天下风云出我辈")
	if me != nil {
		fmt.Println(me.Error())
		return me
	}
	return nil
}

func TestError(t *testing.T) {
	myErrorTest()
}
```