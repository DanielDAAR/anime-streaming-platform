package controllers

import (
	"net/http"
	"strconv"

	"anime-streaming-platform/models"
	"anime-streaming-platform/services"
	"anime-streaming-platform/utils"

	"github.com/gin-gonic/gin"
)

// UserController handles user HTTP requests
type UserController struct {
	service *services.UserService
}

// NewUserController creates a new user controller
func NewUserController() *UserController {
	return &UserController{
		service: services.NewUserService(),
	}
}

// GetAll handles GET /api/users
func (ctrl *UserController) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	users, total, totalPages, err := ctrl.service.GetAllUsers(c.Request.Context(), page, limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	meta := utils.Meta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}

	utils.PaginatedResponse(c, http.StatusOK, "Users retrieved successfully", users, meta)
}

// GetByID handles GET /api/users/:id
func (ctrl *UserController) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := ctrl.service.GetUserByID(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}

// UpdateRole handles PUT /api/users/:id/role
func (ctrl *UserController) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Role models.UserRole `json:"role" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	if err := ctrl.service.UpdateUserRole(c.Request.Context(), id, req.Role); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User role updated successfully", nil)
}

// ToggleActive handles PUT /api/users/:id/toggle-active
func (ctrl *UserController) ToggleActive(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.service.ToggleUserActive(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User status toggled successfully", nil)
}

// GetStats handles GET /api/users/stats
func (ctrl *UserController) GetStats(c *gin.Context) {
	stats, err := ctrl.service.GetStats(c.Request.Context())
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User statistics retrieved", stats)
}
