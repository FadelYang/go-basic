package main

import (
	"fmt"
)

func probablyGetError(error bool) {
	if error {
		panic("\nAda error yang terjadi, saya panic\n")
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi panic, recover:", r)
		}
		fmt.Println("Program selesai dijalankan")
	}()

	probablyGetError(true)
	fmt.Println("Tidak ada error")
}
