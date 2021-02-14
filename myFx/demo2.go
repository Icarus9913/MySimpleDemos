package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

func main() {
	type t3 struct {
		Name string
	}

	type t4 struct {
		Age int
	}

	var (
		v1 *t3
		v2 *t4
	)

	app := fx.New(
		fx.Provide(func() *t3 { return &t3{"hello everybody!!!"} }),
		fx.Provide(func() *t4 { return &t4{2019} }),

		fx.Populate(&v1),
		fx.Populate(&v2),
	)
	app.Start(context.Background())
	defer app.Stop(context.Background())

	fmt.Printf("the result is %v , %v \n", v1.Name, v2.Age)

}
