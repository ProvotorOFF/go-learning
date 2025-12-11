package auth

import (
	"fmt"
	"log"
	"math/rand"
	"order-api-start/internal/session"
	"order-api-start/internal/user"
)

type AuthService struct {
	userRepo    *user.UserRepository
	sessionRepo *session.SessionRepository
}

func (service *AuthService) Login(user *user.User) (*session.Session, error) {
	user, err := service.userRepo.FindOrCreate(user)

	if err != nil {
		return nil, err
	}

	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	currentSession, err := service.sessionRepo.FindOrCreateByUserId(int(user.ID), code)

	if err != nil {
		return nil, err
	}

	log.Printf("Generated SMS code for user %d: %s\n", user.ID, code)

	return currentSession, nil
}
