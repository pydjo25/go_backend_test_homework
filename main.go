package main

import (
	"fmt"
	"sort" // или "slices"
)

// Median возвращает медиану числовой последовательности.
// Напишите код функции
// ...
func Median(a []int) int {
	if len(a) == 0 {
		return 0
	}

	if len(a) < 2 {
		return a[len(a)-1]
	}

	aCp := make([]int, len(a))
	copy(aCp, a)
	sort.Ints(aCp)

	if len(aCp) == 2 {
		return (aCp[0] + aCp[1]) / 2
	}

	num := len(aCp)
	flag := num / 2

	if len(aCp)%2 == 0 {
		return (aCp[flag] + aCp[flag+1]) / 2
	} else {
		return aCp[flag]
	}

}

func main() {
	lists := [][]int{
		// {},
		// {57},
		// {78, -7},
		// {99, 200, 0},
		// {4, 4, 4, 3},
		// {102, -7, 44, -7, 102},
		{82, -23, 1, 5, 98, 100},
		// {100000, 90000, 20000, 20000, 20000, 22000, 25500, 22000},
	}
	medians := []int{
		0, 57, 35, 99, 4, 44, 43, 22000,
	}

	for i, list := range lists {
		if median := Median(list); median != medians[i] {
			fmt.Printf("median %d: %d != %d\n", i, medians[i], median)
		}
	}
	fmt.Println("Тестирование завершено")
}
