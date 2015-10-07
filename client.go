package momentum

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
	"net/http"
)

type MomentumClient struct {
	mutex *sync.Mutex
	log   *log.Logger
}

func InitClient() *MomentumClient {
	return &MomentumClient{mutex: &sync.Mutex{}, log: log.New(os.Stderr, "", log.LstdFlags)}
}

type AsyncResult struct {
	result chan interface{}
	date   *time.Time
}

//Add provides call remote procedure
func (client *MomentumClient) Get(title string, arguments interface{}) {
	_, err := json.Marshal(arguments)
	if err != nil {
		panic(err)
	}
}

//Call provides sending request to the server
func (client *MomentumClient) Call(title string, args interface{}) *AsyncResult {
	return client.sendArgs(title, args)
}

func (client *MomentumClient) sendArgs(title string) *AsyncResult {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	return &AsyncResult{}
}
