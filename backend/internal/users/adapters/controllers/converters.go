package controllers

import (
	v1 "github.com/ghandic/grpc-web-go-react-example/backend/api/proto/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/db/users"
)

func PGUsersToUsers(pgUsers []users.User) []*v1.User {
	var acc []*v1.User

	for _, u := range pgUsers {
		tmpUser := v1.User{Id: u.ID, Name: u.Name}
		acc = append(acc, &tmpUser)
	}

	return acc
}
