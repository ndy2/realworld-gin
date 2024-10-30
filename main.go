package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"log"
	"ndy/realworld-gin/auth"
	"ndy/realworld-gin/route"
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

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// users
		users := api.Group("/users")
		{
			users.POST("/login", common.HandleJsonRootMiddleware("user", "user"), auth.AuthenticationHandler(&authLogic))
			users.POST("/", common.HandleJsonRootMiddleware("user", "user"), user.RegisterHandler(&userLogic))
		}
	}

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
