package main

import (
	"github.com/huahuayu/playground/go/design_pattern/producer_consumer/consumer"
	"github.com/huahuayu/playground/go/design_pattern/producer_consumer/pool"
	"github.com/huahuayu/playground/go/design_pattern/producer_consumer/producer"
	"time"
)

const (
	consumerBufferSize = 8
	workerPoolSize     = 4
)

func main() {
	consumer := consumer.New(make(chan int, 1), make(chan int, consumerBufferSize))
	producer := producer.New(consumer.CallbackFunc)
	pool := pool.New(consumer, producer, workerPoolSize)
	cancel := pool.Start()
	// Graceful shutdown after 10 seconds, but you can also use ctrl+c to interrupt the program
	time.Sleep(10 * time.Second)
	cancel()
}
