package usecases

import (
	"github.com/ghandic/grpc-web-go-react-example/backend/db/users"
)

func (a *UserUsecases) CreateUser(Name string) (*users.User, error) {
	return a.userService.CreateUser(Name)
}
