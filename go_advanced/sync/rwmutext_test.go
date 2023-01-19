package sync_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 下面的代码极大概率出现这种打印:
// go4: getReadLock...
// go4: releaseReadLock...
// go1: getWriteLock...
// go1: releaseWriteLock...
// go3: getReadLock...
// go2: getReadLock...
// go2: releaseReadLock...
// go3: releaseReadLock...
// 这是因为一旦一个写锁尝试上锁，则后续所有的读锁和写锁动作都将被阻塞
// 等写锁释放了之后，所有的读锁可以同时上锁

func TestRWMutex(t *testing.T) {
	var wg sync.WaitGroup
	var rwMutex sync.RWMutex
	wg.Add(4)
	go func() {
		rwMutex.Lock()
		fmt.Println("go1: getWriteLock...")
		time.Sleep(time.Second * 5)
		fmt.Println("go1: releaseWriteLock...")
		rwMutex.Unlock()
		wg.Done()
	}()
	go func() {
		rwMutex.RLock()
		fmt.Println("go2: getReadLock...")
		time.Sleep(time.Second * 5)
		fmt.Println("go2: releaseReadLock...")
		rwMutex.RUnlock()
		wg.Done()
	}()
	go func() {
		rwMutex.RLock()
		fmt.Println("go3: getReadLock...")
		time.Sleep(time.Second * 5)
		fmt.Println("go3: releaseReadLock...")
		rwMutex.RUnlock()
		wg.Done()
	}()
	go func() {
		rwMutex.RLock()
		fmt.Println("go4: getReadLock...")
		time.Sleep(time.Second * 5)
		fmt.Println("go4: releaseReadLock...")
		rwMutex.RUnlock()
		wg.Done()
	}()
	wg.Wait()
}
