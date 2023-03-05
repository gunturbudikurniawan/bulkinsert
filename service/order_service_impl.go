package service

import (
	"bpjs/helper"
	"bpjs/model/domain"
	"bpjs/model/response"
	"bpjs/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/sourcegraph/conc/iter"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewCategoryService(orderRepository repository.OrderRepository, DB *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *OrderServiceImpl) Create(ctx context.Context, request []response.OrderCreateRequest) []response.OrderResponse {
	var orders []domain.Orders
	// err := service.Validate.Struct(request)
	// helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	for _, value := range request {
		order := domain.Orders{
			Customer: value.Customer,
			Quantity: value.Quantity,
			Price:    value.Price,
		}
		orders = append(orders, order)
	}

	orderlist := service.Routing(orders, ctx, tx)

	return helper.ToOrderResponses(orderlist)

}
func (service *OrderServiceImpl) Routing(orders []domain.Orders, ctx context.Context, tx *sql.Tx) []domain.Orders {
	return iter.Map(orders, func(o *domain.Orders) domain.Orders {
		return service.OrderRepository.Save(ctx, tx, *o)
	})
}

func (service *OrderServiceImpl) FindAll(ctx context.Context) []response.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.OrderRepository.FindAll(ctx, tx)

	return helper.ToOrderResponses(categories)
}
