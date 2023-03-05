package response

type OrderResponse struct {
	Id         int
	Customer   string
	Quantity   int
	Price      float32
	Timestamps string
}
