package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Everyone!")
	fmt.Println("My name is Gonzalo 'Gonzo' Fernandez")
	foo()
	FinalCountdown()
}

func foo() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// FinalCountdown is FinalExit
func FinalCountdown() {
	fmt.Println("duh duh duuuhhh duuuuhh....duh duh dooh dooh dooh")
}
