package usecases

import "context"

func (a *UserUsecases) DeleteUser(ctx context.Context, UserId int32) (bool, error) {
	return a.userService.DeleteUser(ctx, UserId)
}
