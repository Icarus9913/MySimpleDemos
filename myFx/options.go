package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"io"
	"io/ioutil"
	"strings"
)

/*
	当provide构造函数有很多的时候，或者在不同包下有很多的时候，就可以抽出来，扔进一个options集合里
*/

func main() {
	var reader io.Reader

	m := fx.Provide(func() io.Reader {
		return strings.NewReader("hello world!")
	})

	var module = fx.Options(		//options集合里存放的是不同package下的provide构造函数
		m,
	)

	app := fx.New(
		module,		//new的时候直接扔进这个module就行了，不需要再写一大堆provide
		fx.Populate(&reader),
	)
	app.Start(context.Background())
	defer app.Stop(context.Background())

	bs, err := ioutil.ReadAll(reader)
	if nil != err {
		panic("read occur error: " + err.Error())
	}
	fmt.Printf("the result is %s \n", string(bs))
}
