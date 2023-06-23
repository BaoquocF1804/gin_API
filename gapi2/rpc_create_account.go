package gapi2

import (
	"context"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/pb_account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateAccount(ctx context.Context, req *pb_account.CreateAccountRequest) (*pb_account.CreateAccountResponse, error) {

	arg := db.CreateAccountParams{
		Owner:    req.GetOwner(),
		Balance:  req.GetBalance(),
		Currency: req.GetCurrency(),
	}
	err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account-----------: %s", err)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get account: %s", err)
	}

	//rsp := &pb_account.CreateAccountResponse{
	//	Account: convertUser(user),
	//}
	return nil, nil
}
