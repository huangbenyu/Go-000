package dao

import (
	"database/sql"

	"github.com/pkg/errors"
)

type Dao struct {
}

func New() *Dao {
	return &Dao{}

}

func (dao *Dao) Init() {

}

func (dao *Dao) GetUserInfo(userid int) (int, error) {

	return 0, errors.Wrap(sql.ErrNoRows, "dao error")
}
