package main

import (
	"fmt"
	"log"

	"tron_robot/internal/cryptocurrencies/cryptobot-sdk-golang/cryptobot"
)

func getBalance(client *cryptobot.Client) {
	balance, err := client.GetBalance()
	if err != nil {
		log.Fatalln(err)
	}

	for _, asset := range balance {
		fmt.Printf("Currency - %s, available - %s\n", asset.CurrencyCode, asset.Available)
	}
}
