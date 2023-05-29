package usecases

import (
	"context"
	"github.com/ghandic/grpc-web-go-react-example/backend/db/users"
)

func (a *UserUsecases) CreateUser(ctx context.Context, Name string) (*users.User, error) {
	return a.userService.CreateUser(ctx, Name)
}
