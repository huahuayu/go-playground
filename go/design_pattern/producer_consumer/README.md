# Producer-Consumer

Producer-consumer pattern with a worker pool, graceful shutdown supported.

When cancel or interrupt signal received, the program will wait for all workers to finish their jobs in consumer jobsChan.