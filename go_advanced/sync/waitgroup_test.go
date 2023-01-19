package sync_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitgroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("go1 in...")
		time.Sleep(time.Second * 5)
		fmt.Println("go1 out...")
		wg.Done()
	}()
	go func() {
		fmt.Println("go2 in...")
		time.Sleep(time.Second * 5)
		fmt.Println("go2 out...")
		wg.Done()
	}()
	wg.Wait()
}
