//Беляев Е.Г
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

//Variants is ...
type Variants struct {
	A, B, C, D, E int64
}

//Choise is ...
type Choise struct {
	Name  string
	Index int
}

func main() {

	var a, b int64 // //3. Обход конём шахматной доски 5х5. Наброски

	// fmt.Println("Укажите стартовые клетки для положения коня")
	// fmt.Println("Введите первый номер: ")
	// a, err := getNumInt() //проверяем на число
	// if err != nil {
	// 	fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
	// 	os.Exit(1)
	// }

	// fmt.Println("Введите второй номер: ")
	// b, err := getNumInt() //проверяем на число
	// if err != nil {
	// 	fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
	// 	os.Exit(1)
	// }

	a = 1 //для теста
	b = 2

	var v = []Variants{}
	v = []Variants{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	testFunc(&v, a, b)

	// v.slA = []int{1, 2, 3, 4, 5}
	// v.slB = []int{6, 7, 8, 9, 10}
	// v.slC = []int{11, 12, 13, 14, 15}
	// v.slD = []int{16, 17, 18, 19, 20}
	// v.slE = []int{21, 22, 23, 24, 25}

	//test(&v, a, b)
	//test2(v, a, b)

	// var slice []int //занятые позиции
	// slice = append(slice, a)
	// slice = append(slice, b)

	// //нумеруем строки - для случая 5x5
	// slA := []int{1, 2, 3, 4, 5}
	// slB := []int{6, 7, 8, 9, 10}
	// slC := []int{11, 12, 13, 14, 15}
	// slD := []int{16, 17, 18, 19, 20}
	// slE := []int{21, 22, 23, 24, 25}

	//проверить на правильный ввод стартовой пары
	// if (a-b) == 1 || (b-a) == 1 { //принадлежат одной строке
	// 	fmt.Println("ок")
	// 	return
	// }

	// result := testInput(slA, slB, slC, slD, slE, a, b)
	// if !result {
	// 	fmt.Println("Не правильно указана стартовая пара")
	// 	return
	// }

	//проверяем ходы со стороны а
	//if

}

func testFunc(v *[]Variants, a int64, b int64) bool {

	for i := 0; i < len(*v); i++ {
		val := reflect.ValueOf((*v)[i])

		for k := 0; k < val.NumField(); k++ {

			value := val.Field(k).Interface().(int64)

			if value == a {

				typeField := val.Type().Field(k)
				p1 := Choise{Name: typeField.Name, Index: k}
				fmt.Println(p1)

			}

			if value == b {
				typeField := val.Type().Field(k)
				p2 := Choise{Name: typeField.Name, Index: k}
				fmt.Println(p2)
			}

	}

}

func getNumInt() (int, error) {

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	num, err := strconv.Atoi(input.Text())
	return num, err

}

func test2(v Variants, a, b int) {

	val := reflect.ValueOf(v)
	for i := 0; i < val.NumField(); i++ {

		//for _, zn := range val.Field(i) {

		//}
		// result := val.Field(i).Interface().(Variants)
		// fmt.Println(result)
	}
}

// func test(v *Variants, a, b int) bool {

// 	var indA, indB int

// 	for ind, zn := range v.slA {

// 		if zn == a {
// 			indA = ind
// 		}
// 		if zn == b {
// 			indB = ind
// 		}
// 	}

// 	for ind, zn := range v.slB {

// 		if zn == a {
// 			indA = ind
// 		}
// 		if zn == b {
// 			indB = ind
// 		}
// 	}

// 	for ind, zn := range v.slC {

// 		if zn == a {
// 			indA = ind
// 		}
// 		if zn == b {
// 			indB = ind
// 		}
// 	}

// 	for ind, zn := range v.slD {

// 		if zn == a {
// 			indA = ind
// 		}
// 		if zn == b {
// 			indB = ind
// 		}
// 	}

// 	for ind, zn := range v.slE {

// 		if zn == a {
// 			indA = ind
// 		}
// 		if zn == b {
// 			indB = ind
// 		}
// 	}

// 	if indA == indB {
// 		return true //, indA, indB
// 	}

// 	return false
// }

// func testInput(slA, slB, slC, slD, slE []int, a, b int) bool {

// 	var indA, indB int //, counter int //индексы со сдвигом двух чисел.
// 	varMap := make(map[string]int)

// 	//сюда добавляем мапу.

// 	for ind, zn := range slA {

// 		if zn == a {
// 			indA = ind
// 			varMap["slA"] = ind
// 		}
// 		if zn == b {
// 			indB = ind
// 			varMap["slA"] = ind
// 		}
// 	}

// 	for ind, zn := range slB {

// 		if zn == a {
// 			indA = ind
// 			varMap["slB"] = ind
// 		}
// 		if zn == b {
// 			indB = ind
// 			varMap["slB"] = ind
// 		}
// 	}

// 	for ind, zn := range slC {

// 		if zn == a {
// 			indA = ind
// 			varMap["slC"] = ind
// 		}
// 		if zn == b {
// 			indB = ind
// 			varMap["slC"] = ind
// 		}
// 	}

// 	for ind, zn := range slD {

// 		if zn == a {
// 			indA = ind
// 			varMap["slD"] = ind
// 		}
// 		if zn == b {
// 			indB = ind
// 			varMap["slD"] = ind
// 		}
// 	}

// 	for ind, zn := range slE {

// 		if zn == a {
// 			indA = ind
// 			varMap["slE"] = ind
// 		}
// 		if zn == b {
// 			indB = ind
// 			varMap["slE"] = ind
// 		}
// 	}

// 	if indA == indB {
// 		return true //, indA, indB
// 	}

// 	return false

// }
