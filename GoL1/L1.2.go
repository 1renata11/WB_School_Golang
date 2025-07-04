package main

import (
	"log"
	"sync"
)

var m = []int{2, 4, 6, 8, 10}
var n [5]int

func count(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	n[i] = m[i] * m[i]
}

func L12() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go count(i, &wg)
	}
	wg.Wait()
	for i := 0; i < 5; i++ {
		log.Printf("%v ", n[i])
	}
}
