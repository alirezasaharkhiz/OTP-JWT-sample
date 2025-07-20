package main

import (
	"OTP-JWT-sample/model"
	"OTP-JWT-sample/repository/redisrepo"
	"log"
	"os"

	_ "github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"OTP-JWT-sample/controller"
	"OTP-JWT-sample/repository/postgresrepo"
	"OTP-JWT-sample/router"
	"OTP-JWT-sample/service"
)

func main() {
	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	userRepo := postgresrepo.NewPostgresUserRepository(db)
	otpRepo := redisrepo.NewRedisOtpRepository(rdb)

	authService := service.NewAuthService(userRepo, otpRepo)
	authController := controller.NewAuthController(authService)

	r := router.SetupRouter(authController)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
