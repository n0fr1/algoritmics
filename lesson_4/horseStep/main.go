package main

import "fmt"

//AllHorsesSteps is...
type AllHorsesSteps []HorseStep

//HorseStep is...
type HorseStep struct { //можно для уже выполенных ходов именно такую структуру и создать.
	c1 Coordinata
	c2 Coordinata
	c3 Coordinata
	c4 Coordinata
}

//Coordinata is ...
type Coordinata struct {
	cRow map[string]int //ряд
	cNum map[string]int //номер
	cInd map[string]int //индекс
}

//Userchoise is ...как и чем можно заменить Userchoise?
type Userchoise struct {
	uRow int
	uNum int
	uInd int
}

//Choises is ...
type Choises []Userchoise

//VarPoint is ...
type VarPoint struct {
	Row     int
	NumsRow []int
}

//VaryablesField is ...
type VaryablesField []VarPoint

//BorderForStep is...
type BorderForStep []Coordinata //сюда закинем границы для определения хода.

//TestStep is ...
type TestStep struct { //тестируем возможность хода
	VaryablesField []VarPoint
	xchoise        Userchoise
	ychoise        Userchoise
	result1        Coordinata
	result2        Coordinata
	borders        BorderForStep
}

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

	var varyables VaryablesField
	var start1, start2 int //исходим из того, что будут заданы начальные координаты для двух из четырех клеток первого стартового хода коня.
	//var occ []int

	//случай для доски 5х5. Пока указываем явно, потом переделаем.
	varyables.addPoint(0, []int{1, 2, 3, 4, 5})
	varyables.addPoint(1, []int{6, 7, 8, 9, 10})
	varyables.addPoint(2, []int{11, 12, 13, 14, 15})
	varyables.addPoint(3, []int{16, 17, 18, 19, 20})
	varyables.addPoint(4, []int{21, 22, 23, 24, 25})

	start1 = 13 //для теста
	start2 = 14 //нач.координаты

	testPosition, x, y := checkChoice(&varyables, start1, start2) //вся эта проверка - на соседство выбранных клеток. Также определяем местоположение на шахматной доске.
	if !(testPosition) {
		fmt.Println("Введены не верные начальные координаты!") //проверка на соседство клеточек не пройдена
		return
	}

	mstep := TestStep{
		VaryablesField: varyables,
		xchoise:        x,
		ychoise:        y,
	}

	var borders BorderForStep
	makeStep(&mstep, borders)

	//occ.addOcc(x.Cnum,y.Cnum) //передаем координаты выбора в занятые клетки
	//occupiedMoves(x, y)

}

func makeStep(m *TestStep, borders BorderForStep) {

	var mayMove bool

	if len(borders) == 0 {
		borders = findBordersForStep(m)
	}

	for i := 0; i <= len(borders)-1; i++ {

		if _, ok := m.result1.cNum["Num"]; !ok { //проверяем структуру на пустоту, если пустая, то заполняем.
			m.result1 = fillCoordinata(borders[i].cRow["Row"], borders[i].cNum["Num"], borders[i].cInd["Ind"]) //это будет первая координата.
			continue
		}

		m.result2 = fillCoordinata(borders[i].cRow["Row"], borders[i].cNum["Num"], borders[i].cInd["Ind"]) //заполняем вторую предположительную координату хода.

		mayMove = checkStep(m.result1, m.result2, m.xchoise, m.ychoise) //здесь есть возможность - проба загнать все в одну структуру
		if mayMove {                                                    //если результат - верный, то прерываем.
			fmt.Println(mayMove, m.result1.cNum["Num"], m.result2.cNum["Num"], m.xchoise.uNum, m.ychoise.uNum)
			break
		}
	}

	if !mayMove {

		borders = borders[1 : len(borders)-1] //уменьшаем слайс
		m.result1.cRow = make(map[string]int) //очищаем мап.
		m.result1.cNum = make(map[string]int)
		m.result1.cInd = make(map[string]int)

		makeStep(m, borders)
		return
	}
}

//получаем предварительные границы. Чтобы понимать возможные границы хода.
func findBordersForStep(m *TestStep) BorderForStep {
	//func findBordersForStep(v *VaryablesField, x Userchoise, y Userchoise) BorderForStep {

	//var result1, result2 Coordinata //2 оставшиеся координаты предполагаемого хода.

	var borderSlice BorderForStep

	maxRow, minRow := findMaxMinRow(m.xchoise.uRow, m.ychoise.uRow, len(m.VaryablesField)-1) //определяем допустимые варианты по номерам
	maxIndex, minIndex := findMaxMinIndex(m.xchoise.uInd, m.ychoise.uInd)                    //определяем допустимые варианты по колонкам

	for row := minRow; row <= (maxRow); row++ {
		thisRow := (m.VaryablesField)[row]

		for ind, num := range thisRow.NumsRow {

			if num == m.xchoise.uNum || num == m.ychoise.uNum { //2 следующие координаты не должны быть равны двум выбранным (от которых отталкиваемся)
				continue
			}

			if ind < minIndex || ind > maxIndex { //берём только близкие по индексу.
				continue
			}

			result := fillCoordinata(row, num, ind) //заполняем координату
			borderSlice = append(borderSlice, result)

		}

	}

	return borderSlice

}

