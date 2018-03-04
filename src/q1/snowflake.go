package main

import (
	"math"
	"sync"
	"sync/atomic"
)

const STOPITERATOR = math.MaxUint32
const MAXGENERATORNUM = 10

type SnowFlake struct {
	isStop            bool
	atomIndex         uint32
	CustomerIndexChan chan uint32
}

func NewSnowFlaker() *SnowFlake {
	r := &SnowFlake{isStop: false, atomIndex: 0, CustomerIndexChan: make(chan uint32, BUFSIZE)}
	r.CustomerIndexChan <- 0
	return r
}

func (this *SnowFlake) IncreasedIdGenerator(wg *sync.WaitGroup, maxId uint32) {
	wg.Add(1)
	defer wg.Done()

	for {
		if this.atomIndex >= maxId {
			// stop signal in the end
			this.CustomerIndexChan <- STOPITERATOR
			break
		}
		// support multiple generator
		this.CustomerIndexChan <- atomic.AddUint32(&this.atomIndex, 1)
	}
}

func (this *SnowFlake) GetId() uint32 {
	return <-this.CustomerIndexChan
}

func (this *SnowFlake) Start(wg *sync.WaitGroup, workNum int, maxId uint32) {
	if workNum < 0 || workNum > MAXGENERATORNUM {
		workNum = MAXGENERATORNUM
	}

	for i := 1; i <= workNum; i++ {
		go this.IncreasedIdGenerator(wg, maxId)
	}
}

func (this *SnowFlake) Close() {
	close(this.CustomerIndexChan)
}
