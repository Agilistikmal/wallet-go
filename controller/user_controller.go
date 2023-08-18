package controller

import (
	"github.com/agilistikmal/wallet-go/helper"
	"github.com/agilistikmal/wallet-go/model"
	"github.com/agilistikmal/wallet-go/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
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
	userCreateRequest := model.UserCreateRequest{}
	helper.ReadFromRequest(r, &userCreateRequest)

	secretKey := r.Header.Get("Secret-Key")
	var webResponse model.WebResponse

	if os.Getenv("SECRET_KEY") == secretKey {
		userResponse := controller.UserService.Create(r.Context(), userCreateRequest)
		webResponse = model.WebResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   userResponse,
		}
	} else {
		webResponse = model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Data:   nil,
		}
	}
	helper.WriteToResponse(w, webResponse)
}

func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userUpdateRequest := model.UserUpdateRequest{}
	helper.ReadFromRequest(r, &userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(err)
	}

	userUpdateRequest.Id = uint(id)
	userResponse := controller.UserService.FindById(r.Context(), uint(id))
	secretKey := r.Header.Get("Secret-Key")
	apiKey := r.Header.Get("API-Key")
	var webResponse model.WebResponse

	if os.Getenv("SECRET_KEY") == secretKey || userResponse.ApiKey == apiKey {
		userResponse = controller.UserService.Update(r.Context(), userUpdateRequest)
		webResponse = model.WebResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   userResponse,
		}
	} else {
		webResponse = model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Data:   nil,
		}
	}
	helper.WriteToResponse(w, webResponse)
}

func (controller *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(err)
	}

	userResponse := controller.UserService.FindById(r.Context(), uint(id))
	secretKey := r.Header.Get("Secret-Key")
	apiKey := r.Header.Get("API-Key")
	var webResponse model.WebResponse

	if os.Getenv("SECRET_KEY") == secretKey || userResponse.ApiKey == apiKey {
		controller.UserService.Delete(r.Context(), uint(id))
		webResponse = model.WebResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   nil,
		}
	} else {
		webResponse = model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Data:   nil,
		}
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
	secretKey := r.Header.Get("Secret-Key")
	apiKey := r.Header.Get("API-Key")
	var webResponse model.WebResponse

	if os.Getenv("SECRET_KEY") == secretKey || userResponse.ApiKey == apiKey {
		webResponse = model.WebResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   userResponse,
		}
	} else {
		webResponse = model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Data:   nil,
		}
	}
	helper.WriteToResponse(w, webResponse)
}

func (controller *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userResponses := controller.UserService.FindAll(r.Context())
	secretKey := r.Header.Get("Secret-Key")
	var webResponse model.WebResponse

	if os.Getenv("SECRET_KEY") == secretKey {
		webResponse = model.WebResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   userResponses,
		}
	} else {
		webResponse = model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Data:   nil,
		}
	}
	helper.WriteToResponse(w, webResponse)
}
