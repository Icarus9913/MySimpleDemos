## Let's talk some shits with Channel!

如果chan没有缓冲,所以发送的地方只有在有目标接收的时候才能发出

注意:
如果chan不close的话,则会在以下两种情况下出现死锁.  
1.因为for range无限迭代下去;   
2.对于for循环,channel的源码下,close后会对已经close的chan变量,发送0,false
>for range或者  
>```
>for{
>  case v,ok := <-channel 
> }
> ```


