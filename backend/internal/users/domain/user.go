package domain

import "github.com/ghandic/grpc-web-go-react-example/backend/db/users"

type ListUsersResponse struct {
	Users      *[]users.User
	TotalCount int64
}
