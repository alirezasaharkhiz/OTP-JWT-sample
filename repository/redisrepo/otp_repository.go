package redisrepo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"OTP-JWT-sample/repository"
	"github.com/go-redis/redis/v8"
)

type redisOtpRepository struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisOtpRepository constructor
func NewRedisOtpRepository(client *redis.Client) repository.OtpRepository {
	return &redisOtpRepository{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *redisOtpRepository) SetOTP(mobile string, otp string, ttlSeconds int) error {
	return r.client.Set(r.ctx, r.key(mobile), otp, time.Duration(ttlSeconds)*time.Second).Err()
}

func (r *redisOtpRepository) GetOTP(mobile string) (string, error) {
	val, err := r.client.Get(r.ctx, r.key(mobile)).Result()
	if errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("OTP not found")
	}
	return val, err
}

func (r *redisOtpRepository) DeleteOTP(mobile string) error {
	return r.client.Del(r.ctx, r.key(mobile)).Err()
}

func (r *redisOtpRepository) key(mobile string) string {
	return fmt.Sprintf("otp:%s", mobile)
}
