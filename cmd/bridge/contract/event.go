package contract

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	store "github.com/VoroshilovMax/Bettery/cmd/bridge/contract/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

func ConnectToEvent(infuraKey string, bridgeAddress string) error {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws/v3/" + infuraKey)
	if err != nil {
		return errors.Errorf("error with connetction to smart contract: ", err)
	}

	contractAddress := common.HexToAddress(bridgeAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}

	var ch = make(chan types.Log)
	ctx := context.Background()

	sub, err := client.SubscribeFilterLogs(ctx, query, ch)
	if err != nil {
		return errors.Errorf("error to subscribe to the contract event: ", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		return errors.Errorf("error read contract abi: ", err)
	}

	fmt.Println("Connection success")

	for {
		select {
		case err := <-sub.Err():
			return errors.Errorf("error subscribe: ", err)
		case vLog := <-ch:
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())

			event := struct {
				Sender common.Address
				Amount *big.Int
			}{}
			err := contractAbi.Unpack(&event, "Deposit", vLog.Data)
			if err != nil {
				return errors.Errorf("error unpack event: ", err)
			}

			fmt.Println(string(event.Sender.Hex()))
			fmt.Println(event.Amount.String())

			var topics [4]string
			for i := range vLog.Topics {
				topics[i] = vLog.Topics[i].Hex()
			}

			fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
		}
	}
}
