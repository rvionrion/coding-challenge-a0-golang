/*
STUDI KASUS : PERSAMAAN A + B = C

Buatlah program untuk menampilkan semua penyelesaian persamaan a + b = c,
dimana c adalah 20, dengan ketentuan a, b, dan c adalah bilangan bulat >= 0
seperti contoh output dibawah ini :

A      B       C
====================
0      20      20
1      19      20
2      18      20
*/

package main

import (
	"fmt"
)

func main() {
	c := 20

	fmt.Println("A\t\tB\t\tC")
	fmt.Println("========================")

	for a := 0; a <= c; a++ {
		b := c - a
		fmt.Printf("%d\t\t%d\t\t%d\n", a, b, c)
	}
}
