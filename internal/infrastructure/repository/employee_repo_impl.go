package repository

import (
	"ka-auth-service/internal/domain/entity"

	"github.com/jmoiron/sqlx"
)

type EmployeeRepoImpl struct {
	db *sqlx.DB
}

func NewEmployeeRepoImpl(db *sqlx.DB) *EmployeeRepoImpl {
	return &EmployeeRepoImpl{db: db}
}

func (r *EmployeeRepoImpl) GetEmployeeIDByEmail(email string) (*entity.Employee, error) {
	employee := &entity.Employee{}
	query := "SELECT id, email FROM public.users WHERE email = $1 AND is_active = true"
	err := r.db.Get(employee, query, email)
	return employee, err
}
