package main

import (
	"fmt"
	"log"
)

type ClickEvent struct {
	// ...
}

type HoverEvent struct {
	// ...
}

var eventCounts = make(map[string]int) // type -> count

func recordEvent(evt interface{}) { // empty interface
	switch evt.(type) { // type assertion / type switch
	case *ClickEvent: // pointer to ClickEvent
		eventCounts["click"]++
	case *HoverEvent:
		eventCounts["hover"]++
	default:
		log.Printf("warning: unknown event: %#v of type %T\n", evt, evt)
	}
}

func main() {
	recordEvent(&ClickEvent{}) // a pointer is passed, not struct itself
	recordEvent(&HoverEvent{})
	recordEvent(&ClickEvent{})
	recordEvent(3) // default case, since 3 is not a pointer to ClickEvent or HoverEvent
	// 2021/04/30 15:07:17 warning: unknown event: 3 of type int

	fmt.Println("event counts:", eventCounts)
	// event counts: map[click:2 hover:1]
}
