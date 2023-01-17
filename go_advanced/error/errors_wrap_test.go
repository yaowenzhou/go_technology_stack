package go_error

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func errorsWrapTest() error {
	err := errors.New("一入江湖岁月催")
	return err
}

func TestErrorsWrap(t *testing.T) {
	fmt.Printf("errorsWrapTest: %+v", errorsWrapTest())
}
