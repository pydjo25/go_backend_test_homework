package main

import (
	"fmt"
	"math"
)

// Average возвращает среднее арифметическое элементов слайса []int.
// Напишите код функции

func Average(a []int) float64 {
	sum := 0
	count := 0
	if len(a) == 0 {
		return 0
	}
	for _, ch := range a {
		sum += ch
		count++
	}
	return float64(sum / count)
}

// func Range(b []int) int {
// 	countH := 0
// 	countL := 0
// 	if len(b) < 2 {
// 		return 0
// 	}
// 	for i := 0; i < len(b); i++ {
// 		if countH == 0 && countL == 0 && b[i] > b[i+1] && b[i+1] < b[i] {
// 			countH = b[i]
// 			countL = b[i+1]
// 		}
// 		if countH != 0 && countL != 0 && b[i] < countH && b[i] < countL {
// 			countL = b[i]
// 		}
// 		if countH != 0 && countL != 0 && b[i] > countH && b[i] > countL {
// 			countH = b[i]
// 		}
// 	}
// 	s := countH - countL
// 	return s
// }

func main() {
	lists := [][]int{
		{},
		{57},
		{78, -7},
		{99, 200, 0},
		{4, 4, 4, 3},
		{102, -7, 44, -7, 102},
		{82, -23, 1, 5, 98, 100},
		{100000, 90000, 20000, 20000, 20000, 22000, 25500, 22000},
	}
	averages := []float64{
		0, 57, 36, 100, 4, 47, 44, 39938,
	}
	// ranges := []int{
	// 	0, 0, 85, 200, 1, 109, 123, 80000,
	// }

	for i, list := range lists {
		if average := math.Round(Average(list)); average != averages[i] {
			fmt.Printf("average %d: %.2f != %.2f\n", i, averages[i], average)
		}

		// if r := Range(list); r != ranges[i] {
		// 	fmt.Printf("range %d: %d != %d\n", i, ranges[i], r)
		// }
	}

	fmt.Println("Тестирование завершено")
}
