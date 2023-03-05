package response

type OrderCreateRequest struct {
	Customer string  `validate:"required,min=1,max=100" json:"customer"`
	Quantity int     `validate:"required" json:"quantity"`
	Price    float32 `validate:"required" json:"price"`
}
