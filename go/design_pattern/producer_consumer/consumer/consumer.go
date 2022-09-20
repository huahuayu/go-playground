package consumer

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Consumer struct {
	ingestChan chan int
	jobsChan   chan int
	resultChan chan int
}

func New(ingestChan chan int, jobsChan chan int, resultChan ...chan int) *Consumer {
	c := &Consumer{
		ingestChan: ingestChan,
		jobsChan:   jobsChan,
	}
	if len(resultChan) > 0 {
		c.resultChan = resultChan[0]
	}
	return c
}

// CallbackFunc is invoked each time the external lib passes a job to us.
func (c *Consumer) CallbackFunc(job int) {
	c.ingestChan <- job
}

// WorkerFunc starts a single worker function that will range on the jobsChan until that channel closes.
func (c *Consumer) WorkerFunc(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	for job := range c.jobsChan {
		// simulate work  taking between 1-3 seconds
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
		fmt.Printf("Worker %d finished processing job %d\n", id, job)
	}
	fmt.Printf("Worker %d interrupted\n", id)
}

// Start acts as the proxy between the ingestChan and jobsChan, with a select to support graceful shutdown.
func (c *Consumer) Start(ctx context.Context) {
	for {
		select {
		case job := <-c.ingestChan:
			c.jobsChan <- job
		case <-ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobsChan!")
			close(c.jobsChan)
			fmt.Println("Consumer closed jobsChan")
			return
		}
	}
}
