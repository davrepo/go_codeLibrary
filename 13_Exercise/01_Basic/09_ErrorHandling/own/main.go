package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// operators is the map of legal operators and their functions
var operators = map[string]func(x, y float64) (float64, error){
	"+": func(x, y float64) (float64, error) { return x + y, nil },
	"-": func(x, y float64) (float64, error) { return x - y, nil },
	"*": func(x, y float64) (float64, error) { return x * y, nil },
	"/": func(x, y float64) (float64, error) {
		if y == 0 {
			return 0, fmt.Errorf("divide by zero")
		}
		return x / y, nil
	},
}

// parseOperand parses a string to a float64
func parseOperand(op string) (float64, error) {
	parsedOp, err := strconv.ParseFloat(op, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot parse: %v", err)
	}
	return parsedOp, nil // if error occurs, will return zero value, i.e. 0.0
}

// calculate returns the result of a 2 operand mathematical expression
func calculate(expr string) (float64, error) {
	ops := strings.Fields(expr)
	if len(ops) != 3 {
		return 0, fmt.Errorf("invalid expression: need 3 ops, got %d", len(ops))
	}

	left, err := parseOperand(ops[0])
	if err != nil {
		return 0, fmt.Errorf("invalid left operand: %v", err)
	}

	right, err := parseOperand(ops[2])
	if err != nil {
		return 0, fmt.Errorf("invalid right operand: %v", err)
	}

	f, ok := operators[ops[1]]
	if !ok {
		return 0, fmt.Errorf("unknown operator: %s", ops[1])
	}

	result, err := f(left, right)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func main() {
	expr := flag.String("expr", "", "The expression to calculate on, separated by spaces.")
	flag.Parse()

	if *expr == "" {
		log.Fatal("No expression provided. Please use the -expr flag.")
	}

	result, err := calculate(*expr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s = %.2f\n", *expr, result)
}
