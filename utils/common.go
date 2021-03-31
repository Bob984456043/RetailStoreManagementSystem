package utils

import (
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"strconv"
)

var secret = "984456043"

// GenerateToken 生成Token值
func GenerateToken(mapClaims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte(secret))
}

// token: "eyJhbGciO...解析token"
func ParseToken(token string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["cmd"].(string), nil
}
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func StringToInt64(a string) int64 {
	int64Val, int64ValErr := strconv.ParseInt(a, 10, 64)
	if int64ValErr != nil {
		return 0
	}
	return int64Val
}
