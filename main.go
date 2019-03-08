package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	const (
		nickel  = 0.05
		dime    = 0.10
		quarter = 0.25
	)

	var coins = [3]float64{nickel, dime, quarter}
	var balance float64

	for balance <= 20 {
		coin := coins[rand.Intn(3)]
		balance += coin

		fmt.Printf("Balance is now $%4.2f\n", balance)
	}
}
