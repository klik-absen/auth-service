package repository

import "ka-auth-service/internal/domain/entity"

type UserRepository interface {
	GetUserByEmail(email string) (*entity.User, error)
	UpdateLastLogin(email string) error
}
