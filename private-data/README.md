## Golang指针  
- *类型:普通指针类型，用于传递对象地址，不能进行指针运算。
- unsafe.Pointer:通用指针类型，用于转换不同类型的指针，不能进行指针运算，不能读取内存存储的值（必须转换到某一类型的普通指针）。
- uintptr:用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收。  

> unsafe.Pointer 是桥梁，可以让任意类型的指针实现相互转换，也可以将任意类型的指针转换为 uintptr 进行指针运算。  
> <br>
> unsafe.Pointer 不能参与指针运算，比如你要在某个指针地址上加上一个偏移量，Pointer是不能做这个运算的，那么谁可以呢?  
> 就是uintptr类型了，只要将Pointer类型转换成uintptr类型，做完加减法后，转换成Pointer，通过*操作，取值，修改值，随意。  
> <br>
>  总结：unsafe.Pointer 可以让你的变量在不同的普通指针类型转来转去，也就是表示为任意可寻址的指针类型。而 uintptr 常用于与 unsafe.Pointer 打配合，用于做指针运算。

#### unsafe.Pointer    
  unsafe.Pointer称为通用指针，官方文档对该类型有四个重要描述： 
  * 任何类型的指针都可以被转化为Pointer
  * Pointer可以被转化为任何类型的指针
  * uintptr可以被转化为Pointer
  * Pointer可以被转化为uintptr

unsafe.Pointer是特别定义的一种指针类型（译注：类似C语言中的void类型的指针），在golang中是用于各种指针相互转换的桥梁，它可以包含任意类型变量的地址。
***当然，我们不可以直接通过*p来获取unsafe.Pointer指针指向的真实变量的值，因为我们并不知道变量的具体类型.***  
和普通指针一样，unsafe.Pointer指针也是可以比较的，并且支持和nil常量比较判断是否为空指针。

  





