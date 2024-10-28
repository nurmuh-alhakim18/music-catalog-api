package user

import (
	"context"
	"errors"

	"github.com/nurmuh-alhakim18/music-catalog-api/internal/models/user"
	"github.com/nurmuh-alhakim18/music-catalog-api/internal/repositories"
	"github.com/nurmuh-alhakim18/music-catalog-api/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	secretKeyJWT string
	queries      *repositories.Queries
}

func NewUserService(secretKeyJWT string, queries *repositories.Queries) *UserService {
	return &UserService{
		secretKeyJWT: secretKeyJWT,
		queries:      queries,
	}
}

func (s *UserService) Register(ctx context.Context, req user.UserRegisterRequest) (user.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return user.User{}, errors.New("failed to hash password")
	}

	userCreated, err := s.queries.CreateUser(ctx, repositories.CreateUserParams{
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
	})
	if err != nil {
		return user.User{}, errors.New("failed to create user")
	}

	return user.User{
		ID:        userCreated.ID,
		Email:     userCreated.Email,
		Username:  userCreated.Username,
		CreatedAt: userCreated.CreatedAt,
		UpdatedAt: userCreated.UpdatedAt,
	}, nil
}

func (s *UserService) Login(ctx context.Context, req user.UserLoginRequest) (user.UserLoginResponse, error) {
	userDB, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return user.UserLoginResponse{}, errors.New("incorrect email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(req.Password))
	if err != nil {
		return user.UserLoginResponse{}, errors.New("incorrect email or password")
	}

	token, err := auth.GenerateJWT(userDB.ID, s.secretKeyJWT)
	if err != nil {
		return user.UserLoginResponse{}, errors.New("failed to generate access token")
	}

	return user.UserLoginResponse{
		User: user.User{
			ID:        userDB.ID,
			Email:     userDB.Email,
			Username:  userDB.Username,
			CreatedAt: userDB.CreatedAt,
			UpdatedAt: userDB.UpdatedAt,
		},
		Token: token,
	}, nil
}
