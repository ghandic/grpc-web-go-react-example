package controllers

import (
	"context"
	v1 "github.com/ghandic/grpc-web-go-react-example/backend/api/proto/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/db/users"
	"github.com/ghandic/grpc-web-go-react-example/backend/internal/users/domain"
	"log"
)

type UserHandler struct {
	userUsecases UserUsecases
}

type UserUsecases interface {
	GetUser(UserId int32) (*users.User, error)
	CreateUser(Name string) (*users.User, error)
	ListUsers(Req *v1.ListUsersRequest) (*domain.ListUsersResponse, error)
	DeleteUser(UserId int32) (bool, error)
}

func NewUserHandler(userUsecases UserUsecases) *UserHandler {
	return &UserHandler{userUsecases: userUsecases}
}

func (h *UserHandler) CreateUser(ctx context.Context, request *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	user, err := h.userUsecases.CreateUser(request.Name)
	if err != nil {
		log.Fatalf("Error occurred %v", err)
	}

	return &v1.CreateUserResponse{User: &v1.User{
		Id:   user.ID,
		Name: user.Name,
	}}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, request *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	user, err := h.userUsecases.GetUser(request.UserId)
	if err != nil {
		log.Fatalf("Error occurred %v", err)
	}

	return &v1.GetUserResponse{User: &v1.User{
		Id:   user.ID,
		Name: user.Name,
	}}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, request *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	successfulDelete, err := h.userUsecases.DeleteUser(request.UserId)
	if !successfulDelete || err != nil {
		log.Fatalf("Error occurred %v", err)
	}
	return &v1.DeleteUserResponse{}, nil
}

func (h *UserHandler) ListUsers(ctx context.Context, request *v1.ListUsersRequest) (*v1.ListUsersResponse, error) {
	users, err := h.userUsecases.ListUsers(request)

	if err != nil {
		log.Fatalf("Error occurred %v", err)
	}

	return &v1.ListUsersResponse{Users: PGUsersToUsers(*users.Users), Total: users.TotalCount}, nil
}
