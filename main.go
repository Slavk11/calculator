package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите выражение: ")

	result, err := calculate(os.Stdin)
	if err != nil {
		fmt.Println("Ошибка", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
