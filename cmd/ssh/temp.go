package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func contextCancellation(){
	var wg sync.WaitGroup

	context.WithValue(context.TODO(), "a", 123)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	for w:=0; w < 5; w++{
		wg.Add(1)

		randServer := rand.Int32N(5)

		go func(n int, crashes bool, ctx context.Context, cancel context.CancelFunc){
			defer func(){
				fmt.Printf("Server %d shutting down\n", n)
				wg.Done()
			}()

			fmt.Printf("Server %d started\n", n)

			var crashDuration int32

			if crashes {
				crashDuration = rand.Int32N(10)
			} else {
				crashDuration = 1000
			}
			
			timer := time.NewTimer(time.Duration(crashDuration) * time.Second)

			for {
				select {
				case <-timer.C:
					fmt.Println("-----------------------------------")
					fmt.Printf("Server %d has crashed\n", n)
					cancel()
				case <-ctx.Done():
					return
				default:
					fmt.Printf("Server %d doing work\n", n)
					time.Sleep(5 * time.Second)
				}
			}
			
		}(w, randServer == int32(w), ctx, cancel)
	}

	wg.Wait()
}

func blockTil(duration int) {
	done := make(chan int)

	go func() {
		time.Sleep(time.Duration(duration) * time.Second)
		done <- 0
	}()
	
	<-done
	fmt.Printf("Unblocked after %d seconds\n", duration)
}

func semaphoreEx() {
	var wg sync.WaitGroup
	sem := make(chan int, 10)

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			sem <- 0

			duration := rand.Int32N(5)
			time.Sleep(time.Duration(duration) * time.Second)
			fmt.Printf("Worker %d completed after %d seconds\n", i, duration)

			<-sem
		}(i)
	}

	wg.Wait()
}