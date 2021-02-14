package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

func main()  {
	type t3 struct {
		Name string
	}

	//使用group标签
	type result struct {
		fx.Out

		V1 *t3 `group:"g"`
		V2 *t3 `group:"g"`
	}

	targets := struct {
		fx.In

		Group []*t3	`group:"g"`
	}{}

	app := fx.New(
		fx.Provide(func()result {
			return result{
				V1: &t3{"hello-000"},
				V2: &t3{"world-www"},
			}
		}),
		fx.Populate(&targets),
	)

	app.Start(context.Background())
	defer app.Stop(context.Background())

	for _,t := range targets.Group{
		fmt.Printf("the result is %v \n",t.Name)
	}
}
