package controllers

import (
	"net/http"
	"strconv"

	"anime-streaming-platform/models"
	"anime-streaming-platform/services"
	"anime-streaming-platform/utils"

	"github.com/gin-gonic/gin"
)

// EpisodeController handles episode HTTP requests
type EpisodeController struct {
	service *services.EpisodeService
}

// NewEpisodeController creates a new episode controller
func NewEpisodeController() *EpisodeController {
	return &EpisodeController{
		service: services.NewEpisodeService(),
	}
}

// Create handles POST /api/episodes
func (ctrl *EpisodeController) Create(c *gin.Context) {
	var episode models.Episode
	if err := c.ShouldBindJSON(&episode); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	resp, err := ctrl.service.CreateEpisode(c.Request.Context(), &episode)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Episode created successfully", resp)
}

// GetByAnime handles GET /api/animes/:animeRef/episodes.
// animeRef accepts either a Mongo ObjectID or a public anime slug.
func (ctrl *EpisodeController) GetByAnime(c *gin.Context) {
	animeRef := c.Param("animeRef")
	if animeRef == "" {
		animeRef = c.Param("slug")
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	episodes, total, totalPages, err := ctrl.service.GetEpisodesByAnime(c.Request.Context(), animeRef, page, limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	meta := utils.Meta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}

	utils.PaginatedResponse(c, http.StatusOK, "Episodes retrieved successfully", episodes, meta)
}

// GetByAnimeAndNumber handles GET /api/animes/:animeRef/episodes/:number.
func (ctrl *EpisodeController) GetByAnimeAndNumber(c *gin.Context) {
	animeRef := c.Param("animeRef")
	if animeRef == "" {
		animeRef = c.Param("slug")
	}
	number, err := strconv.Atoi(c.Param("number"))
	if err != nil || number < 1 {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid episode number")
		return
	}

	episode, err := ctrl.service.GetEpisodeByAnimeAndNumber(c.Request.Context(), animeRef, number)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Episode retrieved successfully", episode)
}

// GetByID handles GET /api/episodes/:id
func (ctrl *EpisodeController) GetByID(c *gin.Context) {
	id := c.Param("id")
	episode, err := ctrl.service.GetEpisodeByID(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Episode retrieved successfully", episode)
}

// Update handles PUT /api/episodes/:id
func (ctrl *EpisodeController) Update(c *gin.Context) {
	id := c.Param("id")
	var episode models.Episode
	if err := c.ShouldBindJSON(&episode); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	resp, err := ctrl.service.UpdateEpisode(c.Request.Context(), id, &episode)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Episode updated successfully", resp)
}

// Delete handles DELETE /api/episodes/:id
func (ctrl *EpisodeController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.service.DeleteEpisode(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Episode deleted successfully", nil)
}

// GetLatest handles GET /api/episodes/latest
func (ctrl *EpisodeController) GetLatest(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	episodes, err := ctrl.service.GetLatestEpisodes(c.Request.Context(), limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Latest episodes retrieved", episodes)
}
