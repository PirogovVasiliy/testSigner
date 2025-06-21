package internal

import (
	"Signer/internal/contract"
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployContract(client *ethclient.Client, chainID *big.Int, privateKeyStr string) *contract.Accum {

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatalln("Ошибка получения private key!", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalln("Ошибка создания трансактора!", err)
	}

	initialOwner := auth.From

	address, tx, accumInstance, err := contract.DeployAccum(auth, client, initialOwner)
	if err != nil {
		log.Fatalln("Ошибка deploy!", err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalln("Ошибка при ожидании подтверждения транзакции!", err)
	}

	fmt.Println("Контракт успешно развернут")
	fmt.Printf("Аккаунт деплоя (owner): %s\n", initialOwner.Hex())
	fmt.Printf("Адрес контракта: %s\n", address.Hex())
	fmt.Println("----------------------------------------")

	return accumInstance
}
