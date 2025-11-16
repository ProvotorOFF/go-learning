package product

type ProductCreateRequest struct {
	Email string `json:"email" validate:"required,email"`
}
