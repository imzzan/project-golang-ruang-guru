package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	hasil := [][]string{}
	for key, value := range mapData {
		element := []string{key, value}
		hasil = append(hasil, element)
	}

	return hasil
}

func main() {
	mapData := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	fmt.Println(MapToSlice(mapData))

}
