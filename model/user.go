package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
}

func (u *User) CreateDigestPassword(password string) string {
	bytes := []byte(password)
	digestPassword, err := bcrypt.GenerateFromPassword(bytes, 12)
	if err != nil {
		panic("加密错误！")
	}
	return string(digestPassword)
}

func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password)) == nil
}
