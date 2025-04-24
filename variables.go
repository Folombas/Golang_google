package main

import "fmt"

func main() {

	// мой любимый цвет
	var favoriteColor = "Purple"
	fmt.Println("Мой любимый цвет", favoriteColor)

	// мой год рождения и возраст
	birthYear, ageInYear := 1987, 37
	fmt.Println("Родился в", birthYear, "году", "мне", ageInYear, "года")

	// блочное присваивание
	var (
		firstInitial = "A"
		surnameInitial = "F"
)
	fmt.Println("Инициалы =", firstInitial, surnameInitial)


	var ageInDays int
	ageInDays = 365 * ageInYear
	fmt.Println("Мой возраст в днях =", ageInDays)

	// fmt.Println() для вывода информации
	// Вычитание -
	// Сложение +
	// Умножение *
	// Деление /
}
