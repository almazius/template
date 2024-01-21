package usecase

import "template/internal/auth/repo"

type AuthService struct {
	repo *repo.AuthPGRepo
}
