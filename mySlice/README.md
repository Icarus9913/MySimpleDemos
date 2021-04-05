# 关于切片与数组

### 数组  
“An array type definition specifies a length and an element type”，数组由此本身的类型和长度共同决定。类型相同，长度不同，不是一个类型；长度相同，类型不同，更不是一个类型。数组是值类型，长度是不变的。  

----------------------------------------------------------- 

### 切片
A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment)”，简而言之，切片是一段数组的描述，包括指向数组的指针，数组段落的长度，以及充许的最大长度。（既然有长度了，为什么还要有最大长度，即容量，这是Go语言本身的需要，不然它的make函数没有办法工作）  

append主要用于给切片追加元素，如果该切片存储容量cap足够，就直接追加，长度len变长；如果容量不足，就会重新开辟内存，创建新的底层数组，并将之前的元素和新的元素一同拷贝进去。
此时用printf("%p")打印slice变量，会发现切片的地址变了.

```shell
//slice的底层结构
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```