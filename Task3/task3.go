package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d\n", &input)
	hours := input / 3600
	mins := input / 60
	for i := 0; i < hours; i++ {
		mins = mins - 60
	}
	fmt.Printf("It is %d hours %d minutes", hours, mins)
}
