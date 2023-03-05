package controller

import (
	"bpjs/helper"
	"bpjs/model/response"
	"bpjs/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewCategoryController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *OrderControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var orderCreateRequest []response.OrderCreateRequest
	helper.ReadFromRequestBody(request, &orderCreateRequest)

	orderResponse := controller.OrderService.Create(request.Context(), orderCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ordersResponses := controller.OrderService.FindAll(request.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ordersResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
