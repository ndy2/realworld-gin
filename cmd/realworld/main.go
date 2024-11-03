package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	auth2 "ndy/realworld-gin/internal/auth/api"
	"ndy/realworld-gin/internal/auth/app"
	auth3 "ndy/realworld-gin/internal/auth/infra"
	middleware2 "ndy/realworld-gin/internal/middleware"
	profile3 "ndy/realworld-gin/internal/profile"
	user3 "ndy/realworld-gin/internal/user"
	"ndy/realworld-gin/internal/util"
)

func main() {
	util.InitLogger()
	defer util.Sync()

	util.Log.Info("Application starting...")

	// Create a new Gin app
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
		util.Log.Fatal(err.Error())
	}

	// Auth
	authRepo := auth3.NewMysqlRepo(db)
	authLogic := app.NewLogicImpl(authRepo)

	// User
	userRepo := user3.NewMysqlRepo(db)
	userLogic := user3.NewLogic(userRepo)

	// Profile
	profileRepo := profile3.NewMysqlRepo(db)
	profileLogic := profile3.NewLogic(profileRepo)

	// Middlewares
	um := middleware2.JsonRoot("user", "user")
	umRespOnly := middleware2.JsonRoot("", "user")
	am := middleware2.Auth()
	pmRespOnly := middleware2.JsonRoot("", "profile")

	// Routes
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/api")
	{
		// users
		api.POST("/users/login", um, auth2.AuthenticationHandler(&authLogic))
		api.POST("/users", um, user3.RegisterHandler(&userLogic))
		api.GET("/user", am, umRespOnly, user3.GetCurrentUserHandler(&userLogic))
		api.PUT("/user", am, um, user3.UpdateUserHandler(&userLogic))

		// profiles
		api.GET("/profiles/:username", am, pmRespOnly, profile3.GetProfileHandler(&profileLogic))
	}

	// Run the app
	err = r.Run(":8080")
	if err != nil {
		util.Log.Fatal(err.Error())
		return
	}
}
