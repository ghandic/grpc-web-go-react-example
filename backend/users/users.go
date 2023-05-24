package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	v1 "github.com/ghandic/grpc-web-go-react-example/backend/gen/proto/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/gen/proto/users/v1/usersv1connect"

	"github.com/bufbuild/connect-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	usersv1connect.UnimplementedUserServiceHandler
	Pool *pgxpool.Pool
}

func (u *UserService) GetUser(
	ctx context.Context,
	req *connect.Request[v1.GetUserRequest],
) (*connect.Response[v1.GetUserResponse], error) {

	q := New(u.Pool)

	pg_user, err := q.GetUser(ctx, req.Msg.UserId)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Error(codes.NotFound, "User not found")
		} else {
			return nil, status.Errorf(codes.Internal, "GetUser failed: %v\n", err)
		}
	}

	return connect.NewResponse(&v1.GetUserResponse{
		User: &v1.User{Id: pg_user.ID, Name: pg_user.Name},
	}), nil
}

func (u *UserService) CreateUser(
	ctx context.Context,
	req *connect.Request[v1.CreateUserRequest],
) (*connect.Response[v1.CreateUserResponse], error) {

	if req.Msg.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing name")
	}

	q := New(u.Pool)

	pg_user, err := q.AddUser(ctx, req.Msg.Name)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateUser failed: %v\n", err)
	}

	user := &v1.User{
		Id:   pg_user.ID,
		Name: pg_user.Name,
	}

	return connect.NewResponse(&v1.CreateUserResponse{
		User: user,
	}), nil
}

func PGUsersToUsers(pg_users []User) []*v1.User {
	var acc []*v1.User

	for _, u := range pg_users {
		tmpUser := v1.User{Id: u.ID, Name: u.Name}
		acc = append(acc, &tmpUser)
	}

	return acc
}

func getListParams(req *v1.ListUsersRequest) (*GetUsersParams, error) {
	listParams := &GetUsersParams{}
	if req.Sorting != nil {
		switch req.Sorting.Field {
		case "name":
			switch req.Sorting.Direction {
			case v1.SortDirection_ASC:
				listParams.NameAsc = true
			case v1.SortDirection_DESC:
				listParams.NameDesc = true
			default:
				return listParams, errors.New("Invalid sort by direction")
			}
		case "id":
			switch req.Sorting.Direction {
			case v1.SortDirection_ASC:
				listParams.IDAsc = true
			case v1.SortDirection_DESC:
				listParams.IDDesc = true
			default:
				return listParams, errors.New("Invalid sort by direction")
			}
		case "created_at":
			switch req.Sorting.Direction {
			case v1.SortDirection_ASC:
				listParams.CreatedAtAsc = true
			case v1.SortDirection_DESC:
				listParams.CreatedAtDesc = true
			default:
				return listParams, errors.New("Invalid sort by direction")
			}
		default:
			return listParams, errors.New("Invalid field")
		}
	}

	if req.Query != nil {
		if req.Query.Text != "" {
			listParams.Search = req.Query.Text
		}
	}

	if req.PageSize > 0 {
		listParams.LimitAmount = req.PageSize
	}

	if req.Offset >= 0 {
		listParams.OffsetAmount = req.Offset
	}

	return listParams, nil

}

func (u *UserService) ListUsers(
	ctx context.Context,
	req *connect.Request[v1.ListUsersRequest],
) (*connect.Response[v1.ListUsersResponse], error) {
	q := New(u.Pool)

	listParams, err := getListParams(req.Msg)
	fmt.Println(listParams)

	pg_users, err := q.GetUsers(ctx, *listParams)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetUsers failed: %v\n", err)
	}

	search := ""
	if req.Msg.Query != nil {
		search = req.Msg.Query.Text
	}
	total_count, err := q.GetUsersCount(ctx, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetUsersCount failed: %v\n", err)
	}

	return connect.NewResponse(&v1.ListUsersResponse{Users: PGUsersToUsers(pg_users), Total: total_count}), nil
}

func (u *UserService) DeleteUser(
	ctx context.Context,
	req *connect.Request[v1.DeleteUserRequest],
) (*connect.Response[v1.DeleteUserResponse], error) {

	q := New(u.Pool)

	err := q.DeleteUser(ctx, req.Msg.UserId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "DeleteUser failed: %v\n", err)
	}

	return connect.NewResponse(&v1.DeleteUserResponse{}), nil
}
