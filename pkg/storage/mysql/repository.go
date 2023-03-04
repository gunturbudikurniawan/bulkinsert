package mysql

import (
	"bpjs/config"
	"bpjs/pkg/adding"
	"context"
	"database/sql"
	"time"

	"fmt"
	"log"

	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(cfgdb config.MySQL) (*Storage, error) {
	var err error

	s := new(Storage)

	s.db, err = sql.Open("nrmysql", cfgdb.DSN)
	if err != nil {
		return s, err
	}

	log.Println("MySQL connected......")

	return s, nil
}

func (s *Storage) CreateOrderlist(ctx context.Context, aws adding.Order) (int, error) {
	q := fmt.Sprintf("INSERT INTO order_list (customer, quantity, price, timestamp) VALUE ('%s', %d, '%v','%s')", aws.Customer, aws.Quantity, aws.Price, time.Now().In(config.App.TLoc).Format(config.AppTLayout))

	res, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
