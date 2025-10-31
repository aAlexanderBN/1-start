package main

import "fmt"

const fromUSDtoEUR float64 = 0.82
const fromUSDtoRUB float64 = 82.0
const fromEURtoRUB float64 = fromUSDtoRUB / fromUSDtoEUR

func main() {
	fmt.Println(fromEURtoRUB, fromUSDtoEUR, fromUSDtoRUB)
}
