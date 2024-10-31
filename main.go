package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"log"
	"ndy/realworld-gin/auth"
	"ndy/realworld-gin/middleware"
	"ndy/realworld-gin/user"
)

func main() {
	// Create a new Gin application
	r := gin.Default()

	// Capture connection properties
	var db *sql.DB
	cfg := mysql.Config{
		User:   "root",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "realworld",
	}
	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Auth
	authRepo := auth.NewMysqlRepo(db)
	authLogic := auth.NewLogic(authRepo)

	// User
	userRepo := user.NewMysqlRepo(db)
	userLogic := user.NewLogic(userRepo)

	// Middlewares
	um := middleware.JsonRoot("user", "user")
	umRespOnly := middleware.JsonRoot("", "user")
	am := middleware.Auth()

	// Routes
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// users
		api.POST("/users/login", um, auth.AuthenticationHandler(&authLogic))
		api.POST("/users", um, user.RegisterHandler(&userLogic))
		api.GET("/user", am, umRespOnly, user.GetCurrentUserHandler(&userLogic))
		api.PUT("/user", am, um, user.UpdateUserHandler(&userLogic))
	}

	// Run the application
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
