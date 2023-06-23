package gapi2

import (
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/pb_account"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(account db.Account) *pb_account.Account {
	return &pb_account.Account{
		ID:       account.ID,
		Owner:    account.Owner,
		Balance:  int32(account.Balance),
		Currency: account.Currency,
		CreatAt:  timestamppb.New(account.CreatedAt),
	}
}
