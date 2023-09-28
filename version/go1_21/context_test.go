package go1_21

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	const shortDuration = 1 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestContext1(t *testing.T) {
	tooSlow := fmt.Errorf("too slow!")
	_, cancel := context.WithTimeoutCause(context.Background(), 1*time.Second, tooSlow)
	time.Sleep(2 * time.Second)
	cancel()
}
