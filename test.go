package main

import (
	"fmt"
	"time"
)

func main() {
	month := time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(month)
	// 2021/04/30 17:53:15 GOT: {Time:2021-04-30 17:53:14.437272 +0300 IDT CPU:0.23 Memory:87.32}
}
