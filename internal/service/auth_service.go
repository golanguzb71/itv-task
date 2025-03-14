package service

import (
	"errors"
	"itv/internal/dto"
	"itv/internal/repository"
	"itv/pkg/auth"
)

type AuthService struct {
	userRepo   *repository.UserRepository
	jwtService *auth.JWTService
}

func NewAuthService(userRepo *repository.UserRepository, jwtService *auth.JWTService) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *AuthService) Login(loginDTO dto.LoginDTO) (*dto.TokenResponseDTO, error) {
	user, err := s.userRepo.FindByUsername(loginDTO.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !auth.CheckPasswordHash(loginDTO.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponseDTO{Token: token}, nil
}

func (s *AuthService) EnsureAdminExists(username, password string) error {
	return s.userRepo.EnsureAdminExists(username, password)
}
