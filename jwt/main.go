package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	NewTwo()
}

func NewOne() {
	mySigningKey := []byte("dagedaidaiwo")

	//StandardClaims struct
	c := MyClaims{
		Username: "Icarus",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "Icarus",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := t.SignedString(mySigningKey)
	if nil != err {
		fmt.Printf("%s", err)
		return
	}
	fmt.Println("打印：", s)

	token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if nil != err {
		fmt.Printf("%s", err)
		return
	}
	fmt.Println("打印token数据：", token.Claims.(*MyClaims).Username)
}

func NewTwo() {
	mySigningKey := []byte("dagedaidaiwo")

	//MapClaims map

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Unix() + 5,
		"iss":      "icarus",
		"nbf":      time.Now().Unix() - 5,
		"username": "my",
	})
	s, err := t.SignedString(mySigningKey)
	if nil != err {
		fmt.Printf("%s", err)
		return
	}
	fmt.Println("打印：", s)

	token, err := jwt.ParseWithClaims(s, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if nil != err {
		fmt.Printf("%s", err)
		return
	}

	//c := *token.Claims.(*jwt.MapClaims)
	fmt.Println("打印token数据：", ((*token.Claims.(*jwt.MapClaims))["username"]))
}

func ParseOne() {

}
