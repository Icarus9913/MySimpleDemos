package main

import "fmt"

type connection struct{}

type StuffClient interface {
	doStuff() error
}

type stuffClient struct {
	conn    connection
	retries int
	timeout int
}

func (s stuffClient) doStuff() error {
	return nil
}

type stuffClientOption func(*stuffClientOptions)
type stuffClientOptions struct {
	Retries int //number of times to retry the request before giving up
	Timeout int //connection timeout in seconds
}

func withRetries(r int) stuffClientOption {
	return func(o *stuffClientOptions) {
		o.Retries = r
	}
}

func withTimeOut(t int) stuffClientOption {
	return func(o *stuffClientOptions) {
		o.Timeout = t
	}
}

var defaultStuffClientOptions = stuffClientOptions{
	Retries: 1,
	Timeout: 2,
}

func newStuffClient(conn connection, opts ...stuffClientOption) StuffClient {
	options := defaultStuffClientOptions
	for _, o := range opts {
		o(&options)
	}
	return &stuffClient{
		conn:    conn,
		timeout: options.Timeout,
		retries: options.Retries,
	}
}

func main() {
	x := newStuffClient(connection{})
	fmt.Println(x) //prints &{{} 1 2}
	x = newStuffClient(
		connection{},
		withRetries(3),
	)
	fmt.Println(x) //prints &{{} 3 2}
	x = newStuffClient(
		connection{},
		withTimeOut(4),
	)
	fmt.Println(x) //prints &{{} 1 4}
}
