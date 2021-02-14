package main

import (
	"fmt"
)

//CheckMove2 is for check move2
type CheckMove2 struct {
	Xindex int64
	Yindex int64
	Move   []int64
}

//Userchoise is ...
type Userchoise struct {
	Cname  int64
	Cnum   int64
	Cindex int
}

//Choises is ...
type Choises []Userchoise

//Vary is ...
type Vary struct {
	Name int64
	Num  []int64
}

//Coordinats is ...
type Coordinats []Vary

func main() {

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

	var start1, start2 int64 //исходим из того, что будут заданы начальные координаты для двух из четырех клеток первого стартового хода коня.
	var occ []int64

	p := new() //случай для доски 5х5. Пока указываем явно, потом переделаем.

	num1 := []int64{1, 2, 3, 4, 5}
	num2 := []int64{6, 7, 8, 9, 10}
	num3 := []int64{11, 12, 13, 14, 15}
	num4 := []int64{16, 17, 18, 19, 20}
	num5 := []int64{21, 22, 23, 24, 25}

	p.add(0, num1)
	p.add(1, num2)
	p.add(2, num3)
	p.add(3, num4)
	p.add(4, num5)

	start1 = 2 //для теста
	start2 = 3 //нач.координаты

	testPosition, x, y := addChoice(p, start1, start2) //вся эта проверка - на соседство выбранных клеток. Также определяем местоположение на шахматной доске.
	if !(testPosition) {
		fmt.Println("Введены не верные начальные координаты!") //проверка на соседство клеточек не пройдена
		return
	}

	//occ.addOcc(x.Cnum,y.Cnum) //передаем координаты выбора в занятые клетки
	//occupiedMoves(x, y)

	makeMove(p, occ, x, y)

}

func makeMove(p *Coordinats, occ []int64, x Userchoise, y Userchoise) {

	var mNew CheckMove2

	maxNumber, minNumber := findMaxMin(x.Cname, y.Cname)               //определяем допустимые варианты по номерам
	maxIndex, minIndex := findMaxMin(int64(x.Cindex), int64(y.Cindex)) //определяем допустимые варианты по колонкам

	for i := minNumber; i <= (maxNumber); i++ {
		pp := (*p)[i]

		for ind, zn := range pp.Num {

			if zn == x.Cnum || zn == y.Cnum { //берём только близкие колонки и не выходим за пределы
				continue
			}

			if ind < minIndex || ind > maxIndex { //берём только близкие по индексу.
				continue
			}

			occupied := testAlreadyOccupied(zn, occ) //проверка не являются ли уже занятыми координаты
			if occupied {
				continue
			}

			// if len(move1) == 0 {
			// 	move1[1] = make(map[int64]int) //инициализируем внутреннюю мапу
			// 	move1[1][zn] = ind
			// 	continue
			// }

			// move2[2] = make(map[int64]int)
			// move2[2][zn] = ind

			mNew = CheckMove2{
				Xindex: int64(x.Cindex),
				Yindex: int64(y.Cindex),
				M1Map:  move1,
				M2Map:  move2,
			}

			testMove2(mNew) //проверяем, подходит ли нам вторая координата (стоит ли она рядом, можем ли мы походить буквой "г")
			//тут проверяем, одинаковые ли индексы у первых записей?

			fmt.Println(move1, move2)
			break
		}

		if len(move1) != 0 && len(move2) != 0 {
			break //вызываем рекурсивно функцию
		}

	}

}

func testAlreadyOccupied(zn int64, occ []int64) bool { //как сравнить? нужен ли обязательно слайс структур?

	for _, i := range occ {
		if i == zn {
			return true
		}
	}
	return false
}

func testMove2(m *CheckMove2) {

	if m.Xindex == m.Yindex { //в этом случае индексы выбора не должны быть одинаковы

	}

}

//ограничиваем поиск свободных клеток только в пределах близких индексов, номеров
func findMaxMin(x, y int64) (int, int) {

	var max, min int

	if x > y {
		max = int(x)
		min = int(y)
	} else {
		max = int(y)
		min = int(x)
	}

	max++
	if min > 0 {
		min--
	}

	return max, min
}

func (p *Coordinats) add(name int64, num []int64) {

	pnew := Vary{
		Name: name,
		Num:  num,
	}
	*p = append(*p, pnew)
}

//добавляем к выбору: название, номер строки и номер по порядку-индекс.
func (c *Choises) addChoice(cname int64, cnum int64, cindex int) {
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
func addChoice(p *Coordinats, x1 int64, x2 int64) (bool, Userchoise, Userchoise) {

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
			return true, x, y
		}

	} else {
		if x.Cindex == y.Cindex {
			return true, x, y
		}
	}
	return false, x, y
}

//проверяем, что ввели именно число
// func getNumInt() (int, error) {

// 	input := bufio.NewScanner(os.Stdin)
// 	input.Scan()

// 	num, err := strconv.Atoi(input.Text())
// 	return num, err

// }
