package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello Universe")
}

//Greeter exported
var Greeter greeting
