package srv

import (
	"cartoon-user/model"
	"golang.org/x/net/context"
)

type Service interface {
	//Regist(ctx context.Context)(string, error)
	Login(ctx context.Context)(map[string]interface{},error)
	Logout(ctx context.Context)(bool, error)
	FetchOne(ctx context.Context)(model.User, error)
}

type UserService struct {}

func (srv UserService) Login(ctx context.Context) (result map[string]interface{}, err error) {
	return nil, nil
}

func (srv UserService) Logout(ctx context.Context) (ok bool, err error) {
	return false, nil
}

func (srv UserService) FetchOne(ctx context.Context) (user model.User, err error) {
	return model.User{}, nil
}

func NewService() Service {
	return UserService{}
}
