package pool

import (
	"context"
	"fmt"
	"github.com/huahuayu/playground/go/design_pattern/producer_consumer/consumer"
	"github.com/huahuayu/playground/go/design_pattern/producer_consumer/producer"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Pool struct {
	consumer       *consumer.Consumer
	producer       *producer.Producer
	workerPoolSize int
}

func New(consumer *consumer.Consumer, producer *producer.Producer, workerPoolSize int) *Pool {
	return &Pool{
		workerPoolSize: workerPoolSize,
		consumer:       consumer,
		producer:       producer,
	}
}

// Start starts the producer and consumer, and then starts the worker pool. It also handles graceful shutdown.
func (p *Pool) Start() func() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go p.consumer.Start(ctx)
	go p.producer.Start()

	// Start workers and Add [workerPoolSize] to WaitGroup
	wg := &sync.WaitGroup{}
	wg.Add(p.workerPoolSize)
	for i := 0; i < p.workerPoolSize; i++ {
		go p.consumer.WorkerFunc(wg, i)
	}

	once := sync.Once{}
	cancel := func() {
		once.Do(func() {
			cancelFunc()
			wg.Wait()
			fmt.Println("All workers finished")
		})
	}

	go func() {
		// Handle sigterm and await termChan signal
		termChan := make(chan os.Signal)
		signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
		// Blocks here until interrupted
		<-termChan
		fmt.Println("*********************************Shutdown signal received*********************************")
		// Graceful shutdown
		cancel()
	}()
	return cancel
}
