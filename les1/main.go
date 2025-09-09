package main

import "fmt"

func main() {
	// Создаём карту для хранения населения городов
	cityPopulations := make(map[string]int)

	// Добавляем данные
	cityPopulations["Москва"] = 13000000
	cityPopulations["Казань"] = 1300000
	cityPopulations["Санкт-Петербург"] = 5500000

	fmt.Println("Население Москвы:", cityPopulations["Москва"])
	fmt.Println(cityPopulations)
	delete(cityPopulations, "Казань")
	fmt.Println(cityPopulations)
	if _, ok := cityPopulations["Казань"]; ok {
		fmt.Println("Город существует")
	} else {
		fmt.Println("Город не найден")
	}
}
