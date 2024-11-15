package repository

import "ka-auth-service/internal/domain/entity"

type EmployeeRepository interface {
	GetEmployeeIDByEmail(email string) (*entity.Employee, error)
}
