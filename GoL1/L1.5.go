package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var channel1 []int
var mutex1 sync.Mutex

func Generate1(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex1.Lock()
		channel1 = append(channel1, rand.Int())
		mutex1.Unlock()
		time.Sleep(time.Second)
	}
}

func Get1(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex1.Lock()
		if len(channel1) == 0 {
			mutex1.Unlock()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			continue
		}
		val := channel1[len(channel1)-1]
		channel1 = channel1[:len(channel1)-1]
		mutex1.Unlock()
		log.Printf("%v ", val)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	var n int
	log.Print("Введите число: ")
	fmt.Scanln(&n)
	var t int
	log.Print("Введите количество миллисекунд: ")
	fmt.Scanln(&t)
	go Generate1(wg)
	wg.Add(1)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go Get1(wg)
	}
	<-time.After(time.Duration(t) * time.Millisecond)
	fmt.Println("Таймаут")
}
