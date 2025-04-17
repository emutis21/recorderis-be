package drivens // Asegúrate que el paquete sea 'drivens'

import (
	"recorderis/cmd/services/auth/ports/drivens"

	"golang.org/x/crypto/bcrypt"
)

var _ drivens.ForPasswordManager = (*BcryptAdapter)(nil)

type BcryptAdapter struct {
	cost int
}

func NewBcryptAdapter() *BcryptAdapter {
	return &BcryptAdapter{
		cost: bcrypt.DefaultCost,
	}
}

func (b *BcryptAdapter) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (b *BcryptAdapter) ValidatePassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
