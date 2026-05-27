package controllers

import (
	"net/http"

	"anime-streaming-platform/models"
	"anime-streaming-platform/services"
	"anime-streaming-platform/utils"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication HTTP requests
type AuthController struct {
	service *services.AuthService
}

// NewAuthController creates a new auth controller
func NewAuthController() *AuthController {
	return &AuthController{
		service: services.NewAuthService(),
	}
}

// Register handles POST /api/auth/register
func (ctrl *AuthController) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	resp, err := ctrl.service.Register(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User registered successfully", resp)
}

// Login handles POST /api/auth/login
func (ctrl *AuthController) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	resp, err := ctrl.service.Login(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Login successful", resp)
}

// Me handles GET /api/auth/me
func (ctrl *AuthController) Me(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	user, err := ctrl.service.GetCurrentUser(c.Request.Context(), userID.(string))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}
