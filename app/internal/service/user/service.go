package user

import (
	"cex-core-api/app/internal/service"
	user_v1 "cex-core-api/gen/user/v1"
)

type UserService struct {
	userRepo service.UserRepository
	user_v1.UnimplementedUserServiceServer
}

func New(userRepo service.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
