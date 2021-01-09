package test

import (
	"fmt"
	"github.com/accessions/core/jwt"
	"testing"
)

func TestJwt_Builder(t *testing.T)  {

	generate, err := jwt.ServiceJwt.New("hello", func(claims *jwt.Claims) {
		claims.DisplayName = "file"
		claims.Avatar = "test"
		claims.EmailAddress = "haha@163.com"
	}).Generate()

	fmt.Println("generate: jwt : ",generate, err)

	claims, err := jwt.ServiceJwt.ParseClaims(generate)

	fmt.Println("parse - option :", claims, err)
}


