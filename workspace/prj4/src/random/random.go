package main

import (
	"fmt"
	"sync"
	"time"
)

func blockFunction(i int, m chan int, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)

	fmt.Println(" Sending the data the function ")
	m <- i


	x := <-m
	fmt.Printf(" Read from channel %d\n", x)
	time.Sleep(2 * time.Second)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	host := make(chan int, 2)

	i := 0

	wg.Add(5)
	for i < 5 {
		go blockFunction(i, host, &wg)
		i++
	}
	wg.Wait()
	fmt.Println(" DONE.")

}
