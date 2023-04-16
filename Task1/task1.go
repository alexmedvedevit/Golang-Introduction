package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d", &input)
	if input > 999 || input < 100 {
		fmt.Print("Ошибка: Введённое число должно быть трёхзначным!")
	} else if input%10 == 0 {
		fmt.Print("Ошибка: Введённое число не должно заканчиваться на 0!")
	} else {
		output := 0
		for input > 0 {
			temp := input % 10
			output = output*10 + temp
			input = input / 10
		}
		fmt.Print(output)
	}

}
