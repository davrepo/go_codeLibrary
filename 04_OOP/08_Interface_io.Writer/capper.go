package main

import (
	"fmt"
	"io"
	"os"
)

// Capper implements io.Writer and turns everything to uppercase
type Capper struct {
	// c.wtr is a field of the Capper struct, which implements the io.Writer interface.
	// It represents the underlying writer where the modified output will be written.
	wtr io.Writer
}

// Capper implements io.Writer interface because it has the Write() method
func (c *Capper) Write(p []byte) (n int, err error) {
	// lower case a has a value of 97, upper case A has a value of 65
	diff := byte('a' - 'A')

	// create a new slice of bytes with the same length as p as we don't want to modify p
	// make() allocates memory for the slice
	out := make([]byte, len(p))
	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}

	// invokes the Write method of the underlying writer,
	// passing out as the data to be written. This is done by calling Write on c.wtr.
	// The return value of c.wtr.Write(out) is then returned as the return value of the Write method of Capper
	return c.wtr.Write(out)
}

// when the Write method of Capper is called, it transforms the input bytes p to uppercase,
// stores them in the out byte slice, and then passes out to the Write method of the underlying writer (c.wtr).
// The return value of the underlying writer's Write method is
// then returned as the return value of the Write method of Capper

func main() {
	// {} is a struct literal
	// os.Stdout is a field of the os package and it implements the io.Writer interface
	// is os.Stdout is passed as the first argument to the Capper struct? Yes
	// that os.Stdout is stored in the wtr field of the Capper struct
	c := &Capper{os.Stdout}
	// Fprintln() takes an io.Writer as the first argument, and second argument is the string to print
	// what is Fprintln() doing? It is calling the Write() method of the Capper struct
	// and passing the string "Hello there" as the data to be written
	fmt.Fprintln(c, "Hello there")
	// HELLO THERE
}
