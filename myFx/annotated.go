package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

/*
	annotated提供provide注入类型,里面有Name和Group标签
type Annotated struct {
	Name   string
	Group  string
	Target interface{}
}
*/

func main() {
	type t3 struct {
		Name string
	}

	targets := struct {
		fx.In
		V1 *t3 `name:"n1"`
	}{}

	app := fx.New(
		fx.Provide(fx.Annotated{
			Name:"n1",
			Target: func() *t3{
				return &t3{"hello world"}
			},
		}),
		fx.Populate(&targets),
	)

	app.Start(context.Background())
	defer app.Stop(context.Background())

	fmt.Printf("the result is %v \n", targets.V1.Name)
}
