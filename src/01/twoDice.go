package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var dice1 int = rand.Intn(6)+1 // 0 based
	var dice2 int = rand.Intn(6)+1 // 0 based
	fmt.Println("Hello, World!", dice1, dice2)
}
