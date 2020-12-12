package cli

import (
	"bufio"
	"strconv"
	
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
)

func GetCmdCreateMine(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-mine [name] [price] [selling] [efficiency] [invetory] [uraniumCost] [resources]",
		Short: "Creates a new mine",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0])
			argsPrice,_ := sdk.ParseCoins(args[1])
			argsSelling,_ := strconv.ParseBool(args[2])
			argsEfficiency,_ := strconv.Atoi(args[3])
			argsInvetory := string(args[4])
			argsUraniumCost,_ := strconv.Atoi(args[5])
			argsResources := args[6:len(args)]

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateMine(cliCtx.GetFromAddress(), string(argsName), argsPrice, argsSelling, argsEfficiency, argsInvetory, argsResources, argsUraniumCost)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetMine(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-mine [id]  [name] [price] [selling] [efficiency] [invetory] [uraniumCost] [resources]",
		Short: "Set a new mine",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsName := string(args[1])
			argsPrice,_ := sdk.ParseCoins(args[2])
			argsSelling,_ := strconv.ParseBool(args[3])
			argsEfficiency,_ := strconv.Atoi(args[4])
			argsInvetory := string(args[5])
			argsUraniumCost,_ := strconv.Atoi(args[6])
			argsResources := args[7:len(args)]
			
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetMine(cliCtx.GetFromAddress(), id, string(argsName), argsPrice, argsSelling, argsEfficiency, string(argsInvetory), argsResources, argsUraniumCost)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteMine(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-mine [id]",
		Short: "Delete a new mine by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
	
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteMine(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdBuyMine(cdc *codec.Codec) *cobra.Command{
return &cobra.Command{
		Use:   "buy-mine [id] [price]",
		Short: "buy a new mine",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
		argsId := string(args[0] )

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgBuyMine(argsId, cliCtx.GetFromAddress(),coins)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}