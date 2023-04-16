package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("A = ")
	fmt.Scanf("%d\n", &a)
	fmt.Print("B = ")
	fmt.Scanf("%d\n", &b)
	fmt.Print("C = ")
	fmt.Scanf("%d\n", &c)

	if c*c == a*a+b*b {
		fmt.Println("Треугольник ABC — прямоугольный")
	} else {
		fmt.Println("Треугольник ABC — не прямоугольный")
	}
}
