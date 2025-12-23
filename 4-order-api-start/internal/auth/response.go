package auth

type LoginResponse struct {
	SID string `json:"session_id"`
}

type LoginVerifiedResponse struct {
	Token string `json:"token"`
}
