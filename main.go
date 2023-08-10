package main

import (
	"github.com/agilistikmal/wallet-go/controller"
	"github.com/agilistikmal/wallet-go/database"
	"github.com/agilistikmal/wallet-go/handler"
	"github.com/agilistikmal/wallet-go/repository"
	"github.com/agilistikmal/wallet-go/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := database.New()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := httprouter.New()

	router.GET("/api/user", userController.FindAll)
	router.GET("/api/user/:userId", userController.FindById)
	router.POST("/api/user", userController.Create)
	router.PUT("/api/user/:userId", userController.Update)
	router.DELETE("/api/user/:userId", userController.Delete)

	router.PanicHandler = handler.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	println("* Running on localhost:8080")
}
