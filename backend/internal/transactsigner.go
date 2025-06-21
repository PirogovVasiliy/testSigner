package internal

import (
	"Signer/internal/contract"
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CallWithdrawEthers(accumInstance *contract.Accum, chainID *big.Int, privateKeyStr string) {

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatalln("Ошибка получения private key!", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalln("Ошибка создания трансактора!", err)
	}

	_, err = accumInstance.WithdrawEther(auth)
	if err != nil {
		log.Println("Ошбика отправки транзакции WithdrawEthers!", err)
	}
}

func CallWithdrawEthersTwo(client *ethclient.Client, contractAddress string, privateKeyStr string) {
	transactor, _ := contract.NewAccumTransactor(common.HexToAddress(contractAddress), client)

	privateKey, _ := crypto.HexToECDSA(privateKeyStr)
	chainId, _ := client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	transactor.WithdrawEther(auth)
}

func CallCheckBalance(accumInstance *contract.Accum) *big.Int {
	balance, err := accumInstance.CheckBalance(&bind.CallOpts{})
	if err != nil {
		log.Println("Ошбика отправки транзакции CheckBalacne!", err)
	}

	return balance
}

func CallCheckBalanceTwo(client *ethclient.Client, contractAddress string) *big.Int {

	caller, err := contract.NewAccumCaller(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Panic(err)
	}
	balance, _ := caller.CheckBalance(&bind.CallOpts{})
	return balance
}

func CallSendEthers(accumInstance *contract.Accum, chainID *big.Int, privateKeyStr string) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatalln("Ошибка получения private key!", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalln("Ошибка создания трансактора!", err)
	}
	auth.Value = big.NewInt(9000000000000000000)

	_, err = accumInstance.SendEthers(auth)
	if err != nil {
		log.Println("Ошбика отправки транзакции WithdrawEthers!", err)
	}
}

func CallSendEthersTwo(client *ethclient.Client, contractAddress string, privateKeyStr string) {
	tracnsactor, _ := contract.NewAccumTransactor(common.HexToAddress(contractAddress), client)
	privateKey, _ := crypto.HexToECDSA(privateKeyStr)
	chainId, _ := client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	auth.Value = big.NewInt(9000000000000000000)
	tracnsactor.SendEthers(auth)

}

func CallRecieve(accumInstance *contract.Accum, chainID *big.Int, privateKeyStr string) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatalln("Ошибка получения private key!", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalln("Ошибка создания трансактора!", err)
	}
	auth.Value = big.NewInt(4000000000000000000)

	_, err = accumInstance.Receive(auth)
	if err != nil {
		log.Println("Ошбика отправки транзакции WithdrawEthers!", err)
	}
}

func CallRecieveTwo(client *ethclient.Client, contractAddress string, privateKeyStr string) {
	tracnsactor, _ := contract.NewAccumTransactor(common.HexToAddress(contractAddress), client)
	privateKey, _ := crypto.HexToECDSA(privateKeyStr)
	chainId, _ := client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	auth.Value = big.NewInt(4000000000000000000)
	tracnsactor.Receive(auth)
}
