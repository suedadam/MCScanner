package main

import (
	"fmt"
)

type pool struct {
	checks  chan string
	workers int
	ending  chan bool
}

func newPool(workers int) *pool {
	p := &pool{
		// Buffer the channel with 10x as many IPs as workers.
		checks:  make(chan string, workers*10),
		ending:  make(chan bool, workers),
		workers: workers,
	}
	p.spawn()

	return p
}

// Spawns out the number of workers in the pool.
func (p *pool) spawn() {
	for i := 0; i < p.workers; i++ {
		go p.work()
	}
}

// Adds a new IP to be checked.
func (p *pool) add(ip string) {
	p.checks <- ip
}

// Listens from and works on the checks.
func (p *pool) work() {
	for ip := range p.checks {
		if isMinecraft(ip) {
			fmt.Println(ip)
		}
	}

	p.ending <- true
}

// Closes the pool. Should be sent when you're done adding IPs to
// be checked. Blocks until the pool is finished.
func (p *pool) end() {
	close(p.checks)

	for i := 0; i < p.workers; i++ {
		<-p.ending
	}
}
