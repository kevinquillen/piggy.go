package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var coins = [3]float64{.05, .10, .25}
	var balance float64

	for balance <= 20 {
		balance += coins[rand.Intn(3)]
		fmt.Printf("Balance is now $%4.2f\n", balance)
	}
}
