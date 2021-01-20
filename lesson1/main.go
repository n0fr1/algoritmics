//Беляев Е.Г
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	countIndexWeight() //1.  Рассчитать и вывести индекс массы тела.
	findFourNumMax()   //2.  Найти максимальное из четырех чисел. Массивы не использовать.
	changeNumbersA()   //3.  (a) Обмен значениями переменных с использованием третьей переменной.
	changeNumbersB()   //3.  (б) Обмен значениями переменных без использования третьей переменной.
	findArgs()         //4.  Написать программу нахождения корней заданного квадратного уравнения.
	seasonYear()       //5.  С клавиатуры вводится номер месяца. Требуется определить, к какому времени года он относится.
	inputAge()         //6.  Ввести возраст человека (от 1 до 150 лет) и вывести его «год», «года» или «лет».
	chessColor()       //7.  Вводятся координаты двух полей шахматной доски (x1,y1,x2,y2). Определить, относятся ли поля к одному цвету.
	findExp()          //8.  Ввести a и b и вывести квадраты и кубы чисел от a до b.
	findPartDivision() //9.  Даны целые положительные числа N и K. Используя только операции сложения и вычитания, найти частное от деления нацело N на K, а также остаток от этого деления
	findOddNum()       //10. Дано целое число N (> 0). С помощью операций деления нацело и взятия остатка от деления определить, имеются ли в записи числа N нечетные цифры.
	findMassiveMax()   //12. Написать функцию нахождения максимального из трех чисел.
}

//1. Рассчитать и вывести индекс массы тела.
func countIndexWeight() {

	var index float64

	fmt.Print("Введите вес (m): ")
	m, err := getNumFloat() //проверяем 1-й введенный аргумент на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	fmt.Print("Введите рост (h): ")
	h, err := getNumFloat() //проверяем 2-й введенный аргумент на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	h = h / 100 //для удобства считаем, что пользователь вводит рост в см.

	if h != 0 { //проверка на ноль
		index = m / (h * h)
	} else {
		index = 0
	}

	fmt.Printf("%s %f", "Индекс массы тела", index) //вывод сообщения
}

//2.  Найти максимальное из четырех чисел. Массивы не использовать.
func findFourNumMax() {

	var result int

	a := -25
	b := 68
	c := 10
	d := 0

	result = maxNum(a, b)
	result = maxNum(result, c)
	result = maxNum(result, d)

	fmt.Printf("%s %v", "максимум из 4-х чисел:", result) //вывод сообщения
}

//2.
func maxNum(argA, argB int) int {

	var max int

	if argA > argB {
		max = argA
	} else {
		max = argB
	}

	return max
}

//3. (a) Обмен значениями переменных с использованием третьей переменной
func changeNumbersA() {

	a := 10
	b := 5

	c := b
	b = a
	a = c
	//fmt.Printf("%s %v %s %v", "\n a = ", a, "\n b = ", b) //вывод сообщения

}

//3. (б) Обмен значениями переменных без использования третьей переменной
func changeNumbersB() {

	a := 10
	b := 5

	a, b = b, a

	//или можно так, но будет работать если оба числа положительные.
	//b = b + a
	//a = b - a
	//b = b - a

	fmt.Printf("%s %v %s %v", "\n a = ", a, "\n b = ", b) //вывод сообщения
}

//4. Написать программу нахождения корней заданного квадратного уравнения.
func findArgs() {

	str := "Решение квадратного уравнения: 5*(x*x) + 8*x + 3 = 0"
	diskrim := (8 * 8) - (4 * 5 * 3) //дискриминант = 4

	//поскольку дискриминант > 0, то уравнение имеет два корня.
	x1 := (-8 + math.Sqrt(float64(diskrim))) / (2 * 5) //корень можно вычислить только если тип данных - float64, поэтому переводим тип данных в float64
	x2 := (-8 - math.Sqrt(float64(diskrim))) / (2 * 5)

	fmt.Printf("%s %s", str, "\n")
	fmt.Printf("%s %v %s", "x1 =", x1, "\n")
	fmt.Printf("%s %v", "x2 =", x2)

}

//5. С клавиатуры вводится номер месяца. Требуется определить, к какому времени года он относится
func seasonYear() {

	fmt.Print("Введите месяц: ")
	month, err := getNumInt() //проверяем на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	if month < 1 || month > 12 {
		fmt.Println("Введен не правильный номер месяца")
		return
	}

	getSeason(month)

}

//5.
func getSeason(month int) {

	if month > 2 && month < 6 {
		fmt.Println("Весна")
		return
	}

	if month > 5 && month < 9 {
		fmt.Println("Лето")
		return
	}

	if month > 8 && month < 12 {
		fmt.Println("Осень")
		return
	}

	fmt.Println("Зима")
}

//6.  Ввести возраст человека (от 1 до 150 лет) и вывести его «год», «года» или «лет».
func inputAge() {

	fmt.Print("Введите возраст человека: ")
	age, err := getNumInt() //проверяем на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	if age < 1 || age > 150 {
		fmt.Println("Введен возраст за пределами функции")
		return
	}

	stringAge := getStringAge(age)

	fmt.Println(stringAge)
}

