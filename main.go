package main

import (
	"context"
	"fmt"
	"github.com/dscamargo/crud-clean-architecture/src/adapters"
	"github.com/dscamargo/crud-clean-architecture/src/presentation/controllers"
	"github.com/dscamargo/crud-clean-architecture/src/user/repository"
	"github.com/dscamargo/crud-clean-architecture/src/user/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func connectMongoDB() (*mongo.Database, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
	}
	db := conn.Database("golang-test")
	fmt.Print("[Database] Connected")
	return db, cancel
}

func main() {
	db, cancel := connectMongoDB()
	defer cancel()

	app := gin.Default()

	//Dependencias
	userRepo := repository.NewMongoUserRepository(db)
	hasher := adapters.NewHasher()

	//Usecases
	createUserUsecase := usecase.NewCreateUserUsecase(userRepo, hasher)
	getUserUsecase := usecase.NewGetUserUsecase(userRepo)

	//Controllers
	userController := controllers.NewUserController(createUserUsecase, getUserUsecase)

	app.POST("/users", userController.CreateUser)
	app.GET("/users/:id", userController.GetUser)

	err := app.Run(":8090")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("[Server] Connected")
}
