package momentum

import (
	"log"
	"os"
	"sync"
	"time"
)

type MomentumClient struct {
	mutex *sync.Mutex
	log   *log.Logger
}

func InitClient() *MomentumClient {
	return MomentumClient{mutex: *sync.Mutex{}, log: log.New(os.Stderr, "", log.LstdFlags)}
}

type AsyncResult struct {
	result interface{}
	date   *time.Time
}

//Add provides call remote procedure
func (client *MomentumClient) Get(title string, arguments interface{}) {
	err := json.Marshal(arguments)
	if err != nil {
		panic(err)
	}
}

func (client *MomentumClient) Call(title string) *AsyncResult {
	client.mutex.Lock()
	defer client.mutex.Unlock()
}
