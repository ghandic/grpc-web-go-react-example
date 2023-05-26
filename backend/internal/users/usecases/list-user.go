package usecases

import (
	v1 "github.com/ghandic/grpc-web-go-react-example/backend/api/proto/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/internal/users/domain"
)

func (a *UserUsecases) ListUsers(req *v1.ListUsersRequest) (*domain.ListUsersResponse, error) {
	return a.userService.ListUsers(req)
}
