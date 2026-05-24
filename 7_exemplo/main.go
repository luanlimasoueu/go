package main

import (
	"fmt"
)

func main() {

	var printValue string = "Hello World"
	printMe(printValue)

	var numerador int = 11
	var denominador int = 2
	var result, remainder int = intDivision(numerador, denominador)
	fmt.Println(result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerador int, denominador int) (int, int) {
	var result int = numerador / denominador
	var remainder int = numerador % denominador
	return result, remainder
}
