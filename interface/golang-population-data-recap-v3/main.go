package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]any {
	hasil := []map[string]any{}
	for i := 0; i < len(data); i++ {
		hasil = append(hasil, make(map[string]any))
		arr := strings.Split(data[i], ";")
		name := arr[0]
		age, _ := strconv.Atoi(arr[1])
		address := arr[2]
		hasil[i]["name"] = name
		hasil[i]["age"] = age
		hasil[i]["address"] = address
		if arr[3] != "" {
			height, _ := strconv.ParseFloat(arr[3], 64)
			hasil[i]["height"] = height
		}
		if arr[4] != "" {
			is_married, _ := strconv.ParseBool(arr[4])
			hasil[i]["isMarried"] = is_married
		}
	}

	return hasil
}

func main() {
	data := []string{"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
	output := PopulationData(data)
	fmt.Println(output)
}
