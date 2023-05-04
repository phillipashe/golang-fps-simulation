package main

import (
	"fmt"
	"time"
)

func main() {

	startTimer := time.Now()
	// defer runs after the function exits
	// this uses anon func but can pass a named func too
	defer func() {
		fmt.Println(time.Since(startTimer))
	}()

	names := [4]string{"Thomas", "Caleb", "Shelby", "Phillip"}
	shootLaser(names)

}

func shootLaser(names [4]string) {

	for _, name := range names {
		go fmt.Printf("Shooting a laser at %s\n", name)
	}
}
