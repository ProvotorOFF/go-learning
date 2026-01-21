package order

type OrderCreateRequest struct {
	UserID   uint   `json:"user_id" validate:"required,gt=0"`
	Products []uint `json:"product_ids" validate:"required,min=1,dive,gt=0"`
}
