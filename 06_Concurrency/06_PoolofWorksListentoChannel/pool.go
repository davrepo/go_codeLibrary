package main

import (
	"fmt"
	"log"
	"runtime"
	"sort"
	"sync"
	"time"
)

func median(values []float64) float64 {
	// returns the median value of a vector of floats
	nums := make([]float64, len(values))
	copy(nums, values)
	sort.Float64s(nums)
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i]
	}

	return (nums[i-1] + nums[i]) / 2.0
}

func poolWorker(ch <-chan []float64, wg *sync.WaitGroup) { // receive-only channel
	// receive a vector from channel assign to values
	// loop will terminate when channel is closed
	for values := range ch {
		m := median(values)
		log.Printf("median %v -> %f", values, m)
		// the # of times wg.Done() will run depends on the # of times
		// data is sent through the channel, and not number of goroutines started.
		// Goroutine will wait and do nothing if there is no data to be received from channel.
		// i.e. wg.Done() will run 5 times, even though there are only 2 goroutines started.
		wg.Done()
	}
	log.Printf("shutting down")
}

func multiDot(vectors [][]float64) {
	var wg sync.WaitGroup
	wg.Add(len(vectors)) // 5 number of jobs
	ch := make(chan []float64)

	// pooled workers
	// start a number of goroutines to match the number of CPUs
	for i := 0; i < runtime.NumCPU(); i++ {
		go poolWorker(ch, &wg)
	}
	for _, vec := range vectors {
		// send a vector to channel
		ch <- vec
	}

	// will 5 jobs are done (wg.Done() is called 5 times), the WaitGroup will unblock
	wg.Wait()
	// close the channel so for loop in poolWorker() will terminate
	close(ch)
}

func main() {
	vectors := [][]float64{
		{1.1, 2.2, 3.3},
		{2.2, 3.3, 4.4},
		{3.3, 4.4, 5.5},
		{4.4, 5.5, 6.6},
		{5.5, 6.6, 7.7},
	}
	multiDot(vectors)
	// without this line main goroutine may terminate before the worker goroutines have a chance
	// to log their "shutting down" message. However, the actual computation of the medians
	// won't be affected b/c of placement of wg.Done() in the poolWorker() function.
	time.Sleep(10 * time.Millisecond)
	fmt.Println("DONE")
}
