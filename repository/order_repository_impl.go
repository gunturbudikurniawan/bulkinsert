package repository

import (
	"bpjs/config"
	"bpjs/helper"
	"bpjs/model/domain"
	"context"
	"database/sql"
	"time"
)

type OrderRepositoryImpl struct {
}

func NewCategoryRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, order domain.Orders) domain.Orders {
	SQL := "insert into orders(customer,quantity,price,timestamps) values (?, ?, ?, ?)"
	now := time.Now()
	result, err := tx.ExecContext(ctx, SQL, order.Customer, order.Quantity, order.Price, now.Format(config.AppTLayout))
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	order.Id = int(id)
	order.Timestamps = now
	return order
}

func (repository *OrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Orders {
	SQL := "select id, customer, quantity, price, timestamps  from orders"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orders []domain.Orders
	for rows.Next() {
		order := domain.Orders{}
		err := rows.Scan(&order.Id, &order.Customer, &order.Quantity, &order.Price, &order.Timestamps)
		helper.PanicIfError(err)
		orders = append(orders, order)
	}
	return orders
}
