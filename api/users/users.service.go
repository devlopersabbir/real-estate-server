package users

import (
	"github.com/devlopersabbir/juan_don82-server/api/users/core"
	"github.com/devlopersabbir/juan_don82-server/api/users/domain"
	"github.com/devlopersabbir/juan_don82-server/arch/networks"
	"github.com/devlopersabbir/juan_don82-server/internal/pkg/config"
	"github.com/devlopersabbir/juan_don82-server/internal/pkg/utils"
	v "github.com/devlopersabbir/juan_don82-server/internal/pkg/validator"
	"github.com/gin-gonic/gin"
)

// CreateUser handles user registration
//
//	@Summary		Register a new user
//	@Description	Creates a new user with name, email, and password
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.CreateUserRequest	true	"User Registration Details"
//	@Success		201		{object}	map[string]string			"User created successfully"
//	@Router			/api/v1/auth/register [post]
func CreateUser(c *gin.Context) {
	var body domain.CreateUserRequest
	res := networks.Send(c)

	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequestError("Invalid request body", err)
		return
	}

	// Validate struct fields
	if errs := v.Validate(body); errs != nil {
		res.ValidationError("Validation failed", errs)
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		res.InternalServerError("Failed to hash password", err)
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
		res.InternalServerError("Failed to create user", err)
		return
	}
	// Store elastic database
	if err := StoreElastic(c, user); err != nil {
		res.InternalServerError("Failed to create user in search index", err)
		return
	}

	res.SuccessMsgResponse("User created successfully")
}

// LoginUser handles user login
//
//	@Summary		Log in a user
//	@Description	Authenticates a user and returns access and refresh
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.LoginRequest	true	"Login Credentials"
//	@Success		200		{object}	domain.AuthResponse	"Tokens"
//	@Router			/api/v1/auth/login [post]
func LoginUser(c *gin.Context) {
	var body domain.LoginRequest
	res := networks.Send(c)

	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequestError("Invalid request body", err)
		return
	}

	if errs := v.Validate(body); errs != nil {
		res.ValidationError("Validation failed", errs)
		return
	}

	user, err := FindByEmail(body.Email)
	if err != nil || user.Email == "" {
		res.UnauthorizedError("Invalid credentials", err)
		return
	}

	if !utils.CheckPasswordHash(body.Password, user.Password) {
		res.UnauthorizedError("Incorrect password", nil)
		return
	}

	env, _ := config.LoadEnv()
	accessToken, refreshToken, err := utils.GenerateTokens(user.ID, user.Email, user.Role, env.JWTConfig.Secret, env.JWTConfig.RefreshSecret)
	if err != nil {
		res.InternalServerError("Failed to generate tokens", err)
		return
	}

	res.SuccessDataResponse("Login successful", domain.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// GetUsers fetches all users
//
//	@Summary		Get all users
//	@Description	Fetches a list of all users. Requires authentication.
//	@Tags			Users
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Users fetched successfully"
//	@Router			/api/v1/users [get]
func GetUsers(c *gin.Context) {
	// We can implement fetching all users if needed
	networks.Send(c).SuccessMsgResponse("Users fetched successfully")
}

// RefreshUserToken refreshes the access token
//
//	@Summary		Refresh authentication token
//	@Description	Provides a new access token using a refresh token
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.RefreshRequest	true	"Refresh Token"
//	@Success		200		{object}	domain.AuthResponse		"New tokens"
//	@Router			/api/v1/auth/refresh [post]
func RefreshUserToken(c *gin.Context) {
	var body domain.RefreshRequest
	res := networks.Send(c)

	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequestError("Invalid request body", err)
		return
	}

	if errs := v.Validate(body); errs != nil {
		res.ValidationError("Validation failed", errs)
		return
	}
	env, _ := config.LoadEnv()
	claims, err := utils.VerifyToken(body.RefreshToken, env.JWTConfig.RefreshSecret)
	if err != nil {
		res.UnauthorizedError("Invalid refresh token", err)
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(claims.UserID, claims.Email, claims.Role, env.JWTConfig.Secret, env.JWTConfig.RefreshSecret)
	if err != nil {
		res.InternalServerError("Failed to generate tokens", err)
		return
	}

	res.SuccessDataResponse("Token refreshed successfully", domain.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
