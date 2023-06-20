package main

import (
	"fmt"
)

func main() {
	nums := []int{1}
	// fmt.Println(secondToLast(nums)) // will panic
	fmt.Println(safeSecondToLast(nums)) // 0 runtime error: index out of range [-1]
}

func safeSecondToLast(nums []int) (i int, err error) {
	// defer anomymous function
	// will execute at the end of fxn regardless whether fxn returns normally or through panic
	defer func() {
		// recover intercepts any potential panics, regains control over a panicking goroutine
		// recover() is only useful inside deferred functions
		// during normal execution, recover() returns nil
		// if panic occurs, recover() will capture the panic value
		// this deferred function then assigns this value to 'e'
		if e := recover(); e != nil { // e is interface{}
			// e is converted to string using fmt.Errorf() and assigned to err
			err = fmt.Errorf("%v", e)
		}

		// implicit return of err
		// since i is not assigned, it will have a default return value of 0
	}()

	return secondToLast(nums), nil
}

func secondToLast(nums []int) int {
	return nums[len(nums)-2] // will panic if len(nums) < 2
}
