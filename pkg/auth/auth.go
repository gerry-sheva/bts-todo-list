package auth

import (
	"context"
	"errors"

	"github.com/gerry-sheva/bts-todo-list/pkg/database/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrJWTGenerationError = errors.New("failed to generate JWT")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("User not found")
)

func register(ctx context.Context, dbpool *pgxpool.Pool, i *AuthInput) (string, error) {
	pwd, err := hashPassword(i.Password)
	if err != nil {
		return "", err
	}

	p := repository.NewUserParams{
		Username: i.Username,
		Password: pwd,
	}

	user, err := repository.New(dbpool).NewUser(ctx, p)
	if err != nil {
		return "", err
	}

	jwt, err := createJWT(user)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func login(ctx context.Context, dbpool *pgxpool.Pool, i *AuthInput) (string, error) {
	user, err := repository.New(dbpool).GetUser(ctx, i.Username)
	if err != nil {
		return "", ErrUserNotFound
	}

	match, _, err := verifyPassword(i.Password, user.Password)
	if err != nil {
		return "", err
	}

	if !match {
		return "", ErrInvalidCredentials
	}

	jwt, err := createJWT(user.Username)
	if err != nil {
		return "", ErrJWTGenerationError
	}

	return jwt, nil
}
