package main

import (
	"fmt"
	"github.com/vvasquezh07/golang-testing-libraries/app/service"
)

func main() {
	fmt.Println(PrintMessage())
	fmt.Println(Sum(2,3))
	MocksWithTestify()
}

func PrintMessage() string {
	return "Hello World"
}

func Sum(x, y int) int {
	return x + y
}

func MocksWithTestify() {
	fmt.Println("Here starts the mocking project")
	fmt.Println("Available at: https://github.com/lamida/testify-mock")

	d := service.NewDB()

	g := service.NewGreeter(d, "en")
	fmt.Println(g.Greet()) // Message is: hello
	fmt.Println(g.GreetInDefaultMsg()) // Message is: default message

	g = service.NewGreeter(d, "es")
	fmt.Println(g.Greet()) // Message is: holla

	g = service.NewGreeter(d, "random")
	fmt.Println(g.Greet()) // Message is: bzzzz
}
