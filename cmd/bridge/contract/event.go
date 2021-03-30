package contract

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"strings"

	store "github.com/VoroshilovMax/Bettery/cmd/bridge/contract/abi"
	"github.com/VoroshilovMax/Bettery/x/coinmaker"
	coinmakerTypes "github.com/VoroshilovMax/Bettery/x/coinmaker/types"
	sdkContext "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/tendermint/go-amino"
)

func ConnectToEvent(infuraKey string, bridgeAddress string, cdc *codec.Codec, inBuf io.Reader) error {
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

			event := struct {
				Sender common.Address
				Amount *big.Int
				Cosmos string
			}{}
			err := contractAbi.Unpack(&event, "Deposit", vLog.Data)
			if err != nil {
				return errors.Errorf("error unpack event: ", err)
			}

			owner, err := sdk.AccAddressFromBech32(event.Cosmos)
			if err != nil {
				fmt.Println("Failed to parse owner address")
			} else {

				var msgValue = coinmakerTypes.CreateCoin{
					Symbol:        "eth",
					Owner:         owner,
					Amount:        uint(event.Amount.Uint64()),
					SenderAddress: string(event.Sender.Hex()),
				}

				fmt.Println(msgValue)

				msg := coinmaker.MsgCreateCoin(msgValue)

				msg.ValidateBasic()
				if err != nil {
					fmt.Println(err)
					return err
				} else {
					// TO DO "bettery" get from types
					moniker := "validator"
					chainID := "bettery"
					validator := "cosmosvaloper1xwlvmc4nq4d9tdst69m5ymef67tlpad25rdy64"
					rpcURL := "http://localhost:26657"
					validatorName := "validator"
					validatorAddress, err := sdk.ValAddressFromBech32(validator)
					// validatorAddress, validatorName, err := LoadValidatorCredentials(validator, inBuf)
					if err != nil {
						fmt.Println("from ValAddressFromBech32")
						fmt.Println(err)
					}

					cliCtx := LoadTendermintCLIContext(cdc, validatorAddress, validatorName, rpcURL, chainID)
					txBldr := authtypes.NewTxBuilderFromCLI(nil).
						WithTxEncoder(utils.GetTxEncoder(cdc)).
						WithChainID(chainID)
					// Prepare tx
					txBldr, err = utils.PrepareTxBuilder(txBldr, cliCtx)
					if err != nil {
						fmt.Println("from PrepareTxBuilder")
						fmt.Println(err)
					}

					// Build and sign the transaction
					txBytes, err := txBldr.BuildAndSign(moniker, keys.DefaultKeyPass, []sdk.Msg{msg})
					if err != nil {
						fmt.Println("from BuildAndSign")
						fmt.Println(err)
					}

					// Broadcast to a Tendermint node
					res, err := cliCtx.BroadcastTxSync(txBytes)
					if err != nil {
						fmt.Println("from BroadcastTxSync")
						fmt.Println(err)
					}

					if err = cliCtx.PrintOutput(res); err != nil {
						fmt.Println("from PrintOutput")
						fmt.Println(err)
					}
				}
			}
		}
	}
}

// LoadTendermintCLIContext : loads CLI context for tendermint txs
func LoadTendermintCLIContext(appCodec *amino.Codec, validatorAddress sdk.ValAddress, validatorName string,
	rpcURL string, chainID string) sdkContext.CLIContext {
	// Create the new CLI context
	cliCtx := sdkContext.NewCLIContext().
		WithCodec(appCodec).
		WithFromAddress(sdk.AccAddress(validatorAddress)).
		WithFromName(validatorName)

	if rpcURL != "" {
		cliCtx = cliCtx.WithNodeURI(rpcURL)
	}
	cliCtx.SkipConfirm = true

	// Confirm that the validator's address exists
	accountRetriever := authtypes.NewAccountRetriever(cliCtx)
	err := accountRetriever.EnsureExists((sdk.AccAddress(validatorAddress)))
	if err != nil {
		fmt.Println("from LoadTendermintCLIContext")
		fmt.Println(err)
	}
	return cliCtx
}

// LoadValidatorCredentials : loads validator's credentials (address, moniker, and passphrase)
func LoadValidatorCredentials(validatorFrom string, inBuf io.Reader) (sdk.ValAddress, string, error) {
	// Get the validator's name and account address using their moniker
	validatorAccAddress, validatorName, err := sdkContext.GetFromFields(inBuf, validatorFrom, false)
	if err != nil {
		fmt.Println("PROBLEM 1")
		return sdk.ValAddress{}, "", err
	}
	validatorAddress := sdk.ValAddress(validatorAccAddress)

	// Confirm that the key is valid
	_, err = authtypes.MakeSignature(nil, validatorName, keys.DefaultKeyPass, authtypes.StdSignMsg{})
	if err != nil {
		fmt.Println("PROBLEM 2")
		return sdk.ValAddress{}, "", err
	}

	return validatorAddress, validatorName, nil
}
