/**
 * =================
 * Belanja di warung
 * =================
 *
 * [Description]
 * Ibu meminta tolong kamu untuk belanja di warung.
 * Ibu akan memberikan sejumlah uang dan satu jenis barang yang harus dibeli.
 * Warung tersebut hanya menjual:
 * 1. Wafer seharga 15000
 * 2. Bayam seharga 5000
 * 3. Indomie seharga 2000
 *
 * [Instruction]
 * Hitunglah berapa barang yang dapat dibeli dengan jumlah uang yang telah
 * diberikan dan sisa kembalian uang jika ada.
 *
 * [Example]
 * var @belanja = 'wafer'
 * var @uang = 40000
 *
 * OUTPUT 1. jika ada kembalian
 * 'Kamu membeli 2 wafer dan memiliki kembalian sebanyak 10000'
 *
 * OUTPUT 2. jika tidak ada kembalian
 * 'Kamu membeli 2 wafer dan tidak mendapat kembalian'
 *
 * Karena ibu memberikan uang 40000 dan ingin membeli wafer. Maka kamu akan mendapatkan 2 wafer
 * (2 x 15000) serta memiliki kembalian 10000.
 *
 * [Asumsi]
 * - Uang akan selalu lebih besar atau sama dengan dari harga barang yang akan dibeli.
 * - Barang yang dibeli hanya wafer, bayam, atau indomie
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	name := "Indomie"
	var money float64 = 40000
	var count float64 = 0
	var resid float64 = 0

	switch name {
	case "Wafer":
		switch price := 1500; {
		case money > float64(price):
			count = math.Floor(money / float64(price))
			resid = money - (count * float64(price))

			switch value := 0; {
			case resid > float64(value):
				fmt.Println("Kamu membeli", count, "wafer dan memiliki kembalian sebanyak", resid)
			case resid == float64(value):
				fmt.Println("'Kamu membeli", count, "wafer dan tidak mendapat kembalian")
			}

		case money < float64(price):
			fmt.Println("Uang tidak cukup untuk membeli wafer")
		}

	case "Bayam":
		switch price := 5000; {
		case money > float64(price):
			count = math.Floor(money / float64(price))
			resid = money - (count * float64(price))

			switch value := 0; {
			case resid > float64(value):
				fmt.Println("Kamu membeli", count, "bayam dan memiliki kembalian sebanyak", resid)
			case resid == float64(value):
				fmt.Println("'Kamu membeli", count, "bayam dan tidak mendapat kembalian")
			}

		case money < float64(price):
			fmt.Println("Uang tidak cukup untuk membeli bayam")
		}

	case "Indomie":
		switch price := 2000; {
		case money > float64(price):
			count = math.Floor(money / float64(price))
			resid = money - (count * float64(price))

			switch value := 0; {
			case resid > float64(value):
				fmt.Println("Kamu membeli", count, "indomie dan memiliki kembalian sebanyak", resid)
			case resid == float64(value):
				fmt.Println("'Kamu membeli", count, "indomie dan tidak mendapat kembalian")
			}

		case money < float64(price):
			fmt.Println("Uang tidak cukup untuk membeli indomie")
		}
	}
}
