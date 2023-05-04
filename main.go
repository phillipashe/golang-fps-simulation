package main

import "fmt"

func main() {

	names := [4]string{"Thomas", "Caleb", "Shelby", "Phillip"}
	shootLaser(names)

}

func shootLaser(names [4]string) {

	for _, name := range names {
		fmt.Printf("Shooting a laser at %s\n", name)
	}
}
