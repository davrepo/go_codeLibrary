package main

import (
	"fmt"
	"sync"
)

// Dining Philosophers - 5 philosophers.
// state: either eat, or not eat
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
		<-c.ch // Wait for a philosopher to *request* the chopstick
		<-c.ch // Wait for a philosopher to *release* the chopstick
	}
}

func (p *Philosopher) eat(c1, c2 *Chopsticks, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < NUMBER_OF_TIME_EAT; i++ {
		// introduce asymmetry, make Philosopher 0 pick up the right chopstick first
		// while all other philosophers pick up the left chopstick first,
		if p.number == 0 {
			c2.ch <- 1
			c1.ch <- 1
		} else {
			c1.ch <- 1
			c2.ch <- 1
		}
		p.state = EATING
		fmt.Printf("starting to eat %d\n", p.number+1)
		p.ch <- p.number // how to get out of that channel??
		p.state = THINKING
		fmt.Printf("finishing eating, now thinking %d\n", p.number+1)
		<-c1.ch
		<-c2.ch
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(NUMBER_OF_PHILOSOPHERS)

	chopsticks := make([]*Chopsticks, NUMBER_OF_CHOPSTICKS)
	for i := 0; i < NUMBER_OF_CHOPSTICKS; i++ {
		c := make(chan int, 1)
		chopstick := &Chopsticks{
			number: i,
			ch:     c,
		}
		chopsticks[i] = chopstick
	}

	for i := 0; i < NUMBER_OF_CHOPSTICKS; i++ {
		go chopsticks[i].manage()
	}

	for i := 0; i < NUMBER_OF_PHILOSOPHERS; i++ {
		c := make(chan int, 1)
		philosopher := &Philosopher{
			number: i,
			state:  THINKING,
			ch:     c,
		}
		go philosopher.eat(chopsticks[i], chopsticks[(i+1)%NUMBER_OF_CHOPSTICKS], &wg)
	}

	wg.Wait()
}
