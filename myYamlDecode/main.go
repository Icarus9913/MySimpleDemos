package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Hello struct {
	MyTest myTest `yaml:"my.Test"`
}

type myTest struct {
	Title    string   `yaml:"title"`
	Category string   `yaml:"category"`
	Tag      []string `yaml:"tag"`
}

func main() {

	/*
		如果使用os.open则返回一个*os.File, 然后使用bufio包转成*bufio.Reader类型,
		解析的时候使用yaml.NewDecoder().Decode(&result)就行了
	*/

	//file, err := os.Open("myTest.yml")
	//if nil != err {
	//	panic(err)
	//}
	//defer file.Close()
	//reader := bufio.NewReader(file)
	//decoder := yaml.NewDecoder(reader)
	//err = decoder.Decode(result)

	//直接返回byte数组
	readFile, err := ioutil.ReadFile("myTest.yml")
	if nil != err {
		panic(err)
	}

	result := new(Hello)

	err = yaml.Unmarshal(readFile, result)
	if nil != err {
		panic(err)
	}
	fmt.Println(result)
}
