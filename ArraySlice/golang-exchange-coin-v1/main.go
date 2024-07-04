package main

func ExchangeCoin(amount int) []int {
	hasil := make([]int, 0)

	for amount > 0 {
		if amount >= 1000 {
			hasil = append(hasil, 1000)
			amount -= 1000
		} else if amount >= 500 {
			hasil = append(hasil, 500)
			amount -= 500
		} else if amount >= 200 {
			hasil = append(hasil, 200)
			amount -= 200
		} else if amount >= 100 {
			hasil = append(hasil, 100)
			amount -= 100
		} else if amount >= 50 {
			hasil = append(hasil, 50)
			amount -= 50
		} else if amount >= 20 {
			hasil = append(hasil, 20)
			amount -= 20
		} else if amount >= 10 {
			hasil = append(hasil, 10)
			amount -= 10
		} else if amount >= 5 {
			hasil = append(hasil, 5)
			amount -= 5
		} else {
			hasil = append(hasil, 1)
			amount -= 1
		}
	}

	return hasil
}
