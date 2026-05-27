package controllers

import (
	"net/http"
	"strconv"

	"anime-streaming-platform/services"
	"anime-streaming-platform/utils"

    "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CommentController handles comment HTTP requests
type CommentController struct {
	service *services.CommentService
}

// NewCommentController creates a new comment controller
func NewCommentController() *CommentController {
	return &CommentController{
		service: services.NewCommentService(),
	}
}

// GetByAnimeID handles GET /api/comments/:animeId
func (ctrl *CommentController) GetByAnimeID(c *gin.Context) {
	animeID := c.Param("animeId")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	comments, total, totalPages, err := ctrl.service.GetCommentsByAnimeID(c.Request.Context(), animeID, page, limit)
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

	utils.PaginatedResponse(c, http.StatusOK, "Comments retrieved successfully", comments, meta)
}

// Create handles POST /api/comments
func (ctrl *CommentController) Create(c *gin.Context) {
	var req struct {
		AnimeID string `json:"animeId" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	userID, _ := c.Get("userId")
	animeObjID, err := primitive.ObjectIDFromHex(req.AnimeID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid anime ID")
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	comment, err := ctrl.service.CreateComment(c.Request.Context(), animeObjID, userObjID, req.Content)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Comment created successfully", comment)
}

// CreateReply handles POST /api/comments/:id/reply
func (ctrl *CommentController) CreateReply(c *gin.Context) {
	commentID := c.Param("id")
	var req struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	userID, _ := c.Get("userId")
	parentObjID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid comment ID")
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	reply, err := ctrl.service.CreateReply(c.Request.Context(), parentObjID, userObjID, req.Content)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Reply created successfully", reply)
}

// Like handles POST /api/comments/:id/like
func (ctrl *CommentController) Like(c *gin.Context) {
	commentID := c.Param("id")
	userID, _ := c.Get("userId")

	commentObjID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid comment ID")
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := ctrl.service.LikeComment(c.Request.Context(), commentObjID, userObjID); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Comment liked successfully", nil)
}

// GetRecent handles GET /api/comments (admin)
func (ctrl *CommentController) GetRecent(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	comments, total, totalPages, err := ctrl.service.GetRecentComments(c.Request.Context(), page, limit)
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

	utils.PaginatedResponse(c, http.StatusOK, "Recent comments retrieved", comments, meta)
}

// Delete handles DELETE /api/comments/:id (admin)
func (ctrl *CommentController) Delete(c *gin.Context) {
	commentID := c.Param("id")
	commentObjID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid comment ID")
		return
	}

	if err := ctrl.service.DeleteComment(c.Request.Context(), commentObjID); err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Comment deleted successfully", nil)
}
