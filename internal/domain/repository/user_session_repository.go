package repository

type UserSessionRepository interface {
	CountUserSession(userID int) (int, error)
	CountUserSessionByStatus(userID int, status string) (int, error)
	Insert(userID int, sessionToken string) error
	UpdateLastAccessed(userID int) error
	Delete(userID int) error
}
