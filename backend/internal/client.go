package internal

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Connect(nodeURL string) (client *ethclient.Client, chainID *big.Int) {
	var err error
	client, err = ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatalln("Ошибка подключения к Hardhat!", err)
	}
	fmt.Println("Подключено к Hardhat!")

	chainID, err = client.ChainID(context.Background())
	if err != nil {
		log.Fatalln("Не удалось получить Chain ID!", err)
	}
	fmt.Printf("Chain ID: %s\n", chainID.String())
	fmt.Println("----------------------------------------")
	return
}

func Disconnect(client *ethclient.Client) {
	client.Client().Close()
	fmt.Println("Закрытие соединения с Hardhat!")
}
