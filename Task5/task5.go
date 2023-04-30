package main

import "fmt"

func printSquares(input int) {
	var count int
	n := input
	for n > 0 {
		n = n / 10
		count++
	}

	nums := []int{}

	for i := 0; i < count; i++ {
		nums = append(nums, (input%10)*(input%10))
		input = input / 10
	}

	reversed := []int{}

	for i := 0; i < count; i++ {
		reversed = append(reversed, nums[count-1-i])
		fmt.Print(reversed[i])
	}
}

func main() {
	var input int
	fmt.Scanf("%d", &input)

	printSquares(input)
}
