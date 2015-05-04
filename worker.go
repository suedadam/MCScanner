package main

import (
	"fmt"
)

type pool struct {
	checks  chan string
	workers int
	endings []chan bool
}

func newPool(workers int) *pool {
	p := &pool{
		// Buffer the channel with 10x as many IPs as workers.
		checks:  make(chan string, workers*10),
		workers: workers,
	}
	p.spawn()

	return p
}

// Spawns out the number of workers in the pool.
func (p *pool) spawn() {
	for i := 0; i < p.workers; i++ {
		end := make(chan bool)
		p.endings = append(p.endings, end)
		go p.work(end)
	}
}

// Adds a new IP to be checked.
func (p *pool) add(ip string) {
	p.checks <- ip
}

// Listens from and works on the checks until the "end" signal is
// sent.
func (p *pool) work(end chan bool) {
	for {
		select {
		case ip := <-p.checks:
			if isMinecraft(ip) {
				fmt.Println(ip)
			}
		case <-end:
			return
		}
	}
}

// Closes the pool. Should be sent when you're done adding IPs to
// be checked. Blocks until the pool is finished working.
func (p *pool) end() {
	for _, end := range p.endings {
		end <- true
	}
}
