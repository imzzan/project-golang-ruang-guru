package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

var c = internal.Calculator{}

func AdvanceCalculator(calculate string) float32 {
	split := strings.Split(calculate, " ")
	base, _ := strconv.ParseFloat(split[0], 64)
	Cal := internal.NewCalculator(float32(base))
	for i := 0; i < len(split); i++ {
		if split[i] == "*" {
			num, _ := strconv.ParseFloat(split[i+1], 64)
			Cal.Multiply(float32(num))
		} else if split[i] == "/" {
			num, _ := strconv.ParseFloat(split[i+1], 64)
			Cal.Divide(float32(num))
		} else if split[i] == "+" {
			num, _ := strconv.ParseFloat(split[i+1], 64)
			Cal.Add(float32(num))
		} else if split[i] == "-" {
			num, _ := strconv.ParseFloat(split[i+1], 64)
			Cal.Subtract(float32(num))
		}
	}
	return Cal.Result() // TODO: replace this
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
