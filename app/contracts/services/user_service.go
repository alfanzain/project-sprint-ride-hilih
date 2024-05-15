package services

type IUserService interface {
	GetUsers(any) (any, error)
}
