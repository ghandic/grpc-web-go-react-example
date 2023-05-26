package repositories

import (
	"context"
	"github.com/ghandic/grpc-web-go-react-example/backend/api/proto/users/v1/usersv1connect"

	v1 "github.com/ghandic/grpc-web-go-react-example/backend/api/proto/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/db/users"
	"github.com/ghandic/grpc-web-go-react-example/backend/internal/users/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRepository struct {
	usersv1connect.UnimplementedUserServiceHandler
	Pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{Pool: pool}
}

func (u *UserRepository) CreateUser(ctx context.Context, Name string) (*users.User, error) {

	if Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing name")
	}

	q := users.New(u.Pool)

	pgUser, err := q.AddUser(ctx, Name)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateUser failed: %v\n", err)
	}

	return &pgUser, nil
}

func (u *UserRepository) GetUser(
	ctx context.Context,
	userId int32,
) (*users.User, error) {

	q := users.New(u.Pool)

	pgUser, err := q.GetUser(ctx, userId)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Error(codes.NotFound, "User not found")
		} else {
			return nil, status.Errorf(codes.Internal, "GetUser failed: %v\n", err)
		}
	}

	return &pgUser, nil
}

func (u *UserRepository) ListUsers(
	ctx context.Context,
	req *v1.ListUsersRequest,
) (*domain.ListUsersResponse, error) {
	q := users.New(u.Pool)

	listParams, err := getListParams(req)

	pgUsers, err := q.GetUsers(ctx, *listParams)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetUsers failed: %v\n", err)
	}

	search := ""
	if req.Query != nil {
		search = req.Query.Text
	}
	totalCount, err := q.GetUsersCount(ctx, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetUsersCount failed: %v\n", err)
	}

	return &domain.ListUsersResponse{Users: &pgUsers, TotalCount: totalCount}, nil
}

func (u *UserRepository) DeleteUser(
	ctx context.Context,
	userId int32,
) (bool, error) {
	q := users.New(u.Pool)

	err := q.DeleteUser(ctx, userId)

	if err != nil {
		return false, status.Errorf(codes.Internal, "DeleteUser failed: %v\n", err)
	}

	return true, nil
}
