package internal

import (
	"Signer/internal/contract"
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type AddChanal struct {
	Address string
	Amount  big.Int
}

func AddListener(accumInstance *contract.Accum, outChanal chan AddChanal) {

	addEventChanal := make(chan *contract.AccumAdding)

	opts := &bind.WatchOpts{
		Context: context.Background(),
	}

	sub, err := accumInstance.WatchAdding(opts, addEventChanal)
	if err != nil {
		log.Println("Ошибка подписки на события Adding!", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Fatalln("Ошибка чтения события Adding!", err)
		case event := <-addEventChanal:
			outChanal <- AddChanal{event.From.Hex(), *event.Amount}
		}
	}
}

type ExcChanal struct {
	Amount  big.Int
	Message string
}

func ExcListener(accumInstance *contract.Accum, outChanal chan ExcChanal) {
	excEventChanal := make(chan *contract.AccumExceed)

	opts := &bind.WatchOpts{
		Context: context.Background(),
	}

	sub, err := accumInstance.WatchExceed(opts, excEventChanal)
	if err != nil {
		log.Println("Ошибка подписки на события Exceed!", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Println("Ошибка чтения события Exceed!", err)
		case event := <-excEventChanal:
			outChanal <- ExcChanal{*event.Amount, event.Message}
		}
	}
}
