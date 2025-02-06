package service

import (
	"github.com/google/wire"
	userV1 "github.com/sunmery/kratos-template/api/user/v1"
	"github.com/sunmery/kratos-template/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	userV1.UnimplementedUserServiceServer

	uc *biz.UserUsecase
}

// NewUserService new a User service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}
