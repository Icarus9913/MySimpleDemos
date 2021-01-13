package main

import (
	"context"
	"fmt"
)

func main() {
	ProcessRequest("admin", "123456")
}

type ctxKey int

const (
	ctxUserName ctxKey = iota
	ctxPassWord
)

func UserName(c context.Context) string {
	return c.Value(ctxUserName).(string)
}

func PassWord(c context.Context) string {
	return c.Value(ctxPassWord).(string)
}

func ProcessRequest(UserName, PassWord string) {
	ctx := context.WithValue(context.Background(), ctxUserName, UserName)
	ctx = context.WithValue(ctx, ctxPassWord, PassWord)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf("处理响应 用户名:%v 密码:%v\n",
		UserName(ctx),
		PassWord(ctx),
	)
}
