package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/VoroshilovMax/Bettery/app"
	contract "github.com/VoroshilovMax/Bettery/cmd/bridge/contract"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
)

var cdc *codec.Codec

const (
	EnvPrefix  = "BRIDGE"
	FlagRPCURL = "rpc-url"
)

func init() {

	// Read in the configuration file for the sdk
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	cdc = app.MakeCodec()

	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentFlags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend,
		"Select keyring's backend (os|file|test)")
	rootCmd.PersistentFlags().String(flags.FlagChainID, "", "Chain ID of tendermint node")
	rootCmd.PersistentFlags().String(FlagRPCURL, "", "RPC URL of tendermint node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}

	// Construct Root Command
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		initBridgeCmd(),
	)

	DefaultCLIHome := os.ExpandEnv("$HOME/.batterycli")
	executor := cli.PrepareMainCmd(rootCmd, EnvPrefix, os.ExpandEnv(DefaultCLIHome))
	err := executor.Execute()
	if err != nil {
		log.Fatal("failed executing CLI command", err)
	}
}

var rootCmd = &cobra.Command{
	Use:          "bridge",
	Short:        "Streams live events from Ethereum and Cosmos and relays event information to the opposite chain",
	SilenceUsage: true,
}

func initBridgeCmd() *cobra.Command {
	initBridgeCmd := &cobra.Command{
		Use:     "init [infuraKey] [bridgeAddress]",
		Short:   "Run my bridge",
		Args:    cobra.ExactArgs(2),
		Example: "bridge init e15499b20fd64706bedea8339f77755d 0x2Cf2D6A4B92fF17c85273d847A5A4D4dBb4Cb943",
		RunE:    RunInitBridgeCmdCmd,
	}

	return initBridgeCmd
}

func RunInitBridgeCmdCmd(cmd *cobra.Command, args []string) error {
	// Validate and parse arguments
	if len(strings.Trim(args[0], "")) == 0 {
		return errors.Errorf("invalid [infuraKey]: %s", args[0])
	}
	infuraKey := args[0]

	// Validate and parse arguments
	if len(strings.Trim(args[1], "")) == 0 {
		return errors.Errorf("invalid [infuraKey]: %s", args[1])
	}
	bridgeAddress := args[1]

	inBuf := bufio.NewReader(cmd.InOrStdin())

	contract.ConnectToEvent(infuraKey, bridgeAddress, cdc, inBuf)
	return nil
}

func initConfig(cmd *cobra.Command) error {
	return viper.BindPFlag(flags.FlagChainID, cmd.PersistentFlags().Lookup(flags.FlagChainID))
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
