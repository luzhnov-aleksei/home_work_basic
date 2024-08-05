package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	C  int
	m  sync.Mutex
	wg sync.WaitGroup
}

func (counter *Counter) Incrementation(size int) {
	for i := 0; i < size; i++ {
		counter.wg.Add(1)
		go func(i int) {
			defer counter.wg.Done()
			defer counter.printer(i)
			counter.m.Lock()
			counter.C++
			counter.m.Unlock()
		}(i)
	}
	counter.wg.Wait()
}

func (counter *Counter) printer(num int) {
	counter.m.Lock()
	fmt.Printf("Goroutin %d is done \n\tCounter value: %d\n", num, counter.C)
	counter.m.Unlock()
}

func main() {
	c1 := Counter{}
	c1.C = 275

	c1.Incrementation(50)
}
