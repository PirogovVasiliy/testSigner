package main

import (
	"Signer/internal"
	"fmt"
	"time"

	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	nodeURL := os.Getenv("NODE_URL")
	client, _ := internal.Connect(nodeURL)
	defer internal.Disconnect(client)

	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	accumInstance := internal.GetDeployedContract(client, contractAddress)

	addChan := make(chan internal.AddChanal)
	excChan := make(chan internal.ExcChanal)

	go internal.AddListener(accumInstance, addChan)
	go internal.ExcListener(accumInstance, excChan)

	privateKey := os.Getenv("DEPLOYER_PRIVATE_KEY")

	accPK := os.Getenv("ACC_PRIVATE_KEY")
	go func() {
		time.Sleep(15 * time.Second)
		internal.CallSendEthersTwo(client, contractAddress, accPK)
		time.Sleep(15 * time.Second)
		internal.CallRecieveTwo(client, contractAddress, accPK)
	}()
	for {
		select {
		case excEvent := <-excChan:
			fmt.Println("Получено событие Exceed!")
			fmt.Println("amount:", excEvent.Amount.String())
			fmt.Println("message:", excEvent.Message)
			fmt.Println("----------------------------------------")
			internal.CallWithdrawEthersTwo(client, contractAddress, privateKey)
		case addEvent := <-addChan:
			fmt.Println("Получено событие Adding!")
			fmt.Println("address:", addEvent.Address)
			fmt.Println("amount:", addEvent.Amount.String())
			balance := internal.CallCheckBalanceTwo(client, contractAddress)
			fmt.Println("balance:", balance)
			fmt.Println("----------------------------------------")
		}
	}
}
