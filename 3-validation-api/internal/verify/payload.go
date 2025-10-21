package verify

type SendMailRequest struct {
	Email string `json:"email" validate:"required,email"`
}
