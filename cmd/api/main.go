package main

import (
	"context"
	"fmt"
	"itv/internal/config"
	"itv/internal/controller"
	"itv/internal/middleware"
	"itv/internal/repository"
	"itv/internal/service"
	"itv/pkg/auth"
	"itv/pkg/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	_ "golang.org/x/crypto/bcrypt"
	_ "itv/docs"
)

// @title						Movies CRUD API
// @version					1.0
// @description				API Server for Movies CRUD Application
// @host						localhost:8080
// @BasePath					/api/v1
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			database.NewDatabase,
			repository.NewMovieRepository,
			repository.NewUserRepository,
			auth.NewJWTService,
			service.NewMovieService,
			service.NewAuthService,
			middleware.NewAuthMiddleware,
			controller.NewMovieController,
			controller.NewAuthController,
			newRouter,
		),
		fx.Invoke(registerHooks),
	)

	app.Run()
}

func newRouter(
	config *config.Config,
	movieController *controller.MovieController,
	authController *controller.AuthController,
	authMiddleware *middleware.AuthMiddleware,
) *gin.Engine {
	if config.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	api := router.Group("/api/v1")
	{
		authApi := api.Group("/authApi")
		{
			authApi.POST("/login", authController.Login)
		}

		movies := api.Group("/movies")
		movies.Use(authMiddleware.JWTAuth())
		{
			movies.GET("", movieController.GetAllMovies)
			movies.GET("/:id", movieController.GetMovieByID)
			movies.GET("/search", movieController.SearchMovies)

			adminRoutes := movies.Group("")
			adminRoutes.Use(authMiddleware.RoleAuth("admin"))
			{
				adminRoutes.POST("", movieController.CreateMovie)
				adminRoutes.PUT("/:id", movieController.UpdateMovie)
				adminRoutes.DELETE("/:id", movieController.DeleteMovie)
			}
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func registerHooks(
	lifecycle fx.Lifecycle,
	config *config.Config,
	router *gin.Engine,
	authService *service.AuthService,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				err := authService.EnsureAdminExists(config.AdminUser, config.AdminPass)
				if err != nil {
					return fmt.Errorf("failed to create admin user: %w", err)
				}

				go func() {
					addr := config.GetAppAddress()
					log.Printf("Starting server on %s", addr)
					if err := http.ListenAndServe(addr, router); err != nil {
						log.Fatalf("Server failed to start: %v", err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Println("Shutting down server")
				return nil
			},
		},
	)
}
