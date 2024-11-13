package repository

import (
	"github.com/jmoiron/sqlx"
)

type UserSessionRepoImpl struct {
	db *sqlx.DB
}

func NewUserSessionRepoImpl(db *sqlx.DB) *UserSessionRepoImpl {
	return &UserSessionRepoImpl{db: db}
}

func (r *UserSessionRepoImpl) CountUserSession(userID int) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM public.user_session WHERE user_id = $1"
	err := r.db.Get(&count, query, userID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *UserSessionRepoImpl) CountUserSessionByStatus(userID int, status string) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM public.user_session WHERE user_id = $1 AND status = $2"
	err := r.db.Get(&count, query, userID, status)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *UserSessionRepoImpl) Insert(userID int, sessionToken string) error {
	query := "INSERT INTO public.user_session (user_id, session_token, created_at, expires_at, status, last_accessed_at) VALUES ($1, $2, now() AT TIME ZONE 'Asia/Jakarta', now() AT TIME ZONE 'Asia/Jakarta' + INTERVAL '30 days', 'Active', now() AT TIME ZONE 'Asia/Jakarta')"
	_, err := r.db.Exec(query, userID, sessionToken)
	return err
}

func (r *UserSessionRepoImpl) UpdateLastAccessed(userID int) error {
	query := "UPDATE public.user_session SET last_accessed_at = now() AT TIME ZONE 'Asia/Jakarta' WHERE user_id = $1"
	_, err := r.db.Exec(query, userID)
	return err
}

func (r *UserSessionRepoImpl) UpdateStatusExpired(userID int) error {
	query := "UPDATE public.user_session SET status = 'Expired' WHERE user_id = $1"
	_, err := r.db.Exec(query, userID)
	return err
}

func (r *UserSessionRepoImpl) Delete(userID int) error {
	query := "DELETE FROM public.user_session WHERE user_id = $1"
	_, err := r.db.Exec(query, userID)
	return err
}
