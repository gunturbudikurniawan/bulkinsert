package helper

import (
	"bpjs/config"
	"bpjs/model/domain"
	"bpjs/model/response"
)

func ToOrderResponse(order domain.Orders) response.OrderResponse {
	return response.OrderResponse{
		Id:         order.Id,
		Customer:   order.Customer,
		Quantity:   order.Quantity,
		Price:      order.Price,
		Timestamps: order.Timestamps.Format(config.AppTLayout),
	}
}

func ToOrderResponses(categories []domain.Orders) []response.OrderResponse {
	var categoryResponses []response.OrderResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToOrderResponse(category))
	}
	return categoryResponses
}
