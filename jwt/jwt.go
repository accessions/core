package jwt

import (
	"encoding/json"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var ServiceJwt *JWT


type JWT struct{
	Secret []byte
	option *Claims
}
type Claims struct {
	Unique_name  string `json:"unique_name"`
	Guid         string `json:"guid"`
	Avatar       string `json:"avatar"`
	DisplayName  string `json:"displayName"`
	LoginName    string `json:"loginName"`
	EmailAddress string `json:"emailAddress"`
	UserType     string `json:"userType"`
	Time         string `json:"time"`
	jwt.StandardClaims
}

func init()  {
	ServiceJwt = &JWT{}
}

type newOption func(claims *Claims)

// New
func (js *JWT) New (secret string, opts ...newOption) *JWT {

	option := &Claims{
		Unique_name:    "",
		Guid:           "",
		Avatar:         "",
		DisplayName:    "",
		LoginName:      "smoke",
		EmailAddress:   "smoke.mvp@gmail.com",
		UserType:       "git",
		Time:           time.Now().String(),
		StandardClaims: jwt.StandardClaims{},
	}
	js.Secret = []byte(secret)
	for _, o := range opts {
		o(option)
	}
	js.option = option
	return js
}



// Generate 生成token
func (js *JWT) Generate() (string, error) {
	return _generate(js.option, js.Secret)
}


// ParseClaims 解析成内置对象
func (js *JWT) ParseClaims(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(tokens *jwt.Token) (interface{}, error) {
		return js.Secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// ParseJson 解析成json字符串
func (js *JWT) ParseJson (token string) (string, error) {
	tokenRst, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return js.Secret, nil
	})
	if tokenRst != nil {
		if tokenRst.Valid {
			claimsJson, _ := json.Marshal(tokenRst.Claims)
			return string(claimsJson), nil
		}
	}
	return "", err
}

// _generate
func _generate(claims interface{}, secret []byte) (string, error) {
	if tmp, ok := claims.(jwt.Claims); ok {
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, tmp)
		token, err := tokenClaims.SignedString(secret)
		return token, err
	}
	return "", errors.New("claims is not jwt.Claims")
}