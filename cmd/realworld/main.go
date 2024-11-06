package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	authapi "ndy/realworld-gin/internal/auth/api"
	authapp "ndy/realworld-gin/internal/auth/app"
	authinfra "ndy/realworld-gin/internal/auth/infra"
	profileapi "ndy/realworld-gin/internal/profile/api"
	profileapp "ndy/realworld-gin/internal/profile/app"
	profileinfra "ndy/realworld-gin/internal/profile/infra"
	userapi "ndy/realworld-gin/internal/user/api"
	userapp "ndy/realworld-gin/internal/user/app"
	userinfra "ndy/realworld-gin/internal/user/infra"
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
	authRepo := authinfra.NewMysqlRepo(db)
	authLogic := authapp.NewLogicImpl(authRepo)

	// User
	userRepo := userinfra.NewMysqlRepo(db)
	userLogic := userapp.NewLogic(userRepo)

	// Profile
	profileRepo := profileinfra.NewMysqlRepo(db)
	profileLogic := profileapp.NewLogicImpl(profileRepo)

	// Register Api Routes
	api := r.Group("/api")
	{
		authapi.Routes(api, &authLogic)
		userapi.Routes(api, &userLogic)
		profileapi.Routes(api, &profileLogic)
	}

	// Run the app
	err = r.Run(":8080")
	if err != nil {
		util.Log.Fatal(err.Error())
		return
	}
}
