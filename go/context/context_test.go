package context

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// https://www.sohamkamani.com/golang/context-cancellation-and-values/

func TestContextWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()
	select {
	case <-ctx.Done():
		t.Log("done")
	}
}

func TestContextWithTimeout(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		t.Log("done")
	}
}

func TestContextWithDeadline(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	defer cancel()
	select {
	case <-ctx.Done():
		t.Log("done")
	}
}

func TestContextWithValue(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "foo", "bar")
	v := ctx.Value("foo")
	t.Log(v)
}

func TestContextInMultiGoroutines(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	fn := func(ctx context.Context, name string, wg *sync.WaitGroup) {
		defer wg.Done()
		t0 := time.Now()
		for {
			select {
			case <-ctx.Done():
				t.Log(name, "terminate after", time.Since(t0))
				return
			default:
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
				t.Log(name, "job done at", time.Since(t0))
			}
		}
	}
	wg.Add(3)
	go fn(ctx, "fn0", &wg)
	go fn(ctx, "fn1", &wg)
	go fn(ctx, "fn2", &wg)
	wg.Wait()
}
