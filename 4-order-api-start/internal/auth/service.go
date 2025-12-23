package auth

import (
	"fmt"
	"log"
	"math/rand"
	"order-api-start/configs"
	"order-api-start/internal/session"
	"order-api-start/internal/user"
	"order-api-start/pkg/jwt"
)

type AuthService struct {
	userRepo    *user.UserRepository
	sessionRepo *session.SessionRepository
	config      *configs.Config
}

type ServiceDeps struct {
	UserRepo    *user.UserRepository
	SessionRepo *session.SessionRepository
	Config      *configs.Config
}

func NewAuthService(deps ServiceDeps) *AuthService {
	return &AuthService{
		sessionRepo: deps.SessionRepo,
		userRepo:    deps.UserRepo,
		config:      deps.Config,
	}
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

func (service *AuthService) Verify(session *session.Session) (string, error) {
	verifiedSession, err := service.sessionRepo.Verify(session.SID, session.Code)

	if err != nil {
		return "", err
	}

	user, err := service.userRepo.FindById(int(verifiedSession.UserID))

	if err != nil {
		return "", err
	}

	token, err := jwt.NewJWT(service.config.Secret).Create(user.Phone)

	if err != nil {
		return "", err
	}

	return token, nil
}
