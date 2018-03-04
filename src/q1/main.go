package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const ARRAYLEN = 200
const BUFSIZE = 10

var (
	wg          sync.WaitGroup
	sharedArray [ARRAYLEN]int
)

func initSharedArray(signalChan chan int, sharedArray *[ARRAYLEN]int) {
	for i := 1; i < ARRAYLEN; i++ {
		sharedArray[i] = rand.Intn(ARRAYLEN)
	}
	// init done
	signalChan <- 1
}

func idConsumer(snowFlake *SnowFlake, do func(uint32)) {
	defer wg.Done()

	if snowFlake.isStop {
		return
	}

	index := snowFlake.GetId()
	if index == STOPITERATOR {
		// notify all routines to stop
		snowFlake.isStop = true
		return
	}

	// consum it
	do(index)

	wg.Add(1)
	go idConsumer(snowFlake, do)
}

func worker(id uint32) {
	fmt.Println(id)
}

func main() {
	startSignal := make(chan int)
	go initSharedArray(startSignal, &sharedArray)

	snowFlake := NewSnowFlaker()
	defer snowFlake.Close()

	snowFlake.Start(&wg, 2, ARRAYLEN-1)

	// waiting for Shared Array initialized
	<-startSignal

	wg.Add(1)
	go idConsumer(snowFlake, worker)

	// waiting for all workers done
	wg.Wait()
}
