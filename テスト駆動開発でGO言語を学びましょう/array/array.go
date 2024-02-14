package main

import "fmt"

/*
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(Sum(numbers))
}
*/

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	//	fmt.Println(lengthOfNumbers)

	sums := make([]int, lengthOfNumbers)
	//	fmt.Println(sums)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
		//	fmt.Println(sums[i])
	}

	return sums
}

/*
	func main() {
		fmt.Println(SumAll([]int{11, 2}, []int{0, 1, 9}))
	}
*/
/*
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		// 1つ目の要素を除いたスライスを作成
		tail := numbers[1:]
		// fmt.Println(tail)

		sums = append(sums, Sum(tail))
	}

	return sums
}
*/

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		// numbersの長さが0の場合
		if len(numbers) == 0 {
			// sumsに0を追加
			sums = append(sums, 0)
			fmt.Println("sums", sums)
		} else {
			// tailsに1つ目の要素を除いたスライスを作成
			tail := numbers[1:]
			fmt.Println("tail", tail)

			// sumsにSum(tail)を追加
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
func main() {
	fmt.Println(SumAllTails([]int{}, []int{0, 1, 9}))
}
