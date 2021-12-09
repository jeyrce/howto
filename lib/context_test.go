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
