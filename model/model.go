package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DataBase(connPath string) {
	db, err := gorm.Open(mysql.Open(connPath), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("数据库连接成功！")
	DB = db
}

func Migrator() error {
	if err1 := DB.AutoMigrate(&User{}); err1 != nil {
		return err1
	}

	if err2 := DB.AutoMigrate(&Task{}); err2 != nil {
		return err2
	}
	return nil
}
