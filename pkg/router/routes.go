package router

import (
	"database/sql"

	"github.com/Dev-Awaab/go-base-api/config"
	"github.com/Dev-Awaab/go-base-api/internal/user"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func SetupRoutes(dbConn *sql.DB, cfg config.Config) {
	r = gin.Default()
	// User feature
	userRepo :=  user.NewUserRepository(dbConn)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// Register user routes
	userGroup := r.Group("/users")
	{
		userGroup.POST("", userHandler.CreateUser)
	}

	// Add more feature groups (e.g., `/products`, `/orders`) here
}

func Start(addr string) error {

	return r.Run(addr)
}