package sync_test

import (
	"log"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

type OnceDoConfig struct {
	Server string
	Port   int64
}

var (
	once   sync.Once
	config *OnceDoConfig
)

func ReadConfig() *OnceDoConfig {
	once.Do(func() {
		var err error
		config = &OnceDoConfig{Server: os.Getenv("TT_SERVER_URL")}
		config.Port, err = strconv.ParseInt(os.Getenv("TT_PORT"), 10, 0)
		if err != nil {
			config.Port = 8080 // default port
		}
		log.Println("init config")
	})
	return config
}

func TestOnceDo(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			_ = ReadConfig()
		}()
	}
	time.Sleep(time.Second)
}
