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

### 1.3 后台执行应该交给调用者

### 1.4 使用WaitGroup等待所有goroutine返回
[WaitGroup的Demo](./sync/waitgroup_test.go)

```go
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
```

### 1.5 读写锁
[读写锁demo代码](./sync/rwmutext_test.go)
读写锁上写锁时，所有上锁操作都将被阻塞。
读写锁上杜锁时，上写锁操作会被阻塞，但是上读锁的操作不会。
特别的，当一个读写锁已经被上了读锁，但是有一个写锁在阻塞中(尝试上锁)时，后续所有的读锁都会被阻塞
所以下面的代码极大概率会出现这种打印:
```
// go4: getReadLock...
// go4: releaseReadLock...
// go1: getWriteLock...
// go1: releaseWriteLock...
// go3: getReadLock...
// go2: getReadLock...
// go2: releaseReadLock...
// go3: releaseReadLock...
```

```go
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
```

### 1.6 atomic.Value 进行原子替换(适用于读多写少的场景)
[atomic.Value的Demo](./sync/atomic_value_test.go)

```go
package sync_test

import (
	"log"
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

func TestAtomicValue(t *testing.T) {
	var v atomic.Value
	v.Store(&Config{}) // 这行不能省
	go func() {
		i := 0
		for {
			i++
			// 错误示范
			// cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
			cfg := &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			v.Store(cfg)
		}
	}()
	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				cfg := v.Load().(*Config)
				log.Printf("%v\n\n", cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
```

### 1.7 使用ErrGroup同时请求多条数据
[ErrGroup的官方demo](./sync/errgroup_test.go)
[ErrGroup的官方文档连接](https://pkg.go.dev/golang.org/x/sync/errgroup)

### 1.8 once.Do进行线程安全的初始化
once.Do可以提供init函数一样的初始化功能，但是init函数会在程序已开始执行的时候就加载，当变量很久都没使用到时，会造成内存浪费。once.Do可以保证代码只执行一次，配合闭包即可完成初始化。
但是说实话，还不如用init，once.Do涉及到加锁的操作，假如后面要频繁读取这个配置，会影响到程序性能，总的来说感觉可有可无。
[once.Do的demo](./sync/once_do_test.go)
