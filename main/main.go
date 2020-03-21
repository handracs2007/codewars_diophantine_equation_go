package main

import (
	"fmt"
	"math"
	"reflect"
)

func Solequa(n int) [][]int {
	var result = make([][]int, 0)

	for i := 1; i <= int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			j := n / i

			if (i+j)%2 == 0 && (j-i)%4 == 0 {
				x := (i + j) / 2
				y := (j - i) / 4

				result = append(result, []int{x, y})
			}
		}
	}

	return result
}

func dotest(n int, result [][]int) {
	var realResult = Solequa(n)

	if reflect.DeepEqual(realResult, result) {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}

func main() {
	dotest(5, [][]int{{3, 1}})
	dotest(12, [][]int{{4, 1}})
	dotest(13, [][]int{{7, 3}})
	dotest(9005, [][]int{{4503, 2251}, {903, 449}})
	dotest(9008, [][]int{{1128, 562}})
	dotest(90002, [][]int{})
	fmt.Println(Solequa(9000000041))
}
