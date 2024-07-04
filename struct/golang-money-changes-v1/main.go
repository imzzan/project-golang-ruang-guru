package main

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {

	pecahanKembalian := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	total := []int{}
	for _, item := range products {
		total = append(total, item.Price+item.Tax)
	}

	totalBayar := 0
	for _, item := range total {
		totalBayar += item
	}

	sisaBelanja := amount - totalBayar
	var totalKembalian []int
	for i := 0; i < len(pecahanKembalian); i++ {
		for sisaBelanja >= pecahanKembalian[i] {
			if sisaBelanja >= pecahanKembalian[i] {
				totalKembalian = append(totalKembalian, pecahanKembalian[i])
			}
			sisaBelanja -= pecahanKembalian[i]
		}
	}

	if totalKembalian == nil {
		return []int{}
	}

	return totalKembalian
}
