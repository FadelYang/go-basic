package main

import "fmt"

func main() {
	// deklarasi new(int32) untuk memberi tahu bahwa kita perlu mengalokasikan memoru sebesar int32
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value p variable point is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)

	p = &i
	*p = 12

	fmt.Printf("\n\n %v", i)

	fmt.Printf("\n\nThe value p variable point is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)

	// Versi sederhana
	var a = 12

	// b adalah variable pointer yang menyimpan alamat a
	// bisa dibilang b menunjuk alamat a
	var b *int = &a

	fmt.Println("\nvalue dari a", a)

	// sehingga ketika diprint seperti ini, atau diprint dengan cara melihat apa yang b tunjuk
	// hasilnya adalah value dari a
	fmt.Println("value dari b", *b)
}
