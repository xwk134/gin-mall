package util

import "github.com/dgrijalva/jwt-go"

var jwtSecret = []byte("yijiansanlian")

type Clims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `josn:"authority"`
	jwt.StandardClaims
}

// 签发token
func GenerateToken(id uint, userName string, authority int) (string, error) {

}