//6.
func getStringAge(age int) string {

	var result int

	if (age >= 5 && age < 21) || (age >= 105 && age < 121) { //в этом промежутке строго задаём.
		return "лет"
	}

	for i := 9; i > 0; i-- { //а дальше "придумываем", как можно проверить остальное.

		if (age+i)%10 == 0 { //проверяем, сколько не хватает числу до целого деления на 10.
			result = i
			break
		}
	}

	if result >= 0 && result <= 5 { //промежуток от 0 до 5 - лет
		return "лет"
	}

	if result == 9 { //если 9 - год.
		return "год"
	}

	return "года" //в остальных случаях	- года.
}

//7.  Вводятся координаты двух полей шахматной доски (x1,y1,x2,y2). Определить, относятся ли поля к одному цвету.
func chessColor() {

	fmt.Print("Введите координату x1: ")
	x1, err := getNumInt() //проверяем на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	if testPredel(x1) { //проверяем на вхождение. Для шахматной доски координаты в пределах: 1-8
		return
	}

	fmt.Print("Введите координату y1: ")
	y1, err := getNumInt() //проверяем на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	if testPredel(y1) {
		return
	}

	color1 := getColor(x1, y1) //получаем цвет 1-й клетки
	fmt.Printf("%s %s %s", "Цвет первого поля -", color1, "\n")

	fmt.Print("Введите координату x2: ")
	x2, err := getNumInt() //проверяем на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	if testPredel(x2) {
		return
	}

	fmt.Print("Введите координату у2: ")
	у2, err := getNumInt() //проверяем на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	if testPredel(у2) {
		return
	}

	color2 := getColor(x2, у2) //цвет 2-й клетки
	fmt.Printf("%s %s %s", "Цвет второго поля -", color2, "\n")

	if color1 == color2 {
		fmt.Println("Результат: одинаковые")
	} else {
		fmt.Println("Результат: разные")
	}
}

//7. проверяем пределы
func testPredel(xyz int) bool {

	if xyz < 1 || xyz > 8 {
		fmt.Println("Введена координата за пределами")
		return true
	}
	return false
}

//7.
func getColor(x1, y1 int) string {

	//определяем цвет первого поля.
	if testCoordinata(x1) && testCoordinata(y1) {
		return "черный" //черный цвет
	}

	if !(testCoordinata(x1)) && !(testCoordinata(y1)) {
		return "черный" //черный цвет
	}

	if !(testCoordinata(x1)) && testCoordinata(y1) {
		return "белый"
	}

	return "белый"
}

//7. проверяем четность/нечетность
func testCoordinata(xyz int) bool {

	if xyz%2 == 0 {
		return true
	}
	return false
}

//8.  Ввести a и b и вывести квадраты и кубы чисел от a до b.
func findExp() {

	fmt.Println("Введите аргумент a: ")
	a, err := getNumInt() //проверяем на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	fmt.Println("Введите аргумент b: ")
	b, err := getNumInt() //проверяем на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	for i := a; i <= b; i++ {
		square := i * i
		cube := square * i

		fmt.Printf("%v %s %v %s %v %s %v %s", i, "(2)=", square, ",", i, "(3)=", cube, "\n") //вывод сообщения результата
	}
}

//9. Даны целые положительные числа N и K. Используя только операции сложения и вычитания, найти частное от деления нацело N на K, а также остаток от этого деления
func findPartDivision() {

	fmt.Println("Введите аргумент n: ")
	n, err := getNumInt() //проверяем на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	fmt.Println("Введите аргумент k: ")
	k, err := getNumInt() //проверяем на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	if n < 0 || k < 0 {
		fmt.Println("Введено отрицательное число!")
		return
	}

	partDiv(n, k)
}

//9.
func partDiv(n, k int) {

	del := 0        //частное
	ostatokDel := 0 //остаток деления

	if n > k {

		for n > ostatokDel {
			ostatokDel = ostatokDel + k
			del++
		}

		if ostatokDel != n { //на цело не делится
			del--
			ostatokDel = ostatokDel - k
			ostatokDel = n - ostatokDel
			fmt.Printf("%s %v %s", "остаток от деления:", ostatokDel, "\n")
		}

	} else {
		del = 0
		ostatokDel = n
		fmt.Printf("%s %v %s", "остаток от деления:", ostatokDel, "\n")
	}

	fmt.Printf("%s %v", "частное:", del)
}

//10. Дано целое число N (> 0). С помощью операций деления нацело и взятия остатка от деления определить, имеются ли в записи числа N нечетные цифры. Если имеются, то вывести True, если нет — вывести False.
func findOddNum() {

	var result bool

	fmt.Println("Введите целое число: ")
	a, err := getNumInt() //проверяем на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	if a < 0 {
		fmt.Println("Введено отрицательное число!")
		return
	}

	result = false
	if a%10%2 == 1 {
		result = true
	} else { //в случае, если 70, 50 и т.д
		a = a / 10
		if a%2 == 1 {
			result = true
		}
	}

	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

//12. Написать функцию нахождения максимального из трех чисел.
func findMassiveMax() {

	slice := []int{-22, 0, 48}

	result := getMax(slice)
	fmt.Printf("%s %v", "\n maximum =", result) //вывод сообщения
}

//12.
func getMax(slice []int) int {

	result := slice[0]

	for i := 1; i < len(slice); i++ {

		if slice[i] > result {
			result = slice[i]
		}
	}

	return result
}

//необязательные вспомогательные функции: ввод числа и проверка на то, что это именно - число.
func getNumFloat() (float64, error) {

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	num, err := strconv.ParseFloat(input.Text(), 64)
	return num, err

}

func getNumInt() (int, error) {

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	num, err := strconv.Atoi(input.Text())
	return num, err

}
