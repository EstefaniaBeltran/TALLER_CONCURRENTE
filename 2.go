package main

import "fmt"

func main() {
	ch := make(chan int)

	// Productor
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// Consumidor
	for n := range ch {
		if n%2 == 0 {
			fmt.Println(n)
		}
	}
}
