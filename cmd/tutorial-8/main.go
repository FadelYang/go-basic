// method
package main

import "fmt"

type Persegi struct {
	panjang, lebar int
}

func (p Persegi) menghitungLuas() int {
	luas := p.panjang * p.lebar

	return luas
}

func (p Persegi) menghitungKeliling() int {
	keliling := (2 * p.panjang) + (2 * p.lebar)

	return keliling
}

func (p *Persegi) mengubahNilaiPanjang(panjangBaru int) {
	p.panjang = panjangBaru
}

func main() {
	buku := Persegi{panjang: 10, lebar: 5} // in cm
	fmt.Printf("Luas buku adalah: %v cm\n", buku.menghitungLuas())
	fmt.Printf("Keliling buku adalah: %v cm\n\n", buku.menghitungKeliling())

	televisi := Persegi{panjang: 20, lebar: 10}
	fmt.Printf("Luas televisi adalah : %v cm\n", televisi.menghitungLuas())
	fmt.Printf("Keliling televisi adalah : %v cm\n\n", televisi.menghitungKeliling())

	fmt.Println("Mengubah nilai dari panjang televisi")
	televisi.mengubahNilaiPanjang(5)
	fmt.Printf("Panjang televisi menjadi: %v cm", televisi.panjang)
}
