package main

import "fmt"

/*
	简单的拦截器,就是在执行业务之前,进行拦截然后执行其他的操作. 可以理解为中间件?
*/

type account interface {
	query(id string) int
	update(id string, value int)
}

type accountImpl struct {
	id    string
	name  string
	value int
}

func (a *accountImpl) query(_ string) int {
	fmt.Println("AccountImpl.Query")
	return 100
}

func (a *accountImpl) update(_ string, _ int) {
	fmt.Println("AccountImpl.Update")
}

var newA = func(id, name string, value int) account {
	return &accountImpl{id, name, value}
}

func main() {
	id := "1001"
	a := newA(id, "icarus", 18)
	a.query(id)
	a.update(id, 22)
}

type proxy struct {
	account account
}

func (p *proxy) query(id string) int {
	fmt.Println("Proxy.Query begin")
	value := p.account.query(id)
	fmt.Println("Proxy.Query end")
	return value
}

func (p *proxy) update(id string, value int) {
	fmt.Println("Proxy.Update begin")
	p.account.update(id, value)
	fmt.Println("Proxy.Update end")
}

func init() {
	newA = func(id, name string, value int) account {
		a := &accountImpl{id, name, value}
		p := &proxy{a}
		return p
	}
}
