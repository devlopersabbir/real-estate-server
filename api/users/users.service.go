package users

import (
	"net/http"

	"github.com/devlopersabbir/juan_don82-server/api/users/core"
	"github.com/devlopersabbir/juan_don82-server/api/users/domain"
	cf "github.com/devlopersabbir/juan_don82-server/internal/pkg/config"
	"github.com/devlopersabbir/juan_don82-server/internal/pkg/utils"
	v "github.com/devlopersabbir/juan_don82-server/internal/pkg/validator"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body domain.CreateUserRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Validate struct fields
	if errs := v.Validate(body); errs != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := &core.Users{
		Name:     body.Name,
		Email:    body.Email,
		Password: hashedPassword,
		Role:     "user", // default role
	}

	// Store pg database
	if err := Store(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	// Store elastic database
	if err := StoreElastic(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func LoginUser(c *gin.Context) {
	var body domain.LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if errs := v.Validate(body); errs != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}

	user, err := FindByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(body.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	secret := cf.GetEnv("JWT_SECRET", "supersecretkey")
	refreshSecret := cf.GetEnv("JWT_REFRESH_SECRET", "superrefreshsecretkey")

	accessToken, refreshToken, err := utils.GenerateTokens(user.ID, user.Email, user.Role, secret, refreshSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	c.JSON(http.StatusOK, domain.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func GetUsers(c *gin.Context) {
	// We can implement fetching all users if needed
	c.JSON(http.StatusOK, gin.H{
		"message": "Users fetched successfully",
	})
}

func RefreshUserToken(c *gin.Context) {
	var body domain.RefreshRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if errs := v.Validate(body); errs != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}

	refreshSecret := cf.GetEnv("JWT_REFRESH_SECRET", "superrefreshsecretkey")
	claims, err := utils.VerifyToken(body.RefreshToken, refreshSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	secret := cf.GetEnv("JWT_SECRET", "supersecretkey")

	accessToken, refreshToken, err := utils.GenerateTokens(claims.UserID, claims.Email, claims.Role, secret, refreshSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	c.JSON(http.StatusOK, domain.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
