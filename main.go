package main

import (
	"ddyy/go-algorithms/stack"
	"fmt"
)

func main() {
	expn := "(3+4)*5"
	value := stack.InfixToPostfix(expn)
	fmt.Println(expn)
	fmt.Println(value)
}