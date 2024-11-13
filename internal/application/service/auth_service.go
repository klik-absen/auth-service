package service

import (
	"errors"
	"fmt"
	"ka-auth-service/internal/domain/entity"
	"ka-auth-service/internal/domain/repository"
)

type AuthService struct {
	userRepo        repository.UserRepository
	userSessionRepo repository.UserSessionRepository
}

func NewAuthService(userRepo repository.UserRepository, userSessionRepo repository.UserSessionRepository) *AuthService {
	return &AuthService{
		userRepo:        userRepo,
		userSessionRepo: userSessionRepo,
	}
}

func (s *AuthService) Authenticate(email, password string) (*entity.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)

	// validate user by email
	if err != nil {
		return nil, errors.New("user not found")
	}

	// validate password
	if user.Password != password {
		return nil, errors.New("incorrect password")
	}

	// validate active user
	if !user.IsActive {
		return nil, errors.New("your account has been disabled")
	}

	// updating last login
	if err := s.userRepo.UpdateLastLogin(user.Email); err != nil {
		return nil, errors.New(err.Error())
	}

	return user, nil
}

func (s *AuthService) CreateUserSession(userID int, sessionToken string) error {
	// count user session
	count, err := s.userSessionRepo.CountUserSession(userID)
	if err != nil {
		message := fmt.Sprintf("error count user session: %v", err.Error())
		return errors.New(message)
	}

	if count == 0 {
		// insert user session when not found
		if err := s.userSessionRepo.Insert(userID, sessionToken); err != nil {
			message := fmt.Sprintf("error insert user session: %v", err.Error())
			return errors.New(message)
		}
	} else {
		// count active user session
		countActive, err := s.userSessionRepo.CountUserSessionByStatus(userID, "Active")
		if err != nil {
			message := fmt.Sprintf("error count user session active: %v", err.Error())
			return errors.New(message)
		}

		if countActive == 0 {
			// delete user session when expired
			if err := s.userSessionRepo.Delete(userID); err != nil {
				message := fmt.Sprintf("error delete user session: %v", err.Error())
				return errors.New(message)
			}

			// then insert user session
			if err := s.userSessionRepo.Insert(userID, sessionToken); err != nil {
				message := fmt.Sprintf("error insert user session: %v", err.Error())
				return errors.New(message)
			}
		} else {
			// update last accessed when user session active
			if err := s.userSessionRepo.UpdateLastAccessed(userID); err != nil {
				message := fmt.Sprintf("error update last accessed user session: %v", err.Error())
				return errors.New(message)
			}
		}
	}

	return nil
}