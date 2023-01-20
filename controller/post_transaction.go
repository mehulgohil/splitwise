package controller

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/splitwise/service"
)

type HandlerStruct struct {
	ServiceStruct service.ServiceStruct
}

func (h *HandlerStruct) PostTransactionHandler(ctx iris.Context) {
	err := h.ServiceStruct.PostTransactionToDB()
	if err != nil {
		golog.Error(err)
		return
	}
	_, _ = ctx.JSON("success")
}
