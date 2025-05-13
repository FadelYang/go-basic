package main

import "fmt"

func main() {
	// Buffered channel
	pp := make(chan int, 3)
	go func() {
		for p := range pp {
			p = <-pp
			fmt.Println("Terima data ====", p)
		}
	}()

	for i := range 10 {
		fmt.Println("Kirim data")
		pp <- i
	}
}
