package database_handler

import (
	"fmt"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDB(conString string) *gorm.DB {

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名使用单数形式 例如 User 表名会是 User，而不是默认的 users）
			NoLowerCase:   true, // 表示表名和字段名不会自动转换为小写
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		panic(fmt.Sprintf("不能连接到数据库 : %s", err.Error()))
	}

	return db
}
