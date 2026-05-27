package controllers

import (
	"net/http"

	"anime-streaming-platform/services"
	"anime-streaming-platform/utils"

	"github.com/gin-gonic/gin"
)

// DashboardController handles dashboard statistics
type DashboardController struct {
	animeService *services.AnimeService
	userService  *services.UserService
}

// NewDashboardController creates a new dashboard controller
func NewDashboardController() *DashboardController {
	return &DashboardController{
		animeService: services.NewAnimeService(),
		userService:  services.NewUserService(),
	}
}

// GetStats handles GET /api/dashboard/stats
func (ctrl *DashboardController) GetStats(c *gin.Context) {
	ctx := c.Request.Context()

	// Get anime stats
	latestAnimes, err := ctrl.animeService.GetLatestAnimes(ctx, 5)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Get user stats
	userStats, err := ctrl.userService.GetStats(ctx)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	stats := map[string]interface{}{
		"latestAnimes": latestAnimes,
		"userStats":    userStats,
	}

	utils.SuccessResponse(c, http.StatusOK, "Dashboard statistics retrieved", stats)
}
