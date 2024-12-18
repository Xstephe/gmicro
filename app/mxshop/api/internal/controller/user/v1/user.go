package user

import (
	ut "github.com/go-playground/universal-translator"
	"mxshop/app/mxshop/api/internal/service"
)

type userServer struct {
	trans ut.Translator
	//srv   user.UserSrv
	srv service.ServiceFactory
}

func NewUserController(trans ut.Translator, srv service.ServiceFactory) *userServer {
	return &userServer{trans, srv}
}