package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	jumlahTiket := VIP + regular + student
	var hargaTiketVip float32 = 30 * float32(VIP)
	var hargaTiketReguler float32 = 20 * float32(regular)
	var hargaTiketStudent float32 = 10 * float32(student)
	var totalHarga float32 = hargaTiketVip + hargaTiketReguler + hargaTiketStudent
	if day%2 == 1 {
		if jumlahTiket < 5 && totalHarga >= 100 {
			return totalHarga - float32(float32(15)/float32(100)*totalHarga)
		} else if jumlahTiket >= 5 && totalHarga >= 100 {
			return totalHarga - float32(float32(25)/float32(100)*totalHarga)
		} else {
			return totalHarga
		}
	} else if day%2 == 0 {
		if jumlahTiket < 5 && totalHarga >= 100 {
			return totalHarga - float32(float32(10)/float32(100)*totalHarga)
		} else if jumlahTiket >= 5 && totalHarga >= 100 {
			return totalHarga - float32(float32(20)/float32(100)*totalHarga)
		} else {
			return totalHarga
		}
	} else {
		return 0.0
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
