package main

import (
	"fmt"
)

const fromUSDtoEUR float64 = 0.82
const fromUSDtoRUB float64 = 82.0
const fromEURtoRUB float64 = fromUSDtoRUB / fromUSDtoEUR

type tAllConv map[string]tConv

type tConv map[int]float64

var convToUSD = tConv{
	1: 1.0,
	2: 1.0 / fromUSDtoRUB,
	3: 1.0 / fromUSDtoEUR,
}

var convFromUSD = tConv{
	1: 1.0,
	2: fromUSDtoRUB,
	3: fromUSDtoEUR,
}

var allConv = tAllConv{
	"from": convFromUSD,
	"to":   convToUSD,
}

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

	_ = fromEURtoRUB
}

func readsum() float64 {
	var sum float64

	fmt.Println("Введите сумму:")
	for {
		_, er := fmt.Scan(&sum)
		if er != nil {
			fmt.Println("Ошибка ввода, повторите ввод")
		} else {
			return sum
		}
	}

}

func exchange(sum float64, inPutCurrency, outPutCurrency int) float64 {
	//rate := allConv["from"][outPutCurrency]
	// Сумма * курс перевода в доллар из входящей валюты * курс перевода из доллара в получаему валюту
	return sum * allConv["to"][inPutCurrency] * allConv["from"][outPutCurrency]
	//return sum * convToUSD[inPutCurrency] * convFromUSD[outPutCurrency]
}

func readcurrency() int {
	var readcurrency int
	fmt.Println("USD - 1")
	fmt.Println("RUB - 2")
	fmt.Println("EUR - 3")

	for {
		_, err := fmt.Scan(&readcurrency)

		if err != nil {
			fmt.Println("Ошибка ввода, повторите ввод")
			continue
		}

		if readcurrency == 1 || readcurrency == 2 || readcurrency == 3 {
			return readcurrency
		}

	}
}
