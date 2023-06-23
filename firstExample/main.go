package main

import (
	"fmt"
	"sync"
)

func printSomeThing(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, w := range words {
		go printSomeThing(fmt.Sprintf("%d: %s", i, w), &wg)
	}

	wg.Wait()

	wg.Add(1)

	printSomeThing("This is the second thing to be printed!", &wg)
}
