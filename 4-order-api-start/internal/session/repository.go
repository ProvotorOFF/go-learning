package session

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"order-api-start/pkg/db"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SessionRepository struct {
	Database *db.Db
}

func (repo *SessionRepository) FindOrCreateByUserId(id int, code string) (*Session, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	session := Session{}
	result := repo.Database.Where("user_id = ?", id).First(&session)
	SID, err := generateSID()

	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			session = Session{
				SID:       SID,
				UserID:    uint(id),
				Code:      string(hash),
				ExpiresAt: time.Now().Add(5 * time.Minute),
			}
			if err := repo.Database.Create(&session).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, result.Error
		}
	} else {
		session.Code = string(hash)
		session.ExpiresAt = time.Now().Add(5 * time.Minute)
		session.SID = SID

		if err := repo.Database.Save(&session).Error; err != nil {
			return nil, err
		}
	}

	return &session, nil
}

func generateSID() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