func fillCoordinata(row, num, ind int) Coordinata {

	cRow := make(map[string]int) //создаем и заполняем мап'у.
	cNum := make(map[string]int)
	cInd := make(map[string]int)

	cRow["Row"] = row //сюда нужно передавать обязательно ряд
	cNum["Num"] = num
	cInd["Ind"] = ind

	cresult := Coordinata{ //Заполняем структуру
		cRow: cRow,
		cNum: cNum,
		cInd: cInd,
	}

	return cresult
}

func checkStep(result1, result2 Coordinata, x, y Userchoise) bool {

	if x.uInd == y.uInd { //index - одинаковый

		if result1.cInd["Ind"] == result2.cInd["Ind"] {
			return false
		}

		if (result1.cInd["Ind"]-result2.cInd["Ind"]) > 1 || (result2.cInd["Ind"]-result1.cInd["Ind"]) > 1 {
			return false
		}

	}

	if x.uInd != y.uInd { //row - одинаковый

		if result1.cInd["Ind"] != result2.cInd["Ind"] {
			return false
		}

		if (result1.cRow["Row"]-result2.cRow["Row"]) > 1 || (result2.cRow["Row"]-result1.cRow["Row"]) > 1 {
			return false
		}

		if result1.cRow["Row"] != x.uRow && result1.cRow["Row"] != y.uRow { //расположены не на одном ряду, что и выбранные координаты.
			if result1.cInd["Ind"] != x.uInd && result1.cInd["Ind"] != y.uInd && result2.cInd["Ind"] != x.uInd && result2.cInd["Ind"] != y.uInd { //нет пересечений по индексу.
				return false
			}

		}
	}

	return true
}

func testAlreadyOccupied(num int, occ []int) bool { //как сравнить? нужен ли обязательно слайс структур?

	for _, i := range occ {
		if i == num {
			return true
		}
	}
	return false
}

//ограничиваем поиск свободных клеток только в пределах близких индексов, номеров
func findMaxMinIndex(x, y int) (int, int) {
	return getBoundaryMaxMin(x, y)
}

func findMaxMinRow(x, y, lenField int) (int, int) {

	var max, min int

	if x == y {

		if (x + 2) > lenField {
			max = lenField
		} else {
			max = x + 2
		}

		if (x - 2) < 0 {
			min = 0
		} else {
			min = x - 2
		}

	} else {
		max, min = getBoundaryMaxMin(x, y)
	}

	return max, min
}

func getBoundaryMaxMin(x, y int) (int, int) {

	var max, min int

	if x > y {
		max = x
		min = y
	} else {
		max = y
		min = x
	}

	max++
	if min != 0 {
		min--
	}

	return max, min

}

func (v *VaryablesField) addPoint(row int, num []int) {

	newPoint := VarPoint{
		Row:     row,
		NumsRow: num,
	}
	*v = append(*v, newPoint)
}

//добавляем к выбору: название, номер строки и номер по порядку-индекс.
func (ch *Choises) addChoice(row int, num int, index int) {

	chnew := Userchoise{
		uRow: row,
		uNum: num,
		uInd: index,
	}

	*ch = append(*ch, chnew)
}

//сравниваем выбор пользователя с тем, что есть в таблице с номерами.
func checkChoice(v *VaryablesField, x1 int, x2 int) (bool, Userchoise, Userchoise) {

	var chSlice Choises

	for _, thisRow := range *v {
		for ind, num := range thisRow.NumsRow {
			if num == x1 || num == x2 {
				chSlice.addChoice(thisRow.Row, num, ind)
			}
		}
	}

	x := chSlice[0] //Поскольку координаты - две. Поэтому обращаемся напрямую по индексу
	y := chSlice[1]

	if x.uRow == y.uRow { //проверка на соседство двух клеточек шахматной доски. Если клетки находятся поодаль - (А1, B10) - то это ошибка и запускать программу нет смысла.
		if (x.uNum-y.uNum) == 1 || (y.uNum-x.uNum) == 1 {
			return true, x, y
		}

	} else {
		if x.uInd == y.uInd {
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
