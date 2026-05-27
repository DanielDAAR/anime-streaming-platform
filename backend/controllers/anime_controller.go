package controllers

import (
	"net/http"
	"strconv"

	"anime-streaming-platform/models"
	"anime-streaming-platform/services"
	"anime-streaming-platform/utils"

	"github.com/gin-gonic/gin"
)

// AnimeController handles anime HTTP requests
type AnimeController struct {
	service *services.AnimeService
}

// NewAnimeController creates a new anime controller
func NewAnimeController() *AnimeController {
	return &AnimeController{
		service: services.NewAnimeService(),
	}
}

// Create handles POST /api/animes
func (ctrl *AnimeController) Create(c *gin.Context) {
	var anime models.Anime
	if err := c.ShouldBindJSON(&anime); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	resp, err := ctrl.service.CreateAnime(c.Request.Context(), &anime)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Anime created successfully", resp)
}

// GetAll handles GET /api/animes
func (ctrl *AnimeController) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	filters := map[string]interface{}{
		"search": c.Query("search"),
		"genre":  c.Query("genre"),
		"status": c.Query("status"),
	}
	if year := c.Query("year"); year != "" {
		if y, err := strconv.Atoi(year); err == nil {
			filters["year"] = y
		}
	}

	animes, total, totalPages, err := ctrl.service.GetAllAnimes(c.Request.Context(), page, limit, filters)
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

	utils.PaginatedResponse(c, http.StatusOK, "Animes retrieved successfully", animes, meta)
}

// GetBySlug handles GET /api/animes/:slug
func (ctrl *AnimeController) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")
	anime, err := ctrl.service.GetAnimeBySlug(c.Request.Context(), slug)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Anime retrieved successfully", anime)
}

// Update handles PUT /api/animes/:id
func (ctrl *AnimeController) Update(c *gin.Context) {
	id := c.Param("id")
	var anime models.Anime
	if err := c.ShouldBindJSON(&anime); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	resp, err := ctrl.service.UpdateAnime(c.Request.Context(), id, &anime)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Anime updated successfully", resp)
}

// Delete handles DELETE /api/animes/:id
func (ctrl *AnimeController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.service.DeleteAnime(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Anime deleted successfully", nil)
}

// GetLatest handles GET /api/animes/latest
func (ctrl *AnimeController) GetLatest(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	animes, err := ctrl.service.GetLatestAnimes(c.Request.Context(), limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Latest animes retrieved", animes)
}

// GetTopRated handles GET /api/animes/top-rated
func (ctrl *AnimeController) GetTopRated(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	animes, err := ctrl.service.GetTopRatedAnimes(c.Request.Context(), limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Top rated animes retrieved", animes)
}
