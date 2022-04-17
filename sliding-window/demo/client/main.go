package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func Client() {
	res, err := http.Get("http://localhost:8001/api/up/v1")
	if err != nil {
		log.Println(err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("client log:", err)
		return
	}
	log.Printf("%d: %s", res.StatusCode, string(data))
}

func main() {
	rand.Seed(time.Now().Unix())
	wg := sync.WaitGroup{}
	// todo 还要由main 控制住这些 goroutine 的生命周期
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			Client()
			wg.Done()
		}()
	}
	wg.Wait()
}
