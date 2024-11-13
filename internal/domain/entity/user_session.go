package entity

type UserSession struct {
	UserID         int    `db:"user_id" json:"user_id"`
	SessionToken   string `db:"session_token" json:"session_token"`
	CreatedAt      string `db:"created_at" json:"created_at"`
	ExpiresAt      string `db:"expires_at" json:"expires_at"`
	Status         string `db:"status" json:"status"`
	LastAccessedAt string `db:"last_accessed_at" json:"last_accessed_at"`
}
