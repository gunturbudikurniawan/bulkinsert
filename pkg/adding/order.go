package adding

import (
	"time"
)

type Order struct {
	ID        int       `json:"id"`
	Customer  string    `json:"customer"`
	Quantity  int       `json:"last_login_on"`
	Price     float32   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}
