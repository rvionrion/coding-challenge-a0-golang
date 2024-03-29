/*
SALURAN TV

Diberikan daftar saluran TV seperti berikut:
 1=TVRI         4=ANTV
 2=RCTI         5=TRANSTV
 3=SCTV         6=METROTV

Buatlah program yang dapat menerima inputan nomor saluran TV, kemudian program menampilkan nama stasiun penyiaran.
Jika nomor yang diinputkan tidak ada didalam daftar saluran, maka tampilkan "Saluran tidak tersedia".

*/

package main

import (
	"fmt"
)

func main() {
	var daftar_saluran []string = []string{"TVRI", "RCTI", "SCTV", "ANTV", "TRANSTV", "METROTV"}

	var saluran int

	fmt.Print("Masukkan Saluran: ")
	fmt.Scan(&saluran)

	var index_saluran = saluran - 1

	if index_saluran < 0 || index_saluran > len(daftar_saluran)-1 {
		fmt.Println("Saluran tidak tersedia")
	} else {
		fmt.Println("Nama Stasiun Penyiaran:", daftar_saluran[index_saluran])
	}
}
