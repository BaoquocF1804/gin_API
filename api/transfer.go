package api

//
//import (
//	"github.com/gin-gonic/gin"
//	db "github.com/techschool/simplebank/db/sqlc"
//	"net/http"
//)
//
//type transferRequest struct {
//	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
//	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
//	Amount        int64  `json:"amount" binding:"required,gt=0"`
//	Currency      string `json:"currency" binding:"required,currency"`
//}
//
//func (server *Server) createTransfer(ctx *gin.Context) {
//	var req transferRequest
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	arg := db.Transfer{
//		FromAccountID: req.FromAccountID,
//		ToAccountID:   req.FromAccountID,
//		Amount:        req.Amount,
//	}
//
//	result, err := server.store.GetTransfer(ctx, arg)
//	if !valid {
//		return
//	}
//
//	//server.store.CreateAccount(ctx, arg)
//	err := server.store.CreateAccount(ctx, arg)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//}
