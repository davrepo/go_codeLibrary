package main

import (
	"fmt"
	"log"
	"sync"
)

// Consumer sends an order to the barista, and receives a coffee
// Producer receives an order from the customer, and sends a coffee
// Assumptions:
// - once a certain amount of orders is fulfilled, the coffee shop closes
// - do not care about which customer has ordered which coffee
// - not maintain any coffee types; all orders represented by empty struct, i.e. struct{}

// setup constants
const baristaCount = 3
const customerCount = 6
const maxOrderCount = 20

// the total amount of drinks that the bartenders have made
type coffeeShop struct {
	orderCount int
	orderLock  sync.Mutex // mutux

	orderCoffee  chan struct{}
	finishCoffee chan struct{}
	closeShop    chan struct{} // signal channel for closing the shop
}

// registerOrder ensures that the order made by the baristas is counted
// once the orderCount reaches maxOrderCount, the shop is closed
func (p *coffeeShop) registerOrder() {
	p.orderLock.Lock()
	defer p.orderLock.Unlock()
	p.orderCount++
	if p.orderCount == maxOrderCount {
		// when channel is closed, all receivers of the channel
		// will receive a zero value of the channel type, i.e. an empty struct{}
		// i.e. any goroutine that is blocked while trying to read from this channel
		// will immediately proceed once the channel is closed
		close(p.closeShop) // close the signal channel
	}
}

// Producer
func (p *coffeeShop) barista(name string) {
	for {
		select {
		case <-p.orderCoffee: // receives an order
			p.registerOrder()
			log.Printf("%s makes a coffee.\n", name)
			p.finishCoffee <- struct{}{} // sends a coffee
		case <-p.closeShop: // abort signal from signal channel
			log.Printf("%s stops working. Bye!\n", name)
			return
		}
	}
}

// Consumer
func (p *coffeeShop) customer(name string) {
	for {
		select {
		case p.orderCoffee <- struct{}{}: // sends an order
			log.Printf("%s orders a coffee!", name)
			<-p.finishCoffee // receives a coffee
			log.Printf("%s enjoys a coffee!\n", name)
		case <-p.closeShop: // abort signal from signal channel
			log.Printf("%s leaves the shop! Bye!\n", name)
			return
		}
	}
}

func main() {
	log.Println("Welcome to the coffee shop!")

	orderCoffee := make(chan struct{}, baristaCount)
	finishCoffee := make(chan struct{}, baristaCount)
	closeShop := make(chan struct{}) // signal channel for closing the shop, unbuffered

	p := coffeeShop{
		orderCoffee:  orderCoffee,
		finishCoffee: finishCoffee,
		closeShop:    closeShop,
	}

	for i := 0; i < baristaCount; i++ {
		go p.barista(fmt.Sprint("Barista-", i))
	}

	for i := 0; i < customerCount; i++ {
		go p.customer(fmt.Sprint("Customer-", i))
	}

	// this line is serving as a blocking point that
	// prevents the main goroutine from terminating and closing the program prematurely
	// When the closeShop channel is closed, this line will receive the zero value,
	// and the main function can continue to its end, effectively closing the application.
	<-closeShop
	log.Println("The coffee shop has closed! Bye!")
}
