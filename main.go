package main

import (
	"context"
	"fmt"
	dbConn "github.com/StaZisS/web/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var (
	server *gin.Engine
	ctx    context.Context
	db     *dbConn.Queries
	//AuthController controllers.AuthController
	//UserController controllers.UserController
	//AuthRoutes     routes.AuthRoutes
	//UserRoutes     routes.UserRoutes
)

func init() {
	ctx = context.TODO()
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer db.Close(context.Background())

	server = gin.Default()
}

func main() {
	router := server.Group("/api")

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to Golang with PostgreSQL"})
	})
	//AuthRoutes.AuthRoutes(router)
	//UserRoutes.UserRoutes(router)
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})
	log.Fatal(server.Run(":" + os.Getenv("PORT")))
}
