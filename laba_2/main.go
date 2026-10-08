package main

import (
	"fmt"
	"unicode/utf8"
)

// Задание 2: определение знака числа.
func numberSign(number int) string {
	if number > 0 {
		return "Positive"
	} else if number < 0 {
		return "Negative"
	}
	return "Zero"
}

// Задание 4: количество символов Unicode, а не байтов.
func stringLength(text string) int {
	return utf8.RuneCountInString(text)
}

// Задание 5: структура прямоугольника и метод площади.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Задание 6: дробная часть среднего сохраняется.
func average(a, b int) float64 {
	return (float64(a) + float64(b)) / 2
}

func main() {
	// Задание 1: ввод целого числа и проверка четности.
	var number int
	fmt.Print("1. Введите целое число: ")
	if _, err := fmt.Scan(&number); err != nil {
		fmt.Println("Ошибка: необходимо ввести целое число.")
		return
	}
	if number%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}

	fmt.Println("\n2. Знак введенного числа:", numberSign(number))

	// Задание 3: цикл for от 1 до 10 включительно.
	fmt.Println("\n3. Числа от 1 до 10:")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	text := "Привет, Go!"
	fmt.Printf("\n4. Длина строки %q: %d\n", text, stringLength(text))

	rectangle := Rectangle{Width: 5, Height: 3}
	fmt.Printf("\n5. Площадь прямоугольника %.2f × %.2f: %.2f\n", rectangle.Width, rectangle.Height, rectangle.Area())

	a, b := 7, 10
	fmt.Printf("\n6. Среднее чисел %d и %d: %.2f\n", a, b, average(a, b))
}
