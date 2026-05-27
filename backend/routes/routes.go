package routes

import (
	"anime-streaming-platform/controllers"
	"anime-streaming-platform/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine) {
	// API v1 group
	api := router.Group("/api")
	{
		setupAuthRoutes(api)
		setupAnimeRoutes(api)
		setupEpisodeRoutes(api)
		setupCommentRoutes(api)
		setupUserRoutes(api)
		setupDashboardRoutes(api)
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "anime-streaming-api"})
	})
}

func setupAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	authController := controllers.NewAuthController()
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.GET("/me", middleware.AuthMiddleware(), authController.Me)
	}
}

func setupAnimeRoutes(api *gin.RouterGroup) {
	anime := api.Group("/animes")
	animeController := controllers.NewAnimeController()
	{
		// Public routes
		anime.GET("", animeController.GetAll)
		anime.GET("/latest", animeController.GetLatest)
		anime.GET("/top-rated", animeController.GetTopRated)
		anime.GET("/:slug", animeController.GetBySlug)

		// Admin routes
		admin := anime.Group("")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			admin.POST("", animeController.Create)
			admin.PUT("/:id", animeController.Update)
			admin.DELETE("/:id", animeController.Delete)
		}
	}
}

func setupEpisodeRoutes(api *gin.RouterGroup) {
	episodeController := controllers.NewEpisodeController()

	// Public routes
	api.GET("/animes/:slug/episodes", episodeController.GetByAnime)
	api.GET("/animes/:slug/episodes/:number", episodeController.GetByAnimeAndNumber)
	api.GET("/episodes/latest", episodeController.GetLatest)
	api.GET("/episodes/:id", episodeController.GetByID)

	// Admin routes
	admin := api.Group("/episodes")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("", episodeController.Create)
		admin.PUT("/:id", episodeController.Update)
		admin.DELETE("/:id", episodeController.Delete)
	}
}

func setupCommentRoutes(api *gin.RouterGroup) {
	comment := api.Group("/comments")
	commentController := controllers.NewCommentController()
	{
		// Public routes (with optional auth for likes)
		comment.GET("/:animeId", commentController.GetByAnimeID)

		// Authenticated routes
		auth := comment.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("", commentController.Create)
			auth.POST("/:id/reply", commentController.CreateReply)
			auth.POST("/:id/like", commentController.Like)
		}

		// Admin routes
		admin := comment.Group("")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			admin.GET("", commentController.GetRecent)
			admin.DELETE("/:id", commentController.Delete)
		}
	}
}

func setupUserRoutes(api *gin.RouterGroup) {
	user := api.Group("/users")
	userController := controllers.NewUserController()
	{
		// Admin routes
		user.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		user.GET("", userController.GetAll)
		user.GET("/stats", userController.GetStats)
		user.GET("/:id", userController.GetByID)
		user.PUT("/:id/role", userController.UpdateRole)
		user.PUT("/:id/toggle-active", userController.ToggleActive)
	}
}

func setupDashboardRoutes(api *gin.RouterGroup) {
	dashboard := api.Group("/dashboard")
	dashboardController := controllers.NewDashboardController()
	{
		dashboard.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		dashboard.GET("/stats", dashboardController.GetStats)
	}
}
