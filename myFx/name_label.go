package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

/*
	fx中的name和group标签
	provide构造函数，相同类型拥有多个值
	在Fx中未使用Name或Group标签时不允许存在多个相同类型的构造函数，一旦存在会触发panic
*/


func main() {
	type t3 struct {
		Name string
	}

	//name标签的使用
	type result struct {
		fx.Out		//proviede出，，，，

		V1 *t3 `name:"n1"`
		V2 *t3 `name:"n2"`
	}

	targets := struct {
		fx.In		//映射入

		V1 *t3 `name:"n1"`
		V2 *t3 `name:"n2"`
	}{}

	app := fx.New(
		fx.Provide(func() result {
			return result{
				V1: &t3{"hello-HELLO"},
				V2: &t3{"world-WORLD"},
			}
		}),
		fx.Populate(&targets),
	)

	app.Start(context.Background())
	defer app.Stop(context.Background())

	fmt.Printf("the result is %v, %v \n", targets.V1.Name, targets.V2.Name)
}
