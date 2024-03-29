/*
POSITIF, NEGATIF, DAN NOL

Diberikan deretan bilangan seperti berikut:

10, -4, 27, 0, -31, 8

Buat program yang dapat menentukan jenis setiap bilangan diatas, apakah bilangan tersebut termasuk kedalam bilangan positif, bilangan negatif atau bilangan nol.

*Contoh Output*
10 adalah bilangan positif
-4 adalah bilangan negatif
27 adalah bilangan positif
0 adalah bilangan nol
-31 adalah bilangan negatif
8 adalah bilangan positif
*/
package main

import (
	"fmt"
)

func main() {
  var data [6]int = [6]int{10, -4, 27, 0, -31, 8}

  for x := 0; x < 6; x++ {
    var bilangan = data[x]
    if bilangan == 0 {
      fmt.Println(bilangan, "adalah bilangan nol")
    } else if bilangan > 0 {
      fmt.Println(bilangan, "adalah bilangan positif")
    } else {
      fmt.Println(bilangan, "adalah bilangan negatif")
    }
  }
}
