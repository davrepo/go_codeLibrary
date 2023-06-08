package main

import (
	"example.com/backend"
)

func main() {
	a := backend.App{}
	a.Port = ":8080"
	a.Initialize()
	a.Run()
}
