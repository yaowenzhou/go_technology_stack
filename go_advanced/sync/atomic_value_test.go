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
