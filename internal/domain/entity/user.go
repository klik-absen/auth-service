package entity

type User struct {
	ID       int    `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
	RoleID   int    `db:"role_id" json:"role_id"`
	IsActive bool   `db:"is_active" json:"is_active"`
}
