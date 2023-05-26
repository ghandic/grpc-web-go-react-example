package usecases

func (a *UserUsecases) DeleteUser(UserId int32) (bool, error) {
	return a.userService.DeleteUser(UserId)
}
