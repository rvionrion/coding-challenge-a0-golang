/*
FAKTOR PEMBAGI

Buatlah program yang dapat menerima input bilangan bulat positif (X) dan menampilkan angka berapa saja dari angka 1 sampai X yang dapat membagi angka X tersebut tanpa sisa.

Batasan nilai:
X >= 1

*Contoh Input*
30

*Contoh Output*
Angka yang dapat membagi 30 tanpa sisa adalah:
1,2,3,5,6,10,15,30
*/
package main

import (
	"fmt"
)

func main() {
	var bilangan_x int
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scan(&bilangan_x)

  fmt.Printf("Angka yang dapat membagi %d tanpa sisa adalah:\n", bilangan_x)
	for angka := 1; angka <= bilangan_x; angka++ {
		if bilangan_x%angka == 0 {
			fmt.Print(angka)
      if bilangan_x != angka {
        fmt.Print(",")
      }
		}
	}
}
