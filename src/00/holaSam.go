package main

import (
	"fmt"
	"time"
)

func main() {
	//name := "Sergio"
	// fmt.Println("Hola " + name)

	var (
		age  int
		when time.Time
	)

	age = 37
	when = time.Now()

	var firstName, lastName string = "Sergio", "Monsalve"
	fmt.Println("Hola " + firstName + " " + lastName)
	fmt.Println(age, when)
}
