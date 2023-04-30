package main

import "fmt"

func addStars(input string) {
	for i := 0; i < len(input)-1; i++ {
		fmt.Printf("%c*", input[i])
	}
	fmt.Printf("%c", input[len(input)-1])
}

func main() {
	var input string
	fmt.Scanf("%s", &input)
	addStars(input)
}
