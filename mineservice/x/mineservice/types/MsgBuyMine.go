package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgBuyMine{}

type MsgBuyMine struct {
  ID      string `json:"id" yaml:"id"`
  Buyer sdk.AccAddress `json:"Buyer" yaml:"Buyer"`
  Bid sdk.Coins `json:"price" yaml:"price"`
}

func NewMsgBuyMine(id string,buyer sdk.AccAddress, bid sdk.Coins) MsgBuyMine {
  return MsgBuyMine{
    ID: uuid.New().String(),
	Buyer: buyer,
    Bid: bid,
	}
}

func (msg MsgBuyMine) Route() string {
  return RouterKey
}

func (msg MsgBuyMine) Type() string {
  return "BuyMine"
}

func (msg MsgBuyMine) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Buyer)}
}

func (msg MsgBuyMine) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgBuyMine) ValidateBasic() error {
  if msg.Buyer.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  if len(msg.ID) == 0 {
	  return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,"id cannot be empty")
  }
  if !msg.Bid.IsAllPositive() {
	      return sdkerrors.ErrInsufficientFunds
  }
  return nil
}