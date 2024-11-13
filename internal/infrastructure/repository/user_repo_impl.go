package repository

import (
	"ka-auth-service/internal/domain/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepoImpl struct {
	db *sqlx.DB
}

func NewUserRepoImpl(db *sqlx.DB) *UserRepoImpl {
	return &UserRepoImpl{db: db}
}

func (r *UserRepoImpl) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	query := "SELECT id, email, password, role_id, is_active FROM public.users WHERE email = $1"
	err := r.db.Get(user, query, email)
	return user, err
}

func (r *UserRepoImpl) UpdateLastLogin(email string) error {
	query := "UPDATE public.users SET last_login = now() AT TIME ZONE 'Asia/Jakarta' WHERE email = $1"
	_, err := r.db.Exec(query, email)
	return err
}
