package main

import (
	"fmt"
	"sync"
)

// Dining Philosophers - 5 philosophers.
// state: either EAT or THINK
// must use 2 chopsticks to eat
// b/c there are only 5 chopsticks, only 2 philosophers can eat at a time
// can eat unlimited times per philosopher
// avoid deadlock in code

// specification:
// - each fork must have its own thread (goroutine)
// - each philosopher must have its own thread (goroutine)
// - philosophers and forks must communicate with each other *only* by using channels
// - each philosopher must eat at least 3 times, no deadlock allowed / *comment* in code why the system does not deadlock
// - philosophers must display (print on screen) any state change (eating or thinking) during their execution.

const (
	THINKING = iota
	EATING
	NUMBER_OF_PHILOSOPHERS = 5
	NUMBER_OF_CHOPSTICKS   = 5
	NUMBER_OF_TIME_EAT     = 3
)

type Chopsticks struct {
	number int
	ch     chan int
}

type Philosopher struct {
	number int
	state  int
	ch     chan int
}

func (c *Chopsticks) manage() {
	for {
		c.ch <- 1 // Send a signal that the chopstick is available
		<-c.ch    // Wait for a philosopher to send a signal that it's done eating
	}
}

func (p *Philosopher) eat(c1, c2 *Chopsticks, wg *sync.WaitGroup) {
	defer wg.Done()

	// introduce asymmetry, make Philosopher 0 pick up the right chopstick first
	// while all other philosophers pick up the left chopstick first, to avoid deadlock
	// Try to acquire both chopsticks
	if p.number == 0 {
		<-c2.ch
		<-c1.ch
	} else {
		<-c1.ch
		<-c2.ch
	}
	p.state = EATING
	fmt.Printf("Philosopher %d is eating.\n", p.number+1)

	// Release both chopsticks
	c1.ch <- 1
	c2.ch <- 1

	p.state = THINKING
	fmt.Printf("Philosopher %d has finished eating and is now thinking.\n", p.number+1)
}

func main() {
	var wg sync.WaitGroup

	chopsticks := make([]*Chopsticks, NUMBER_OF_CHOPSTICKS)
	for i := 0; i < NUMBER_OF_CHOPSTICKS; i++ {
		c := make(chan int, 1)
		chopstick := &Chopsticks{
			number: i,
			ch:     c,
		}
		chopsticks[i] = chopstick
		go chopstick.manage()
	}

	for i := 0; i < NUMBER_OF_PHILOSOPHERS; i++ {
		c := make(chan int, 1)
		philosopher := &Philosopher{
			number: i,
			state:  THINKING,
			ch:     c,
		}

		for j := 0; j < NUMBER_OF_TIME_EAT; j++ {
			wg.Add(1)
			go philosopher.eat(chopsticks[i], chopsticks[(i+1)%NUMBER_OF_CHOPSTICKS], &wg)
		}
	}

	wg.Wait()
}
