package main

import (
	"github.com/gin-gonic/gin"
	authapi "ndy/realworld-gin/internal/auth/api"
	authapp "ndy/realworld-gin/internal/auth/app"
	authinfra "ndy/realworld-gin/internal/auth/infra"
	"ndy/realworld-gin/internal/config"
	profileapi "ndy/realworld-gin/internal/profile/api"
	profileapp "ndy/realworld-gin/internal/profile/app"
	profileinfra "ndy/realworld-gin/internal/profile/infra"
	userapi "ndy/realworld-gin/internal/user/api"
	userapp "ndy/realworld-gin/internal/user/app"
	userinfra "ndy/realworld-gin/internal/user/infra"
	"ndy/realworld-gin/internal/util"
)

func main() {
	util.Log.Info("Application starting...")

	defer util.Sync()
	defer util.DB.Close()

	// Config
	mysqlDsn := config.MysqlConfig.FormatDSN()

	// Auth
	authRepo := authinfra.NewMysqlRepo(mysqlDsn)
	authLogic := authapp.NewLogicImpl(authRepo)

	// User
	userRepo := userinfra.NewMysqlRepo(mysqlDsn)
	userLogic := userapp.NewLogic(userRepo)

	// Profile
	profileRepo := profileinfra.NewMysqlRepo(mysqlDsn)
	profileLogic := profileapp.NewLogicImpl(profileRepo)

	// Create a new Gin app
	r := gin.Default()

	// Register Api Routes
	api := r.Group("/api")
	{
		authapi.Routes(api, &authLogic)
		userapi.Routes(api, &userLogic)
		profileapi.Routes(api, &profileLogic)
	}

	// Run the app
	_ = r.Run(":8080")
}
