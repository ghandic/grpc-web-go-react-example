package usecases

import (
	"github.com/ghandic/grpc-web-go-react-example/backend/db/users"
)

func (a *UserUsecases) GetUser(UserId int32) (*users.User, error) {
	return a.userService.GetUser(UserId)
}
