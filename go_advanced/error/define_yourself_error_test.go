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
