package domain

import "time"

type Orders struct {
	Id         int
	Customer   string
	Quantity   int
	Price      float32
	Timestamps time.Time
}
