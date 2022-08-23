package _select

import (
	"sync/atomic"
	"testing"
	"time"
)

// Empty select blocks forever.
func TestEmptySelect(t *testing.T) {
	t.Log("empty select")
	select {}
	t.Log("never reach here")
}

// Single select blocks until channel is readable.
func TestSingleCaseSelect(t *testing.T) {
	t.Log("single case select")
	ch := make(chan struct{})
	time.AfterFunc(5*time.Second, func() {
		ch <- struct{}{}
	})
	select {
	case <-ch:
		t.Log("block until readable")
	}
}

// Default case make select don't block.
func TestDefaultCaseSelect(t *testing.T) {
	t.Log("default case select")
	ch := make(chan struct{})
	select {
	case <-ch:
		t.Log("never reach here")
	default:
		t.Log("fallback to default case")
	}
}

// Select will evaluate all the cases and choose a random one.
func TestMultiCaseSelect(t *testing.T) {
	t.Log("multi case select")
	ch0 := make(chan int, 1)
	ch1 := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch0 <- 0:
			t.Log("random result: ", <-ch0)
		case ch1 <- 1:
			t.Log("random result: ", <-ch1)
		}
	}
}

// Wrong break for-select loop
func TestWrongBreakForSelectLoop(t *testing.T) {
	t.Log("break for select loop")
	ch := make(chan struct{})
	counter := int32(0)
	go func() {
		for {
			ch <- struct{}{}
		}
	}()
	go func() {
		for {
			select {
			case <-ch:
				if atomic.AddInt32(&counter, 1) >= 10 {
					t.Log("break at: ", counter)
					break
				} else {
					t.Log("counter: ", counter)
				}
			}
		}
	}()
	time.Sleep(5 * time.Second)
}

// Right break for-select loop
func TestBreakForSelectLoop(t *testing.T) {
	t.Log("break for select loop")
	ch := make(chan struct{})
	counter := int32(0)
	go func() {
		for {
			ch <- struct{}{}
		}
	}()
	go func() {
	loop:
		for {
			select {
			case <-ch:
				if atomic.AddInt32(&counter, 1) >= 10 {
					t.Log("break at: ", counter)
					break loop
				} else {
					t.Log("counter: ", counter)
				}
			}
		}
	}()
	time.Sleep(5 * time.Second)
}

// Timeout for select
func TestSelectTimeout(t *testing.T) {
	t.Log("select timeout")
	ch := make(chan struct{})
loop:
	for {
		select {
		case <-ch:
			t.Log("never reach here")
		case <-time.After(2 * time.Second):
			t.Log("timeout")
			break loop
		}
	}
}

// Select a closed channel: keep printing 'chan closed', dead loop.
func TestSelectClosedChannel(t *testing.T) {
	t.Log("select closed channel")
	ch := make(chan struct{})
	time.AfterFunc(2*time.Second, func() {
		ch <- struct{}{}
	})
	time.AfterFunc(3*time.Second, func() {
		close(ch)
	})
	go func() {
		for {
			select {
			case _, ok := <-ch:
				if ok {
					t.Log("read chan")
				} else {
					t.Log("chan closed")
				}
			default:
			}
		}
	}()
	time.Sleep(4 * time.Second)
}

// Wrong priority queue implementation
func TestPriorityChanDeadLoop(t *testing.T) {
	t.Log("priority chan which cause dead loop")
	highCh := make(chan struct{}, 1)
	lowCh := make(chan struct{}, 1)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			highCh <- struct{}{}
			lowCh <- struct{}{}
		}
	}()
	go func() {
		for {
			select {
			case <-highCh:
				t.Log("highCh")
			default:
				select {
				case <-lowCh:
					t.Log("lowCh")
				default:
					// when no chan is ready, dead loop
					t.Log("default")
				}
			}
		}
	}()
	time.Sleep(10 * time.Second)
}

// Right priority queue implementation
func TestPriorityChan(t *testing.T) {
	t.Log("priority chan")
	highCh := make(chan struct{}, 1)
	lowCh := make(chan struct{}, 1)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			highCh <- struct{}{}
			lowCh <- struct{}{}
		}
	}()
	go func() {
		for {
			select {
			case <-highCh:
				t.Log("highCh")
			case <-lowCh:
			priorityLoop:
				for {
					select {
					case <-highCh:
						t.Log("highCh")
					default:
						t.Log("default")
						break priorityLoop
					}
				}
				t.Log("lowCh")
			}
		}
	}()
	time.Sleep(10 * time.Second)
}
