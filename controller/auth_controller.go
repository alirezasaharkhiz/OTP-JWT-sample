package controller

import (
	"OTP-JWT-sample/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.POST("/request-otp", c.RequestOTP)
	auth.POST("/verify-otp", c.VerifyOTP)
}

func (c *AuthController) RequestOTP(ctx *gin.Context) {
	var req struct {
		Mobile string `json:"mobile"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.authService.RequestOTP(req.Mobile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}

func (c *AuthController) VerifyOTP(ctx *gin.Context) {
	var req struct {
		Mobile string `json:"mobile"`
		OTP    string `json:"otp"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := c.authService.VerifyOTP(req.Mobile, req.OTP)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
