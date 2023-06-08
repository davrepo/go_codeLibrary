package main

import (
	"fmt"
)

// Ordered interface is either int, float64 or string
type Ordered interface {
	int | float64 | string
}

// min is a func with generic type T, which is either int, float64 or string
// the parameter values is a slice of T
// so for example, if T is int, then values is a slice of type int
// thus, T is a type parameter, Ordered is a type constraint
// min func returns the minimum value in a slice of Ordered values
func min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T // zero is a keyword? no, it's a variable name
		return zero, fmt.Errorf("min of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}
	return m, nil
}

func main() {
	fmt.Println(min([]int{}))                 // 0 min of empty slice
	fmt.Println(min([]string{}))              //  min of empty slice
	fmt.Println(min([]float64{2, 1, 3}))      // 1 <nil>
	fmt.Println(min([]string{"B", "A", "C"})) // A <nil>
}
