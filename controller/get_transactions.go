package controller

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
)

func (h *HandlerStruct) GetOweToTransactionHandler(ctx iris.Context){
	mobileNo := ctx.Params().Get("mobileNo")
	resp, err := h.ServiceStruct.GetOweToTransactions(mobileNo)
	if err != nil {
		golog.Error(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	_, _ = ctx.JSON(resp)
}

func (h *HandlerStruct) GetOweByTransactionHandler(ctx iris.Context){
	mobileNo := ctx.Params().Get("mobileNo")
	resp, err := h.ServiceStruct.GetOweByTransactions(mobileNo)
	if err != nil {
		golog.Error(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	_, _ = ctx.JSON(resp)
}
