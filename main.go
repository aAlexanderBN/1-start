package main

import "fmt"

const fromUSDtoEUR float64 = 0.82
const fromUSDtoRUB float64 = 82.0
const fromEURtoRUB float64 = fromUSDtoRUB / fromUSDtoEUR

func main() {
	fmt.Println(fromEURtoRUB, fromUSDtoEUR, fromUSDtoRUB)
	var sum float64
	sum = readsum()

	exchange(sum, "RUB", "EUR")
	fmt.Println(sum)
}

func readsum() float64 {
	var sum float64
	fmt.Println("Введите сумму:")
	_, er := fmt.Scan(&sum)
	if er != nil {
		fmt.Println("Ошибка ввода")
	}
	return sum
}

func exchange(sum float64, inputcurrency, outputcurrency string) float64 {

	return sum
}
