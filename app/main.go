package main

import "fmt"

func main() {
	fmt.Println(PrintMessage())
	fmt.Println(Sum(2,3))
}

func PrintMessage() string {
	return "Hello World"
}

func Sum(x, y int) int {
	return x + y
}
