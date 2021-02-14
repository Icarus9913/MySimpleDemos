package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"io"
	"io/ioutil"
	"strings"
)

func main()  {
	var reader io.Reader

	app := fx.New(
		fx.Provide(func() io.Reader {
			return strings.NewReader("hello world!")
		}),
		fx.Populate(&reader),
	)
	app.Start(context.Background())
	defer app.Stop(context.Background())

	bs, err := ioutil.ReadAll(reader)
	if nil!=err{
		panic("read occur error: "+err.Error())
	}
	fmt.Printf("the result is %s \n",string(bs))
}
