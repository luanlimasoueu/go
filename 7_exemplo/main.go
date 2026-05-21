package main

import "fmt"

func main() {

	var printValue string = "Hello World"
	printMe(printValue)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}
