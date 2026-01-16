package jwt

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(phone string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": phone,
	})

	s, err := t.SignedString([]byte(j.Secret))

	if err != nil {
		return "", err
	}

	return s, nil
}

func (j *JWT) Parse(token string) (bool, string) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, ""
	}
	phone := t.Claims.(jwt.MapClaims)["phone"]
	return t.Valid, phone.(string)
}
