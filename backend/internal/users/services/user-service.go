package services

import (
	"context"
	pb "github.com/ghandic/grpc-web-go-react-example/backend/api/proto/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/db/users"
	"github.com/ghandic/grpc-web-go-react-example/backend/internal/users/domain"
)

type UserServiceHandler struct {
	userRepository UserRepository
}

type UserRepository interface {
	GetUser(ctx context.Context, UserId int32) (*users.User, error)
	CreateUser(ctx context.Context, Name string) (*users.User, error)
	ListUsers(ctx context.Context, Req *pb.ListUsersRequest) (*domain.ListUsersResponse, error)
	DeleteUser(ctx context.Context, UserId int32) (bool, error)
}

func NewUserService(userRepository UserRepository) *UserServiceHandler {
	return &UserServiceHandler{userRepository: userRepository}
}

func (s *UserServiceHandler) GetUser(ctx context.Context, UserId int32) (*users.User, error) {
	return s.userRepository.GetUser(ctx, UserId)
}

func (s *UserServiceHandler) CreateUser(ctx context.Context, Name string) (*users.User, error) {
	return s.userRepository.CreateUser(ctx, Name)
}

func (s *UserServiceHandler) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*domain.ListUsersResponse, error) {
	return s.userRepository.ListUsers(ctx, req)
}

func (s *UserServiceHandler) DeleteUser(ctx context.Context, UserId int32) (bool, error) {
	return s.userRepository.DeleteUser(ctx, UserId)
}
