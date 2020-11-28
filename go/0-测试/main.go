package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context, test chan struct{}, index int) {
	for true {
		select {
		case <-ctx.Done():
			fmt.Printf("stop by ctx,index:%d\n", index)
			return
		case <-test:
			fmt.Printf("stop by test,index:%d\n", index)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	test := make(chan struct{})
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			go process(ctx, test, i)
		} else {
			go process(ctx, nil, i)
		}
	}
	test <- struct{}{}
	//close(test)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}
