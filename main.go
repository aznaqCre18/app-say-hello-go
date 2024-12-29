package main

import (
	"fmt"
	go_say_hello "github.com/aznaqCre18/go-say-hello/v2"
)

func main() {
	sayHello := go_say_hello.SayHello("Aziz")
	sayHowAreU := go_say_hello.HowAreYou("Aziz")
	askAge := go_say_hello.HowOlder()

	fmt.Println(sayHello)
	fmt.Println(sayHowAreU)
	fmt.Println(askAge)
}
