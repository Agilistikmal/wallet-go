package controller

import (
	"github.com/agilistikmal/wallet-go/helper"
	"github.com/agilistikmal/wallet-go/model"
	"github.com/agilistikmal/wallet-go/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type WalletController interface {
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type WalletControllerImpl struct {
	WalletService service.WalletService
}

func NewWalletController(walletService service.WalletService) WalletController {
	return &WalletControllerImpl{
		WalletService: walletService,
	}
}

func (controller *WalletControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	wallet := model.Wallet{}
	helper.ReadFromRequest(r, &wallet)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(err)
	}
	wallet.UserId = uint(id)

	walletResponse := controller.WalletService.Update(r.Context(), wallet)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   walletResponse,
	}

	helper.WriteToResponse(w, webResponse)
}
