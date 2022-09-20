package producer

import "time"

type Producer struct {
	callbackFunc func(job int)
}

func New(callbackFunc func(job int)) *Producer {
	return &Producer{
		callbackFunc: callbackFunc,
	}
}

// Start starts the producer to enqueue jobs
func (p *Producer) Start() {
	job := 0
	for {
		p.callbackFunc(job)
		job++
		time.Sleep(time.Millisecond * 100)
	}
}
