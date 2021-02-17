# 相关规则
在变量声明上方，通过 //go:embed 指令指定一个或多个符合 path.Match 模式的要嵌入的文件或目录。相关规则或使用注意如下：
1）跟其他指令一样，// 和 go:embed 之间不能有空格。（不会报错，但该指令会被编译器忽略）

2）指令和变量声明之间可以有空行或普通注释，不能有其他语句；
```shell
//go:embed message.txt

var message string
```
以上代码是允许的，不过建议紧挨着，而且建议变量声明和指令之间也别加注释，注释应该放在指令上方。

3）变量的类型只能是 string、[]byte 或 embed.FS，即使是这三个类型的别名也不行；
```shell
type mystring = string

//go:embed hello.txt
var message mystring // 编译不通过：go:embed cannot apply to var of type mystring
```

4）允许有多个 //go:embed 指令。多个文件或目录可以通过空格分隔，也可以写多个指令。比如：
```shell
//go:embed image template
//go:embed html/index.html
var content embed.FS
```

5）文件或目录使用的是相对路径，相对于指令所在 Go 源文件所在的目录，路径分隔符永远使用 /；当文件或目录名包含空格时，可以使用双引号或反引号括起来。

6）对于目录，会以该目录为根，递归的方式嵌入所有文件和子目录；

7）变量的声明可以是导出或非导出的；可以是全局也可以在函数内部；但只能是声明，不能给初始化值；
```shell
//go:embed message.txt
var message string = "" // 编译不通过：go:embed cannot apply to var with initializer
```

8）只能内嵌模块内的文件，比如 .git/* 或软链接文件无法匹配；空目录会被忽略；

9）模式不能包含 . 或 ..，也不能以 / 开始，如果要匹配当前目录所有文件，应该使用 * 而不是 .；