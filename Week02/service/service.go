package service

import (
	"Week02/dao"
	"log"

	goerrors "github.com/pkg/errors"
)

type UserService struct {
	db dao.Dao
}

func New() *UserService {
	return &UserService{}
}

func (userService *UserService) GetUserInfo(userid int) (int, error) {
	userinfo, err := userService.db.GetUserInfo(userid)
	if err != nil {
		log.Printf("原始错误发生信息：%T %v\n", goerrors.Cause(err), goerrors.Cause(err))
		return 0, err
	}
	return userinfo, err

}
