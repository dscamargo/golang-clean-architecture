package main

import (
	"fmt"
	"github.com/dscamargo/crud-clean-architecture/src/adapters"
	"github.com/dscamargo/crud-clean-architecture/src/domain"
	"github.com/dscamargo/crud-clean-architecture/src/presentation/controllers"
	"github.com/dscamargo/crud-clean-architecture/src/user/repository"
	"github.com/dscamargo/crud-clean-architecture/src/user/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func connectDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(sqlite.Open("account.db"), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalln("[Database] - Error in connection", err)
	}
	return db
}

func main() {
	db := connectDB()
	err := db.AutoMigrate(&domain.User{}, &domain.Address{})
	if err != nil {
		log.Fatalln("[AutoMigrate] - Error in AutoMigrate", err)
	}
	app := gin.Default()

	//Dependencias
	userRepo := repository.NewSQLiteUserRepository(db)
	hasher := adapters.NewHasher()

	//Usecases
	createUserUsecase := usecase.NewCreateUserUsecase(userRepo, hasher)
	getUserUsecase := usecase.NewGetUserUsecase(userRepo)

	//Controllers
	userController := controllers.NewUserController(createUserUsecase, getUserUsecase)

	app.POST("/users", userController.CreateUser)
	app.GET("/users/:id", userController.GetUser)

	err = app.Run(":8090")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("[Server] Connected")
}
