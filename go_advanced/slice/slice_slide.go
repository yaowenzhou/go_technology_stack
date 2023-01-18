package slice

import (
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func SliceSlide() {
	go func() {
		var s []int64
		var ss []int64
		for j := 0; j < 10000; j++ {
			ss = append(ss, rand.Int63())
		}
		s = append(s, ss...)
		for i := 0; i < 1000000000; i++ {
			ss = ss[0:0]
			for j := 0; j < 10000; j++ {
				ss = append(ss, rand.Int63())
			}
			s = append(s, ss...)
			time.Sleep(time.Second * 3)
			s = s[len(s)-1 : len(s)-1]
		}
	}()
	http.ListenAndServe("localhost:6060", nil) // 启动服务
}
