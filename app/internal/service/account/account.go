package account

import (
	"cex-core-api/app/internal/service"
	user_v1 "cex-core-api/gen/user/v1"
)

type AccountService struct {
	accountRepo service.AccountRepository
	user_v1.UnimplementedUserServiceServer
}

func New(accountRepo service.AccountRepository) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}
