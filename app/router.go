package app

import (
	"bpjs/controller"
	"bpjs/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(orderController controller.OrderController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/orders", orderController.Create)
	router.GET("/api/orders", orderController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
