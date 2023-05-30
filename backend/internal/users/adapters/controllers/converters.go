package controllers

import (
	pb "github.com/ghandic/grpc-web-go-react-example/backend/api/proto/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/db/users"
)

func PGUsersToUsers(pgUsers []users.User) []*pb.User {
	var acc []*pb.User

	for _, u := range pgUsers {
		tmpUser := pb.User{Id: u.ID, Name: u.Name}
		acc = append(acc, &tmpUser)
	}

	return acc
}
