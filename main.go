package main

import (
	"bpjs/app"
	"bpjs/controller"
	"bpjs/helper"
	"bpjs/middleware"
	"bpjs/repository"
	"bpjs/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	orderRepository := repository.NewCategoryRepository()
	orderService := service.NewCategoryService(orderRepository, db, validate)
	orderController := controller.NewCategoryController(orderService)
	router := app.NewRouter(orderController)

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
