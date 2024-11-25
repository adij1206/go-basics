package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // SIgnal that this goroutine is done+
	fmt.Printf("Worker %d started\n", i)
	fmt.Printf("Worker %d end\n", i)

	//one way
	//wg.Done()
}

func main() {
	fmt.Println("Learning about Sync and Wait group")

	// Creating WOrker group
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("Worker execution completed")
}
