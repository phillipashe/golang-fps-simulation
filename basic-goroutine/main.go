package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	startTimer := time.Now()
	// defer runs after the function exits
	// this uses anon func but can pass a named func too
	defer func() {
		fmt.Println(time.Since(startTimer))
	}()

	names := [4]string{"Thomas", "Caleb", "Shelby", "Phillip"}
	getNames(names)
	wg.Wait()
}

func shootLaser(name string) {
	fmt.Printf("Shooting a laser at %s\n", name)
	wg.Done()
}

func getNames(names [4]string) {
	for _, name := range names {
		wg.Add(1)
		go shootLaser(name)
	}
}
