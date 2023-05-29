package usecases

import (
	"context"
	"github.com/ghandic/grpc-web-go-react-example/backend/db/users"
)

func (a *UserUsecases) GetUser(ctx context.Context, UserId int32) (*users.User, error) {
	return a.userService.GetUser(ctx, UserId)
}
