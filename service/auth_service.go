package service

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"OTP-JWT-sample/model"
	"OTP-JWT-sample/repository"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userRepo repository.UserRepository
	otpRepo  repository.OtpRepository
}

func NewAuthService(userRepo repository.UserRepository, otpRepo repository.OtpRepository) *AuthService {
	return &AuthService{userRepo: userRepo, otpRepo: otpRepo}
}

func (s *AuthService) RequestOTP(mobile string) error {
	otp := generateOTP()
	err := s.otpRepo.SetOTP(mobile, otp, 120)
	if err != nil {
		return err
	}
	log.Printf("[OTP] mobile=%s otp=%s", mobile, otp)
	return nil
}

func (s *AuthService) VerifyOTP(mobile, otp string) (string, error) {
	storedOtp, err := s.otpRepo.GetOTP(mobile)
	if err != nil {
		return "", err
	}
	if storedOtp != otp {
		return "", fmt.Errorf("invalid OTP")
	}
	_ = s.otpRepo.DeleteOTP(mobile)

	user, err := s.userRepo.FindByMobile(mobile)
	if err != nil {
		user = &model.User{Mobile: mobile, CreatedAt: time.Now()}
		_ = s.userRepo.Create(user)
	}

	token := generateJWT(user.Mobile)
	return token, nil
}

func generateOTP() string {
	mx := int64(1_000_000) // 10^6
	nBig, err := rand.Int(rand.Reader, big.NewInt(mx))
	if err != nil {
		return "000000"
	}
	otp := fmt.Sprintf("%06d", nBig.Int64())
	return otp
}

func generateJWT(mobile string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mobile": mobile,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("secret"))
	return tokenString
}
