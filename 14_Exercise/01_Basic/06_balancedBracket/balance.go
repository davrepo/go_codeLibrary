package main

import (
	"flag"
	"log"
)

// assess whether expression has balanced brackets, {}, (), []
// do not consider order of {} () []
// use a stack to keep track of opening brackets

// isBalanced returns whether the given expression
// has balanced brackets.

// operatorType represents the type of operator in an expression
type operatorType int // alias for int

const (
	// iota is an incrementing int constant generator
	openBracket   operatorType = iota // 0
	closedBracket                     // 1
	otherOperator                     // 2
)

// bracketPairs is the map legal bracket pairs
var bracketPairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

// getOperatorType return the operator type of the give operator.
func getOperatorType(op rune) operatorType {
	for ob, cb := range bracketPairs {
		switch op {
		case ob: // key
			return openBracket // return 0
		case cb: // value
			return closedBracket // return 1
		}
	}

	return otherOperator // return 2
}

// stack is a simple LIFO stack implementation using a slice.
type stack struct {
	elems []rune
}

// push adds a new element to the stack.
func (s *stack) push(e rune) {
	s.elems = append(s.elems, e)
}

// pop removes the last element from the stack.
func (s *stack) pop() *rune {
	if len(s.elems) == 0 {
		return nil
	}
	n := len(s.elems) - 1
	last := s.elems[n]
	s.elems = s.elems[:n]
	// returns a pointer because
	// ensure that caller receives a valid value even after the element is removed from the stack
	// when you remove an element from the stack by slicing s.elems[:n],
	// the original last value is removed from the elems slice.
	// f you were to return last directly (without &), it would be returning a copy of the last value,
	// which could potentially be pointing to invalid or unexpected data.
	return &last
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	s := stack{} // create a new stack
	for _, e := range expr {
		switch getOperatorType(e) {
		case openBracket: // 0
			s.push(e)
		case closedBracket: // 1
			last := s.pop()
			if last == nil || bracketPairs[*last] != e {
				return false
			}
		}
	}

	return len(s.elems) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
