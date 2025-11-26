package main

import (
	"fmt"
	"log"

	"xrobot/internal/cryptocurrencies/cryptobot-sdk-golang/cryptobot"
)

func getExchangeRates(client *cryptobot.Client) {
	exchangeRates, err := client.GetExchangeRates()
	if err != nil {
		log.Fatalln(err)
	}

	for _, exchangeRate := range exchangeRates {
		fmt.Printf(
			"IsValid - %v, Rate - %v, Source - %v, Target - %v\n",
			exchangeRate.IsValid,
			exchangeRate.Rate,
			exchangeRate.Source,
			exchangeRate.Target,
		)
	}
}
