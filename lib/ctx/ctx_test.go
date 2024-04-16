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

func TestTimeout(t *testing.T) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Second)
		times := 0
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				times++
				t.Log("...", times, "...")
			case <-ctx.Done():
				t.Log("child: ", ctx.Err())
				return
			}
			t.Log("xxxxxxx")
		}
	}(timeout)
	//select {
	//case <-timeout.Done():
	//	t.Log("parent: ", timeout.Err())
	//}
	time.Sleep(time.Second * 4)
}
