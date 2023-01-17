package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02"))                    // 精确到日
	fmt.Println(now.Format("2006-01-02 15:04:05"))           // 精确到秒
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))       // 精确到毫秒
	fmt.Println(now.Format("2006-01-02 15:04:05.000000"))    // 精确到微秒
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000")) // 精确到纳秒
}
