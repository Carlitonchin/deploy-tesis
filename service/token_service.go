package service

import (
	"context"
	"crypto/rsa"
	"strconv"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
)

type tokenService struct {
	TokenRepository      model.TokenRepository
	PrivateKey           *rsa.PrivateKey
	PublicKey            *rsa.PublicKey
	RefreshSecret        string
	IDExipirationSec     int64
	RefreshExpirationSec int64
}

type TSConfig struct {
	TokenRepository      model.TokenRepository
	PrivateKey           *rsa.PrivateKey
	PublicKey            *rsa.PublicKey
	RefreshSecret        string
	IDExpirationSec      int64
	RefreshExpirationSec int64
}

func NewTokenService(c *TSConfig) model.TokenService {
	return &tokenService{
		PrivateKey:           c.PrivateKey,
		PublicKey:            c.PublicKey,
		RefreshSecret:        c.RefreshSecret,
		TokenRepository:      c.TokenRepository,
		IDExipirationSec:     c.IDExpirationSec,
		RefreshExpirationSec: c.RefreshExpirationSec,
	}
}

func (s *tokenService) GetNewPairFromUser(
	ctx context.Context,
	user *model.User,
	prevTokenId string) (*model.TokenPair, error) {

	//function body starts here

	userId_str := strconv.FormatUint(uint64(user.ID), 10)
	if prevTokenId != "" {
		err := s.TokenRepository.DeleteRefreshToken(ctx, userId_str, prevTokenId)

		if err != nil {
			return nil, err
		}
	}

	id_token, err := generateIDToken(user, s.PrivateKey, s.IDExipirationSec)
	if err != nil {
		return nil, err
	}

	refresh_token, err := generateRefreshToken(user.ID, s.RefreshSecret, s.RefreshExpirationSec)

	if err != nil {
		return nil, err
	}

	err = s.TokenRepository.SetNewRefreshToken(ctx, userId_str, refresh_token.ID, refresh_token.ExipresIn)

	if err != nil {
		return nil, err
	}

	return &model.TokenPair{
		IDToken: model.IDToken{SS: id_token},
		RefreshToken: model.RefreshToken{
			ID:  refresh_token.ID,
			UID: user.ID,
			SS:  refresh_token.SS,
		},
	}, nil
}

func (s *tokenService) ValidateIdToken(tokenString string) (*model.User, error) {
	claims, err := validateIdToken(tokenString, s.PublicKey)

	if err != nil {
		type_error := apperrors.Authorization
		message := "Token invalido"

		e := apperrors.NewError(type_error, message)
		return nil, e
	}

	return claims.User, nil
}

func (s *tokenService) ValidateRefreshToken(refresh_token string) (*model.RefreshToken, error) {
	claims, err := validateRefreshToken(refresh_token, s.RefreshSecret)

	if err != nil {
		type_error := apperrors.Authorization
		message := "Token invalido"

		e := apperrors.NewError(type_error, message)
		return nil, e
	}

	return &model.RefreshToken{
		ID:  claims.Id,
		UID: claims.ID,
		SS:  refresh_token,
	}, nil
}

func (s *tokenService) SignOut(ctx context.Context, user_id uint) error {
	return s.TokenRepository.DeleteUserRefreshTokens(ctx, strconv.FormatUint(uint64(user_id), 10))
}
