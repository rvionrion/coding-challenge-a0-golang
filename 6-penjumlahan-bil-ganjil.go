/*
PENJUMLAHAN BILANGAN GANJIL

Buatlah program yang dapat melakukan penjumlahan N buah bilangan ganjil pertama ( yaitu, 1 + 3 + 5 + .... + N ). Catatan: N adalah bilangan bulat tidak negatif.

Contoh Tampilan Program :

Jumlah Bilangan Ganjil : 5 <input nilai>
Hasil : 1 + 3 + 5 + 7 + 9 = 25
*/

package main

import "fmt"

func main() {
	var jumlah_bilangan int
	var hasil_penjumlahan int

	fmt.Print("Jumlah Bilangan Ganjil : ")
	fmt.Scan(&jumlah_bilangan)

	fmt.Print("Hasil : ")

	var nilai_ganjil = 1
	for i := 0; i < jumlah_bilangan; i++ {
		hasil_penjumlahan += nilai_ganjil
		fmt.Print(nilai_ganjil)

		if i < jumlah_bilangan-1 {
			fmt.Print(" + ")
		}

		nilai_ganjil += 2
	}

	fmt.Print(" = ", hasil_penjumlahan)
}
