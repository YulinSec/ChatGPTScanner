package models

import (
	"database/sql"

	"github.com/beego/beego/v2/client/orm"
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}

type User struct {
	Id       int            `json:"id"`
	Username sql.NullString `json:"username"`
	Password sql.NullString `json:"password"`
}
