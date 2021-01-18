package main

//modify the code below,to exit the outer for-loop
func main()  {
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			print(i,",",j," ")
			break
		}
		println()
	}
}

//my fucking change
func init()  {
	var j int
FOO:
	for i:=0;i<3;i++{
		for j=0;j<3;j++ {
			print(i,",",j," ")
			break FOO
		}
	}
	println()
}
