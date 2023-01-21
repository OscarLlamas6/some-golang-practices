package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type contextWithoutDeadline struct {
	ctx context.Context
}

func (*contextWithoutDeadline) Deadline() (time.Time, bool) { return time.Time{}, false }
func (*contextWithoutDeadline) Done() <-chan struct{}       { return nil }
func (*contextWithoutDeadline) Err() error                  { return nil }

func (l *contextWithoutDeadline) Value(key interface{}) interface{} {
	return l.ctx.Value(key)
}

func main() {
	wg := new(sync.WaitGroup)
	Task(wg)

	wg.Wait()
}

func Task(wg *sync.WaitGroup) {
	ctx := context.WithValue(context.Background(), "trace-id", "req123")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg.Add(1)
	go func(ctx context.Context) {
		bctx := &contextWithoutDeadline{ctx}
		defer wg.Done()
		time.Sleep(time.Second)

		fmt.Println("backgound process bctx.Err():", bctx.Err())
		fmt.Println(`backgound process bctx.Value("trace-id"):`, bctx.Value("trace-id"))
		fmt.Println("finish backgound process")
	}(ctx)

	fmt.Println("finish main process")
}
