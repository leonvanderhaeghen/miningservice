package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateMine{}

type MsgCreateMine struct {
  ID      string
  Name string `json:"name" yaml:"name"`
  Price sdk.Coins `json:"price" yaml:"price"`
  Owner sdk.AccAddress `json:"owner" yaml:"owner"`
  PlayerID      string `json:"playerid" yaml:"playerid"`
  Selling bool `json:"selling" yaml:"selling"`
  Efficiency int `json:"efficiency" yaml:"efficiency"`
  Resources []string `json:"resources" yaml:"resources"`
  UraniumCost int `json:"uraniumCost" yaml:"uraniumCost"`
  ResourceCounter int `json:"ResourceCounter" yaml:"ResourceCounter"`
}

func NewMsgCreateMine(owner sdk.AccAddress,playerID string, name string, price sdk.Coins, selling bool, efficiency int, resources []string, uraniumCost int) MsgCreateMine {
  return MsgCreateMine{
    ID: uuid.New().String(),
    Name: name,
    Price: price,
	Owner: owner,
	PlayerID: playerID,
    Selling: selling,
    Efficiency: efficiency,
    Resources: resources,
	UraniumCost: uraniumCost,
	}
}

func (msg MsgCreateMine) Route() string {
  return RouterKey
}

func (msg MsgCreateMine) Type() string {
  return "CreateMine"
}

func (msg MsgCreateMine) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgCreateMine) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateMine) ValidateBasic() error {
  if msg.Owner.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}