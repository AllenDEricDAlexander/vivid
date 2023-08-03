package main

import (
	"fmt"
	"time"
	"user/model"
	"user/repo/persistence"
)

func main() {
	user := model.User{ID: 3, Nickname: "nickname", Signature: "hello", Password: "hello", CreateTime: time.Now(), UpdateTime: time.Now()}
	context := &persistence.UserDAOContext{}
	context.SetDB(&persistence.RedisDAO{})
	err := context.DoCreate(user)
	fmt.Println(err)
}
