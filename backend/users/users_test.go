package users

import (
	"testing"

	v1 "github.com/ghandic/grpc-web-go-react-example/backend/gen/proto/users/v1"
	"github.com/stretchr/testify/assert"
)

func TestPGUsersToUsers(t *testing.T) {
	pgUsers := []User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
		{ID: 3, Name: "Bob"},
	}

	expectedUsers := []*v1.User{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 3, Name: "Bob"},
	}

	result := PGUsersToUsers(pgUsers)

	assert.Equal(t, expectedUsers, result)
}
