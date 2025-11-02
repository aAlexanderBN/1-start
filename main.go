package main

import (
	"fmt"
)

const fromUSDtoEUR float64 = 0.82
const fromUSDtoRUB float64 = 82.0
const fromEURtoRUB float64 = fromUSDtoRUB / fromUSDtoEUR

func main() {

	var inputcurrency, outputcurrency int
	var sum float64
	var toExit string

	for {
		fmt.Println("Введите входящую валюту")
		inputcurrency = readcurrency()

		sum = readsum()

		fmt.Println("Введите исходящую валюту")
		outputcurrency = readcurrency()

		sum = exchange(sum, inputcurrency, outputcurrency)

		fmt.Printf("\nВ новой валюте сумма будет: %.0f", sum)
		fmt.Printf("\nДля повтора расчета нажмите [y], для выхода любую другу клавишу")
		fmt.Scan(&toExit)
		if toExit != "y" {
			break
		}
	}
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

func exchange(sum float64, inPutCurrency, outPutCurrency int) float64 {

	var toUSD, toOutPutCurrency float64

	switch inPutCurrency {
	case 1:
		toUSD = 1
	case 2:
		toUSD = 1 / fromUSDtoRUB
	case 3:
		toUSD = 1 / fromUSDtoEUR
	}

	switch outPutCurrency {
	case 1:
		toOutPutCurrency = 1
	case 2:
		toOutPutCurrency = fromUSDtoRUB
	case 3:
		toOutPutCurrency = fromUSDtoEUR
	}

	return sum * toUSD * toOutPutCurrency
}

func readcurrency() int {
	var readcurrency int
	fmt.Println("USD - 1")
	fmt.Println("RUB - 2")
	fmt.Println("EUR - 3")

	for {
		_, err := fmt.Scan(&readcurrency)

		if err != nil {
			continue
		}

		if readcurrency != 1 || readcurrency != 2 || readcurrency != 3 {
			return readcurrency
		}

	}
}
