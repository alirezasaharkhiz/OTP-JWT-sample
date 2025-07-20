package repository

import "OTP-JWT-sample/model"

type UserRepository interface {
	FindByMobile(mobile string) (*model.User, error)
	Create(user *model.User) error
}

type OtpRepository interface {
	SetOTP(mobile string, otp string, ttlSeconds int) error
	GetOTP(mobile string) (string, error)
	DeleteOTP(mobile string) error
}
