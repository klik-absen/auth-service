package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"ka-auth-service/internal/application/service"
	"ka-auth-service/internal/domain/entity"
	"ka-auth-service/internal/interfaces/dto"
	"ka-auth-service/internal/interfaces/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Authentication(ctx *gin.Context) {
	var req dto.AuthRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Invalid request payload", nil))
		return
	}

	var user *entity.User
	user, err := c.authService.Authenticate(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	// generate session token
	sessionToken := generateToken(req.Email)

	// create user session
	if err := c.authService.CreateUserSession(user.ID, sessionToken); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	// create dto
	authResponse := dto.AuthResponse{
		Email: req.Email,
		Token: sessionToken,
	}

	ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Authentication successfully", authResponse))
}

func generateToken(email string) string {
	timestamp := time.Now().Unix()
	token := email + fmt.Sprintf("%d", timestamp)

	hash := sha256.New()
	hash.Write([]byte(token))
	hashedToken := hash.Sum(nil)
	hashString := hex.EncodeToString(hashedToken)

	return hashString
}
