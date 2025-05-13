// Belajar Go Routine
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2) // jumlah Go routine yang perlu ditunggu
	go CetakNama("Bagus")
	go CetakNama("Budi")
	wg.Wait() // Menunggu serangkaian proses selesai (2 Goroutine)
}

func CetakNama(nama string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Hi", nama)
		time.Sleep(time.Second)
	}
	wg.Done() // Mematikan atau menyatakan bahwa proses dari sebuah Goroutine sudah selesai
}
