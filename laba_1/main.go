package main

import (
	"fmt"
	"time"
)

// Задание 5: функция возвращает сумму и разность двух вещественных чисел.
func sumAndDifference(a, b float64) (float64, float64) {
	return a + b, a - b
}

func main() {
	// Задание 1: текущая дата и время в местном часовом поясе.
	fmt.Println("1. Дата и время:", time.Now().Format("02.01.2006 15:04:05"))

	// Задание 2: явное объявление переменных разных типов.
	var age int = 20
	var height float64 = 1.75
	var name string = "Иван"
	var isStudent bool = true
	fmt.Println("\n2. Переменные разных типов:")
	fmt.Printf("int: %d\nfloat64: %.2f\nstring: %s\nbool: %t\n", age, height, name, isStudent)

	// Задание 3: краткая форма объявления переменных.
	count := 10
	price := 99.5
	message := "Привет, Go!"
	ready := true
	fmt.Println("\n3. Краткое объявление:")
	fmt.Println(count, price, message, ready)

	// Задание 4: арифметические операции с целыми числами.
	a, b := 15, 4
	fmt.Printf("\n4. Арифметика: a = %d, b = %d\n", a, b)
	fmt.Println("Сумма:", a+b)
	fmt.Println("Разность:", a-b)
	fmt.Println("Произведение:", a*b)
	if b != 0 {
		fmt.Println("Целочисленное деление:", a/b)
		fmt.Println("Остаток от деления:", a%b)
		fmt.Println("Деление с дробной частью:", float64(a)/float64(b))
	} else {
		fmt.Println("Деление на ноль невозможно")
	}

	// Задание 5: вызов функции для чисел с плавающей запятой.
	x, y := 8.5, 2.3
	sum, difference := sumAndDifference(x, y)
	fmt.Printf("\n5. Для %.2f и %.2f: сумма = %.2f, разность = %.2f\n", x, y, sum, difference)

	// Задание 6: среднее арифметическое трех чисел.
	n1, n2, n3 := 5.0, 8.0, 11.0
	average := (n1 + n2 + n3) / 3
	fmt.Printf("\n6. Среднее чисел %.2f, %.2f, %.2f = %.2f\n", n1, n2, n3, average)
}
