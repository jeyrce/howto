package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func child(ctx context.Context) {
	for second := 0; ; second++ {
		select {
		case <-ctx.Done():
			goto over
		default:

		}
		fmt.Println("child: ", second)
		time.Sleep(time.Second)
	}
over:
	fmt.Println("child over")
}

func TestContext(t *testing.T) {
	var ctx, cancel = context.WithCancel(context.Background())
	go child(ctx)
	for i := 0; i <= 10; i++ {
		fmt.Println("parent: ", i)
		time.Sleep(time.Second * 2)
	}
	fmt.Println("parent canceled")
	cancel()
	time.Sleep(time.Second * 2)
}
