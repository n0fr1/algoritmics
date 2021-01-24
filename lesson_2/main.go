//Беляев Е.Г
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	findExpFor() //2. степень через цикл
	findExpRec() //2. степень через рекурсию
}

func findExpFor() {

	var result float64

	fmt.Println("Введите число: ")
	num, err := getNumfloat64() //число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	fmt.Println("Введите степень: ")
	st, err := getNumfloat64() //степень
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	if st == 0 {
		fmt.Println("Результат: 1")
		return
	}

	result = 1
	stInt := int(st)

	if st > 0 {
		for i := 0; i < stInt; i++ {
			result = result * num
		}
	} else {
		for i := -1; i > stInt; i-- {
			result = 1 / (result * (-num))
		}
	}

	fmt.Printf("%s %v", "Результат:", result)

}

func findExpRec() {

	var result float64

	fmt.Println("Введите число: ")
	num, err := getNumfloat64() //число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	fmt.Println("Введите степень: ")
	st, err := getNumfloat64() //степень
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	if st < 0 {
		result = recMinus(num, st)
		fmt.Println(result)
		return
	}

	result = rec(num, st)
	fmt.Println(result)

}

func recMinus(num, st float64) float64 {

	if st == 0 {
		return num
	}

	return (1 / num) * (recMinus(num, st+1)) // -2
}

func rec(num, st float64) float64 {

	if st == 0 {
		return 1
	}

	return num * rec(num, st-1)
}

func getNumfloat64() (float64, error) {

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	num64, err := strconv.ParseFloat(input.Text(), 64)

	return num64, err

}
