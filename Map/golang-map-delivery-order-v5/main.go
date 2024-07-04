package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {

	hasil := make(map[string]float32)

	for _, data := range data {
		tokens := strings.Split(data, ":")
		firstName := tokens[0]
		lastName := tokens[1]
		price, _ := strconv.ParseFloat(tokens[2], 32)
		city := tokens[3]

		if city == "BDG" {
			if !(day == "rabu" || day == "kamis" || day == "sabtu") {
				continue
			}
		}

		if city == "JKT" {
			if day == "minggu" {
				continue
			}
		}

		if city == "BKS" {
			if !(day == "selasa" || day == "kamis" || day == "jumat") {
				continue
			}
		}

		if city == "DPK" {
			if !(day == "senin" || day == "selasa") {
				continue
			}
		}

		if day == "senin" || day == "rabu" || day == "jumat" {
			price = (price * 0.1) + price
		} else if day == "selasa" || day == "kamis" || day == "sabtu" {
			price = (price * 0.05) + price
		}

		hasil[firstName+"-"+lastName] = float32(price)
	}

	return hasil
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
