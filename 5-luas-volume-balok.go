/*
LUAS DAN VOLUME BALOK

Buatlah Program untuk menghitung Luas dan Volume suatu bangun ruang Balok, dimana diketahui Panjang = p; Lebar = l; Tinggi = t dan Jalankan hasil programnnya untuk Panjang = 10; Lebar = 8 dan Tinggi = 5.

           ________________
  l = 8  /|               /|
        /_|______________/ |
       |  |_____________|__|
t = 5  | /              | /
       |/_______________|/
              p = 10

Ditanyakan :
Luas Permukaan = ?
Volume = ?
*/

package main

import (
	"fmt"
)

func main() {
	var panjang = 10
	var lebar = 8
	var tinggi = 5

	var volume = panjang * lebar * tinggi
	var luas_permukaan = 2 * ((panjang * lebar) + (panjang * tinggi) + (tinggi * lebar))

	fmt.Println("Volume = ", volume)
	fmt.Println("Luas Permukaan = ", luas_permukaan)
}
