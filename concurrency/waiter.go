package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func DoWait() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		worker(i + 1)
		wg.Done()

	}
	wg.Wait()

}

func worker(i int) {
	fmt.Println("Worker ", i, "started")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Worker ", i, "finished")
}
