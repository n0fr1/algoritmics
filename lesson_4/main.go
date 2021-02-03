//Беляев Е.Г
package main

import (
	"fmt"
)

//Userchoise is ...
type Userchoise struct {
	Cname  string
	Cnum   int64
	Cindex int
}

//Choises is ...
type Choises []Userchoise

//Vary is ...
type Vary struct {
	Name string
	Num  []int64
}

//Coordinats is ...
type Coordinats []Vary

func main() {

	var x1, x2 int64 //исходим из того, что будут заданы начальные координаты для двух из четырех клеток первого стартового хода коня.

	// fmt.Println("Укажите стартовые клетки для положения коня")
	// fmt.Println("Введите первый номер: ")
	// x1, err := getNumInt() //проверяем на число
	// if err != nil {
	// 	fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
	// 	os.Exit(1)
	// }

	// fmt.Println("Введите второй номер: ")
	// x2, err := getNumInt() //проверяем на число
	// if err != nil {
	// 	fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
	// 	os.Exit(1)
	// }

	p := new() //случай для доски 5х5. Пока указываем явно, потом переделаем.

	num1 := []int64{1, 2, 3, 4, 5}
	num2 := []int64{6, 7, 8, 9, 10}
	num3 := []int64{11, 12, 13, 14, 15}
	num4 := []int64{16, 17, 18, 19, 20}
	num5 := []int64{21, 22, 23, 24, 25}

	p.add("A", num1)
	p.add("B", num2)
	p.add("C", num3)
	p.add("D", num4)
	p.add("E", num5)

	x1 = 3 //для теста
	x2 = 2

	testPosition, c := addChoice(p, x1, x2) //вся эта проверка - на соседство выбранных клеток. Также определяем местоположение на шахматной доске.
	if !(testPosition) {
		fmt.Println("Введены не верные начальные координаты!") //проверка на соседство клеточек не пройдена
		return
	}
	fmt.Println(c)
}

func (p *Coordinats) add(name string, num []int64) {

	pnew := Vary{
		Name: name,
		Num:  num,
	}
	*p = append(*p, pnew)
}

//добавляем к выбору: название, номер строки и номер по порядку-индекс.
func (c *Choises) addChoice(cname string, cnum int64, cindex int) {
	cnew := Userchoise{
		Cname:  cname,
		Cnum:   cnum,
		Cindex: cindex,
	}
	*c = append(*c, cnew)
}

func new() *Coordinats {
	var arr Coordinats
	return &arr
}

func newChoice() *Choises {
	var arr Choises
	return &arr
}

//сравниваем выбор пользователя с тем, что есть в таблице с номерами.
func addChoice(p *Coordinats, x1 int64, x2 int64) (bool, *Choises) {

	c := newChoice()

	for _, pp := range *p {
		for ind, zn := range pp.Num {
			if zn == x1 || zn == x2 {
				c.addChoice(pp.Name, zn, ind)
			}
		}
	}

	x := (*c)[0] //Поскольку координаты - две. Поэтому обращаемся напрямую по индексу
	y := (*c)[1]

	if x.Cname == y.Cname { //проверка на соседство двух клеточек шахматной доски. Если клетки находятся поодаль - (А1, B10) - то это ошибка и запускать программу нет смысла.
		if (x.Cnum-y.Cnum) == 1 || (y.Cnum-x.Cnum) == 1 {
			return true, c
		}

	} else {
		if x.Cindex == y.Cindex {
			return true, c
		}
	}
	return false, c
}

//проверяем, что ввели именно число
// func getNumInt() (int, error) {

// 	input := bufio.NewScanner(os.Stdin)
// 	input.Scan()

// 	num, err := strconv.Atoi(input.Text())
// 	return num, err

// }
