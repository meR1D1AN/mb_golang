package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func sayHello(name string) {
	fmt.Printf("Привет, %v!\n", name)
}
func add(a int, b int) int {
	return a + b
}
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Нельзя делить на ноль")
	}
	return a / b, nil
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введи своё имя: ")
	scanner.Scan()
	sayHello(scanner.Text())
	fmt.Println("Введи первое число: ")
	scanner.Scan()
	num1Str := scanner.Text()
	num1, err := strconv.Atoi(num1Str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Введи второе число: ")
	scanner.Scan()
	num2Str := scanner.Text()
	num2, err := strconv.Atoi(num2Str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Сумма:", add(num1, num2))
	result, err := divide(num1, num2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Результат деления:", result)
	for i := 0; i < 5; i++ {
		fmt.Println(i + 1)
	}
	i := 0
	for i < 5 {
		fmt.Println(5 - i)
		i++
	}

	for {
		fmt.Println("Введите фигню: ")
		scanner.Scan()
		finya := scanner.Text()
		if finya == "выход" {
			fmt.Println("Отлично, можешь выходить.")
			break
		}
	}
}
