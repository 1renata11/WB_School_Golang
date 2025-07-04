package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var channel []int
var mutex sync.Mutex

func Generate(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex.Lock()
		channel = append(channel, rand.Int())
		mutex.Unlock()
		time.Sleep(time.Second)
	}
}

func Get(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex.Lock()
		if len(channel) == 0 {
			mutex.Unlock()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			continue
		}
		val := channel[len(channel)-1]
		channel = channel[:len(channel)-1]
		mutex.Unlock()
		log.Printf("%v ", val)
	}
}

func L13() {
	wg := &sync.WaitGroup{}
	var n int
	log.Print("Введите число: ")
	fmt.Scanln(&n)
	go Generate(wg)
	wg.Add(1)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go Get(wg)
	}
	wg.Wait()
}
