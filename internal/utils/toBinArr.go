package utils

import "fmt"

func ToBinArr(nums ...int) []int {
	result := make([]int, 0)

	for _, num := range nums {
		result = append(result, toBin(num)...)
	}

	return result
}

func toBin(n int) []int {
	result := make([]int, 0)

	if n == 1 {
		return []int{0, 1}
	}

	for {
		if n%2 == 0 {
			result = append([]int{0}, result...)
		} else {
			result = append([]int{1}, result...)
		}
		if n < 2 {
			break
		}
		n = n / 2

	}
	return result
}

func Test() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, toBin(i))
	}

	fmt.Println(ToBinArr(2, 2, 2))
}
