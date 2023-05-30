package usecases

import (
	"context"
	pb "github.com/ghandic/grpc-web-go-react-example/backend/api/proto/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/db/users"
	"github.com/ghandic/grpc-web-go-react-example/backend/internal/users/domain"
)

type UserUsecases struct {
	userService UserService
}

type UserService interface {
	GetUser(ctx context.Context, UserId int32) (*users.User, error)
	CreateUser(ctx context.Context, Name string) (*users.User, error)
	ListUsers(ctx context.Context, Req *pb.ListUsersRequest) (*domain.ListUsersResponse, error)
	DeleteUser(ctx context.Context, UserId int32) (bool, error)
}

func NewUserUsecases(userService UserService) *UserUsecases {
	return &UserUsecases{userService: userService}
}
