package lib

import (
	"context"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()
	go testTimeout(ctx, time.Second*2)
	for {
		select {
		case <-ctx.Done():
			t.Log(ctx.Err())
		}
	}
}

func testTimeout(ctx context.Context, timeout time.Duration) {
	time.Sleep(timeout)
	ctx.Deadline()
}

func do(ctx context.Context) {
	context.WithValue(ctx, "date", "2024-11-22")
}

func TestValueContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), "author", "jeyrce")
	do(ctx)
	t.Log(ctx.Value("date"))
}
