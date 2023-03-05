package service

import (
	"bpjs/model/response"
	"context"
)

type OrderService interface {
	Create(ctx context.Context, response []response.OrderCreateRequest) []response.OrderResponse
	FindAll(ctx context.Context) []response.OrderResponse
}
