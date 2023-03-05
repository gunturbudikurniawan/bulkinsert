package repository

import (
	"bpjs/model/domain"
	"context"
	"database/sql"
)

type OrderRepository interface {
	Save(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Orders
}
