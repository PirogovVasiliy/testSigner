package internal

import (
	"Signer/internal/contract"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CallWithdrawEthers(accumInstance *contract.Accum, client *ethclient.Client, chainID *big.Int, privateKeyStr string) {

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

func CallCheckBalance(accumInstance *contract.Accum) *big.Int {
	balance, err := accumInstance.CheckBalance(&bind.CallOpts{})
	if err != nil {
		log.Println("Ошбика отправки транзакции CheckBalacne!", err)
	}

	return balance
}

func CallSendEthers(accumInstance *contract.Accum, client *ethclient.Client, chainID *big.Int, privateKeyStr string) {
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

func CallRecieve(accumInstance *contract.Accum, client *ethclient.Client, chainID *big.Int, privateKeyStr string) {
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
