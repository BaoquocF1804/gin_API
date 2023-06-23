package gapi2

import (
	"context"
	"github.com/techschool/simplebank/pb_account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetAccount(ctx context.Context, req *pb_account.GetAccountRequest) (*pb_account.GetAccountResponse, error) {
	arg, err := server.store.GetAccount(ctx, req.GetID())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no account")
	}
	rsp := &pb_account.GetAccountResponse{
		Account: convertUser(arg),
	}

	return rsp, nil
}
