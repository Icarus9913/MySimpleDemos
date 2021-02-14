package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

type FxDemo struct {
	Name string
}

func NewFxDemo() FxDemo {
	return FxDemo{
		Name: "hello,world",
	}
}

func main() {
	//使用Provide将具体反射的类型添加到container中,可以按需添加任意多个构造函数
	//fx.Provide(NewFxDemo())

	//使用Populate完成变量与具体类型间的映射
	var fxd FxDemo
	//fx.Populate(fxd)

	//新建app对象(application容器包括定义注入变量、类型、不同对象lifecycle等)
	app := fx.New(
		fx.Provide(NewFxDemo), //构造函数可以任意多个,,,注意此处是函数变量，而不是完整的函数NewFxDemo()
		fx.Populate(&fxd),     //反射变量也可以任意多个，并不需要和上面构造函数对应
	)

	app.Start(context.Background())      //开启container
	defer app.Stop(context.Background()) //关闭container

	fmt.Printf("the result is %s \n", fxd.Name)
}
