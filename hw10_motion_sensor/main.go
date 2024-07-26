package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func Sensor(max int64, delay time.Duration) chan int {
	c := make(chan int)
	timeout := time.After(delay)
	go func() {
		defer close(c)
		for {
			select {
			case <-timeout:
				return
			default:
				r, err := rand.Int(rand.Reader, big.NewInt(max))
				if err != nil {
					return
				}
				c <- int(r.Int64())
			}
		}
	}()
	return c
}

func Reader(depth int, inputCh chan int) chan int {
	c := make(chan int)
	counter := 0
	storage := 0
	go func() {
		defer close(c)
		for a := range inputCh {
			switch {
			case counter == depth:
				mid := float32(storage/depth) + 0.5
				c <- int(mid)
				counter = 0
				storage = 0
			default:
				storage += a
				counter++
			}
		}
	}()
	return c
}

func main() {
	emulationSensor := Sensor(1000, time.Minute)
	readData := Reader(10, emulationSensor)

	for output := range readData {
		fmt.Println("readData", output)
	}
}
