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