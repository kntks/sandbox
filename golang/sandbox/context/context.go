package context

import (
	"context"
	"fmt"
	"time"
)

func leakChild() {
	go func() {
		time.Sleep(10 * time.Second)
	}()
}

func leakParent() {
	leakChild()
	time.Sleep(time.Second)
}

func child(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
				return
			default:
				continue
			}
		}
	}()
}

func parent1() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	child(ctx)
	time.Sleep(time.Second)
}

func parent2() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	child(ctx)
	time.Sleep(2 * time.Second)
}
