package controller

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/splitwise/models"
	"strconv"
)

func (h *HandlerStruct) PatchTransactionHandler(ctx iris.Context) {
	transactionId := ctx.Params().Get("transactionId")
	var req models.PatchTransactionRequest
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(transactionId)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	err = h.ServiceStruct.UpdatePaymentStatus(id, req)
	if err != nil {
		golog.Error(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON("success")
}
