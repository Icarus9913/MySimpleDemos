package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main()  {
	str := "nihao"
	mybyte := encode(str)
	fmt.Println(decode(mybyte))
}

func encode(str string) []byte {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(str)
	if nil!=err{
		panic(err)
	}
	return buf.Bytes()
}

func decode(mybyte []byte) string {
	var str string
	reader := bytes.NewReader(mybyte)
	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(&str)
	if nil!=err{
		panic(err)
	}
	return str
}
