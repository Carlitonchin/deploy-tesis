package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/go-redis/redis/v8"
)

type redisTokenRepository struct {
	RedisClient *redis.Client
}

func NewTokenRepository(redis_client *redis.Client) model.TokenRepository {
	return &redisTokenRepository{
		RedisClient: redis_client,
	}
}

func (s *redisTokenRepository) SetNewRefreshToken(ctx context.Context, userId string,
	tokenId string, expiresIn time.Duration) error {

	key := fmt.Sprintf("%s:%s", userId, tokenId)
	err := s.RedisClient.Set(ctx, key, 0, expiresIn).Err()

	return err
}

func (s *redisTokenRepository) DeleteRefreshToken(ctx context.Context,
	userId string, prevTokenId string) error {

	key := fmt.Sprintf("%s:%s", userId, prevTokenId)
	result := s.RedisClient.Del(ctx, key)

	if result.Err() != nil {
		return result.Err()
	}

	if result.Val() < 1 {
		type_error := apperrors.Authorization
		message := "Invalid Token"
		return apperrors.NewError(type_error, message)
	}

	return nil
}

func (s *redisTokenRepository) DeleteUserRefreshTokens(ctx context.Context, userId string) error {
	pattern := fmt.Sprintf("%s*", userId)

	iter := s.RedisClient.Scan(ctx, 0, pattern, 5).Iterator()

	fail := false

	for iter.Next(ctx) {
		if err := s.RedisClient.Del(ctx, iter.Val()).Err(); err != nil {
			fail = true
		}
	}

	if err := iter.Err(); err != nil {
		fail = true
	}

	if fail {
		type_error := apperrors.Internal
		message := "Error interno al cerrar sesion del usuario"

		return apperrors.NewError(type_error, message)
	}

	return nil
}
