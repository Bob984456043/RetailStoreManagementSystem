package utils

import (
	"fmt"
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
func ToString(value interface{}) string {
	switch value := value.(type) {
	case string:
		return value
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(int64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(uint64(value), 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	default:
		return fmt.Sprintf("%+v", value)
	}
}
func ToInt64(value interface{}) int64 {
	switch value := value.(type) {
	case string:
		n, _ := strconv.Atoi(value)
		return int64(n)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	default:
		return value.(int64)
	}
}
func GetInt64Fields(input interface{},name string) []int64 {
	object := reflect.ValueOf(input)
	var items []interface{}
	for i := 0; i < object.Len(); i++ {
		items = append(items, object.Index(i).Interface())
	}
	// Populate the rest of the items into <ids>
	var ids []int64
	for _, v := range items {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		id := val.FieldByName(name).Interface()
		ids = append(ids, id.(int64))
	}
	return ids
}
