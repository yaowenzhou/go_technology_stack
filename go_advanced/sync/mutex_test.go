package sync_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var m sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		m.Lock()
		fmt.Println("go1 in...")
		time.Sleep(time.Second * 5)
		m.Unlock()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		m.Lock()
		fmt.Println("go2 in...")
		time.Sleep(time.Second * 5)
		m.Unlock()
		wg.Done()
	}()
	wg.Wait()
}
