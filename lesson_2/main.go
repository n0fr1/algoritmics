//Беляев Е.Г
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	getStringToNum() //1. Перевод из 10-ричной в 2-двоичную систему (через цикл)
	findExpFor()     //2. степень через цикл
	findExpRec()     //2. степень через рекурсию. Плюс использование свойства четности степени.
}

//1. Перевод из 10-ричной в 2-двоичную систему (через цикл)
func getStringToNum() {

	var slice []int
	var result int
	var stringNum string

	str := "16,8,4,2,1"
	strs := strings.Split(str, ",") // "," - это разделитель строки

	for _, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
			os.Exit(1)
		}
		slice = append(slice, num) //добавляем в слайс
	}

	for i := len(slice) - 1; i >= 0; i-- { //обходим с конца и преобразуем в строку

		result = slice[i] % 2
		stringNum = stringNum + strconv.Itoa(result)
	}
	fmt.Println(stringNum)
}

//2. степень через цикл
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

	if st > 0 { //положительная степень
		for i := 0; i < stInt; i++ {
			result = result * num
		}
	} else { //отрицательная степень
		for i := -1; i > stInt; i-- {
			result = 1 / (result * (-num))
		}
	}

	fmt.Printf("%s %v", "Результат:", result)

}

//2. степень через рекурсию
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

	intSt := int(st)
	parity := intSt % 2 //проверяем на четность степени. Если степень четная - то можно использовать четность степени.
	if parity == 0 {
		st = st / 2
	}

	if st < 0 { //указана отрицательная степень

		result = recMinus(num, st)
		if parity == 0 {
			result = result * result
		}

		fmt.Println(result)
		return
	}

	result = rec(num, st) //указана положительная степень
	if parity == 0 {
		result = result * result
	}
	fmt.Println(result)
}

//отрицательная степень
func recMinus(num, st float64) float64 {

	if st == 0 {
		return num
	}

	return (1 / num) * (recMinus(num, st+1))
}

//положительная степень
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
