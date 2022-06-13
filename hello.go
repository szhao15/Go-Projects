package main

import (
	"fmt"
	"github.com/szhao15/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}