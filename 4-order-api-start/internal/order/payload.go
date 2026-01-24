package order

type OrderCreateRequest struct {
	Products []uint `json:"product_ids" validate:"required,min=1,dive,gt=0"`
}
