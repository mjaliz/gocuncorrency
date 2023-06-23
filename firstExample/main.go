package main

import (
	"fmt"
	"time"
)

func printSomeThing(s string) {
	fmt.Println(s)
}

func main() {

	go printSomeThing("This is the first thing to be printed!")

	time.Sleep(1 * time.Second)

	printSomeThing("This is the second thing to be printed!")
}
