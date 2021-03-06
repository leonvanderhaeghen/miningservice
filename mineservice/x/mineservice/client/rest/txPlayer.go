package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createPlayerRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Name string `json:"name"`
	
}

func createPlayerHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPlayerRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(baseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedName := req.Name
		
		

		msg := types.NewMsgCreatePlayer(
			creator,
			parsedName,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setPlayerRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ID 		string `json:"id"`
	Name string `json:"name"`
	Invetory string `json:"invetory"`
	Mines string `json:"mines"`
	
}

func setPlayerHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setPlayerRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(baseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedName := req.Name
		
		

		msg := types.NewMsgSetPlayer(
			creator, 
			req.ID,
			parsedName,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deletePlayerRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	
	ID 		string `json:"id"`
}

func deletePlayerHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deletePlayerRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(baseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgDeletePlayer(req.ID, creator)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
