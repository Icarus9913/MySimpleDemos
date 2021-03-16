## Go's variable declaration in concept  

- For declared variables in Go, they need to be allocated as objects in
memory where either in the heap or on the stack  
- Heap:   
   + A global storage space
   + Where stored objects can be shared 
   + Where stored objects are managed by GC
- Stack frames:   
   + A local storage space belonging to a function
   + Each stack frame is stuck to a goroutine
   + Where stored objects are used privately
   + Where stored objects are managed by the belonging frame's lifecycle 

Heap vs Stack:  
- From a variable declaration perspective   
   + Allocating objects on the stack is faster than in the heap
   + Beacuse goroutines can fully control their stack frames.
   + No locking, no GC, and less overhead.
  

## What is the escape analysis(ESC)?
- The escape analysis is a mechanism to automatically decide whether 
a variable should be allocated in the heap or not in compile time.
   + It tries to keep variables on the stack as much as possible.
   + If a variable would be allocated in the heap, the variable is escaped(from the stack)

### ESC - concept   
A variable's construction or type doesn't determine where it lives.
Only how the variable is shared does.

- ESC would consider assignment relationships between declared variables.
- Generally, a variable scapes if :
   + its address has been captured by the address-of operand(&).
   + and at least one of the related variables has already escaped.


### How does ESC work?
- Basically, ESC determines whether variables escape or not by 
   + the data-flow analysis(shortest path analysis)
   + and other additional rules
  
### ESC - data-flow analysis
- Data-flow is a directed weighted graph
   + Constructed from the abstract syntax tree(AST)
   + It is used to represent relationships between variables.
- Vertices(locations)
   + Represent all declared variables
   + Compound types(struct,slice,and map...)is lowered to the simplest representation.
- Edges
   + Represent assignments between variables.
   + Each edge has a weight representing addressing/dereference counts(derefs)


### Data-flow analysis - process flow   
Step1. Construct locations  
- Walk throgh all functions to collect declared variables. 

Step2. Construct edges
- Walk through all functions again to collect assignments.

Step3. Analyze the built graph  
-  Iteratively walk through the built graph(based on the Bellman-Ford algorithm).
   + Start from every location.
   + Mark a variable as escaped if the source location has escaped and the relative derefs(shortest path) is -1.
   + Stop expanding a variable's incoming edges if the variable escapes.

Step4. Collect escape notes
- Walk through locations to collect the escape reasons of marked variables.


### Huge objects   
- For explicit declarations(var or :=)
   + The variables escape if their sizes are over 10MB
- For implicit declarations(new or make)  
   + The variables escape if their sizes are over 64KB
```go
package main

type hugeExplicitT struct {
	a   [3*1000*1000]int32  //12MB
}

func main()  {
  //dcl1 escapes to heap: too large for stack 
  dcl1 := hugeExplicitT{} 
  //dcl2 escapes to heap: too large for stack
  dcl2 := make([]int32,0,17*1000) //68KB
  _ = dcl1
  _ = dcl2
}
```  

```go
package main

type smallExplicitT struct {
	a   [1000*1000]int32    //4MB
}

func main()  {
  dcl3 := smallExplicitT{}
  dcl4 := make([]int32,0,15*1000)
  _ = dcl3
  _ = dcl4
}
```


#### Slice escape  
A slice variable escapes if its size of the capacity is non-constant.
```go
package main

func main()  {
  const constSize = 10
  var varSize = 10
  
  s1 := []int32{}
  // s2 escapes to heap: non-constant size
  s2 := make([]int32,varSize)
  s3 := make([]int32,constSize)
  // s4 escapses to heap: non-constant size
  s4 := make([]int32,varSize,varSize)
  s5 := make([]int32,varSize,constSize)
  // s6 escapes to heap: no-constant size
  s6 := make([]int32,constSize, varSize)
  s7 := make([]int32,constSize, constSize)
}
```

#### Map escape
- A variable escapse if it is referenced by a map's key or value.
- The escape happens no matter the map escape or not.
```go
package main

func map1()  {
  m1 := make(map[int]int)
  k1 := 0
  v1 := 0
  m1[k1]=v1
}

func map2()  {
  m2 := make(map[*int]*int)
  k2 := 0   //escapes to heap: key of map put
  v2 := 0   //escapes to heap
  m2[&k2] = &v2
}

func map3()  {
  m3 := make(map[interface{}]interface{})
  k3 := 0   //escapes to heap: key of map put
  v3 := 0   //escapes to heap
  m3[&k3] = &v3 //interface-converted happens
}
```

#### Return values   
Returning values is a backward behavior that  
- the referenced variables escape if the return values are pointers
- the values escape if they are map or slice
```go
package main

func f1() **int {
  //t escapes to heap 
  t := 0
  //x1 escapes to heap
  x1 := &t  
  return &x1
}

func f2() *int {
  // t escapes to heap
  t := 0
  x2 := &t
  return x2
}

func f3() int {
  t := 0
  x3 := t
  return x3
}

func f4() map[string]int {
  // kv escapes to heap
  kv := make(map[string]int)
  return kv
}

func f5() []int {
  // s escapes to heap
  s := []int{}
  return s
}
```

#### Input parameters
Passing arguments is a forward behavior that
- the arguments escape if input parameters have leaked(to heap)
```go
package main

func f1(x1 *int) **int {
  // x1 escapes to heap: parameters leaking
  return &x1
}

func f2(x2 *int) *int {
  return x2
}

func f3(x3 *int) int {
  return *x3
}

func main()  {
  v1 := 1   //v1 escapesto heap
  f1(&v1)
  
  v2 := 1
  f2(&v2)
  
  v3 := 1
  f3(&v3)
}
```

#### Closure function
A variable escapes if
- the source variable is captured by a closure function
- and their relationship is address-of(derefs=-1)
```go
package main

func closure1()  {
  var x *int
  func(x1 *int){
  	func(x2 *int){
  		func(x3 *int){
  			y := 1
  			x3 = &y
        }(x2)
    }(x1)
  }(x)

  _ = x
}

func closure2()  {
  var x *int
  func(){
  	func(){
  		func(){
  			//y escapes to heap
  			y := 1
  			//x is captured by a closure
  			x = &y
        }()
    }()
  }()
  
  _ = x
}
```


### observations
Through understanding the concept of ESC, we can find that
- variables usually escape
   + when their addresses are captured by other variables.
   + when ESC does not know their object sizes in compile time.
- And passing arguments to a function is safer than returning values from the function  

So, the first and most important suggestion is :
try not to use pointers as much as possible