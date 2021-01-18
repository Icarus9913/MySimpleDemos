package main

//What will be printed when the code below is executed?
func main()  {
	s := "123"
	ps := &s
	b := []byte(*ps)		//这个地方*ps->*&s就相当于只是取了一个值,值拷贝.因此后面s不管怎样变,对b都没有影响了
	pb := &b

	s += "4"
	*ps += "5"
	b[1] = '0'

	println(*ps)
	println(string(*pb))
}

//!!!!!!!!!!!!!!!!!!!!!!!!!