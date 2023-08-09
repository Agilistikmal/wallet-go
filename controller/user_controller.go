package controller

import (
	"github.com/agilistikmal/wallet-go/helper"
	"github.com/agilistikmal/wallet-go/model"
	"github.com/agilistikmal/wallet-go/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

type UserControllerImpl struct {
	UserService service.UserService
}

func (controller *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userModelCreateRequest := model.UserModelCreateRequest{}
	helper.ReadFromRequest(r, &userModelCreateRequest)

	userResponse := controller.UserService.Create(r.Context(), userModelCreateRequest)
	webResponse := model.WebModelResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   userResponse,
	}

	helper.WriteToResponse(w, webResponse)
}

func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userModelUpdateRequest := model.UserModelUpdateRequest{}
	helper.ReadFromRequest(r, &userModelUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(err)
	}
	userModelUpdateRequest.Id = uint(id)
	userResponse := controller.UserService.Update(r.Context(), userModelUpdateRequest)
	webResponse := model.WebModelResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   userResponse,
	}

	helper.WriteToResponse(w, webResponse)
}

func (controller *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(err)
	}
	controller.UserService.Delete(r.Context(), uint(id))
	webResponse := model.WebModelResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}

	helper.WriteToResponse(w, webResponse)
}

func (controller *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(err)
	}
	userResponse := controller.UserService.FindById(r.Context(), uint(id))
	webResponse := model.WebModelResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   userResponse,
	}

	helper.WriteToResponse(w, webResponse)
}

func (controller *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userResponses := controller.UserService.FindAll(r.Context())
	webResponse := model.WebModelResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   userResponses,
	}

	helper.WriteToResponse(w, webResponse)
}
