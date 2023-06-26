package main

import (
	"fmt"
	"sync"
	"time"
)

const hunger = 3

var philosophers = []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 3 * time.Second

func diningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()

	// print a message
	fmt.Println(philosopher, "is seated.")
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		fmt.Println(philosopher, "is hungry.")
		time.Sleep(sleepTime)

		// lock both forks
		leftFork.Lock()
		fmt.Printf("\t%s picked up the fork to his left.\n", philosopher)
		rightFork.Lock()
		fmt.Printf("\t%s picked up the fork to his right.\n", philosopher)

		// print a message
		fmt.Println(philosopher, "has both forks, and is eating.")
		time.Sleep(eatTime)

		// unlock the mutexes
		rightFork.Unlock()
		fmt.Printf("\t%s put down the fork on his right.\n", philosopher)
		leftFork.Unlock()
		fmt.Printf("\t%s put down the fork on his left.\n", philosopher)
		time.Sleep(sleepTime)
	}

	fmt.Println(philosopher, "is satisfied.")
	time.Sleep(sleepTime)

	fmt.Println(philosopher, "has left the table.")
}

func main() {

	// print intro
	fmt.Println("The Dining Philosophers Problem.")
	fmt.Println("--------------------------------")

	// spawn one goroutine for each philosopher
	wg.Add(len(philosophers))

	forkLeft := &sync.Mutex{}

	for i := 0; i < len(philosophers); i++ {

		forkRight := &sync.Mutex{}

		go diningProblem(philosophers[i], forkLeft, forkRight)

		forkLeft = forkRight
	}

	wg.Wait()

	fmt.Println("The table is empty.")
}
