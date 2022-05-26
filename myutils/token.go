package myutils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	//加密秘钥，采用对称秘钥
	TokenKey []byte = []byte("ksks")
)

//加密内容
type Claim struct {
	Id           uint
	UserName     string
	UserPassword string
	jwt.StandardClaims
}

type TokenData struct {
	UserName     string
	UserPassword string
	Token        string
}

func GenerateToken(id uint, userName string, userPassword string) (string, error) {
	claims := Claim{
		Id:           id,
		UserName:     userName,
		UserPassword: userPassword,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "wangyiran",
		},
	}
	//加密.
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(TokenKey)
	return token, err
}

func BuildTokenData(username string, userpassword string, token string) *TokenData {
	return &TokenData{
		UserName:     username,
		UserPassword: userpassword,
		Token:        token,
	}
}

func ParseToken(token string) (*Claim, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return TokenKey, nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claim)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
		if time.Now().Unix() > claims.ExpiresAt {
			return nil, errors.New("过期")
		}
	}
	return nil, err
}
