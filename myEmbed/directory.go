package main

import (
	"embed"
	"io"
	"log"
	"os"
	"path"
)
/*
	内嵌目录：将static目录内嵌到二进制程序中，然后在当前目录创建static目录中的所有文件
	编译后，可以将二进制文件移到任何地方，运行后，会在当前目录输出以 embed- 开头的文件。
*/

//go:embed static
var local embed.FS

func main()  {
	fis, err := local.ReadDir("static")
	if nil!=err{
		log.Fatal(err)
	}
	for _, fi := range fis{
		in, err := local.Open(path.Join("static", fi.Name()))
		if nil!=err{
			log.Fatal(err)
		}
		out, err := os.Create("embed-" + path.Base(fi.Name()))
		if nil!=err{
			log.Fatal(err)
		}
		io.Copy(out, in)
		out.Close()
		in.Close()
		log.Println("exported","embed-"+path.Base(fi.Name()))
	}
}
