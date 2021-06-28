package async

import (
	"context"
)

type Future interface {
	Await() interface{}
}

type future struct {
	await func(ctx context.Context) interface{}
}

func (f future) Await() interface{} {
	return f.await(context.Background())
}

func Exec(f func() interface{}) Future {
	var result interface{}
	c := make(chan struct{})
	go func() {
		defer close(c)
		result = f()
	}()

	return future{
		await: func(ctx context.Context) interface{} {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				return result
			}
		}}
}
