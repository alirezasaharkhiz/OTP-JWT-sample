package router

import (
	"OTP-JWT-sample/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authController *controller.AuthController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/auth")
	{
		api.POST("/request-otp", authController.RequestOTP)
		api.POST("/verify-otp", authController.VerifyOTP)
	}

	return r
}
