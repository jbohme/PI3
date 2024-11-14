package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jbohme/crud/configs/database/mongodb"
	"github.com/jbohme/crud/internal/app/http/controller"
	"github.com/jbohme/crud/internal/app/http/routes"
	"github.com/jbohme/crud/internal/app/model/repository"
	"github.com/jbohme/crud/internal/app/model/service"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error connecting to MongoDB database: %s",
			err.Error(),
		)
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // ou o domínio do seu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	r := repository.NewUserRepository(database)
	s := service.NewUserDomainService(r)
	return controller.NewUserControllerInterface(s)
}
