package controller

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/splitwise/models"
	"github.com/mehulgohil/splitwise/service"
)

type HandlerStruct struct {
	ServiceStruct service.ServiceStruct
}

func (h *HandlerStruct) PostTransactionHandler(ctx iris.Context) {
	var req models.TransactionModel
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	err = h.ServiceStruct.PostTransactionToDB(req)
	if err != nil {
		golog.Error(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	_, _ = ctx.JSON("success")
}
