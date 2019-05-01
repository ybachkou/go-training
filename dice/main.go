package main

import (
	"math/rand"

	"work1/output"
)

const count int = 10000

func main() {
	var sum int
	slice := make([]int, 11)
	for i := 0; i <= count; i++ {
		firstCube, secondCube := rand.Intn(6), rand.Intn(6)
		sum = firstCube + secondCube
		slice[sum]++
	}
	row := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	output.PrintRow(row)
	output.PrintRow(slice)
	sliceNew := make([]int, 0)
	for _, value := range slice {
		sliceNew = append(sliceNew, value*100/count)
	}
	output.PrintRow(sliceNew)

}
