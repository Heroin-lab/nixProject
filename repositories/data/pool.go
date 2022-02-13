package main

import (
	"log"
	"sync"
)

type Pool struct {
	Count  int
	Sender chan Shop
	Ender  chan bool
}

func NewPool(count int) *Pool {
	return &Pool{
		Count:  count,
		Sender: make(chan Shop, count*2),
		Ender:  make(chan bool),
	}
}

func (p *Pool) Run(wg *sync.WaitGroup, handler func(author Shop)) {
	defer wg.Done()
	var shop Shop
	for {
		select {
		case shop = <-p.Sender:
			handler(shop)
		case <-p.Ender:
			//fmt.Println(<- p.Sender)
			log.Println("I am finish")
			return
		}
	}
}

func (p *Pool) Stop() {
	for i := 0; i < p.Count; i++ {
		p.Ender <- false
	}
	close(p.Sender)
	close(p.Ender)
}
