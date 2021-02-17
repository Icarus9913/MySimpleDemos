##PERM元模型
* subject  
> subject(sub访问实体)，object(obj访问的资源)和action(act访问方法)，  
> eft(策略结果一般指定为空 默认指定 allow) 还可以定义为deny

* Policy策略 p={sub,obj,act,eft}
> 策略一般存储到数据库，因为会有很多

* Matchers匹配规则 Request和Policy的匹配规则
> m = r.sub == p.sub && r.act == p.act && r.obj == p.obj   
> r请求 p策略  
> 这时候会把r和p按照上述描述进行匹配，从而返回匹配结果(eft)，如果不定义会返回allow， 如果定义过了，会返回我们定义过的那个结果

* Effect影响 (它决定我们是否可以放行)
> e = some(where(p.left==allow)) 这种情况下 我们的一个matchers匹配完成 得到了allow 那么这条请求将被放行  
> e = some(where(p.left==allow)) && !some(where(p.eft == deny))  
> 这里的规则是定死的

* Request请求 r= {sub, obj, act}

##role_definition 角色域
* g=_,_ 表示以角色为基础
* g=_,_,_ 表示以域为基础(多商户模式)