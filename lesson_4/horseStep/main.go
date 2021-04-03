package main

import "fmt"

type HorseStep struct {
	xchoice Coordinata
	ychoice Coordinata
	xresult Coordinata
	yresult Coordinata
	index   int
}

type Coordinata struct {
	cRow int //ряд
	cNum int //номер
	cInd int //индекс
}

type Field []Coordinata

type Rows []Field

type BorderForStep []Coordinata // границы для определения хода.

type Alldata struct {
	Variants Rows          //поле
	borders  BorderForStep //возможные границы для хода
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

	//стартовые границы поля
	sizeField := 5
	end := sizeField * sizeField
	num := 0

	var fieldSlice []Coordinata
	var getCoord Coordinata
	rowsSlice := make([]Field, sizeField)

	xNum := 7 //стартовые точки отсчета
	yNum := 8

	//заполняем поле 5*5
	for row := 0; row < sizeField; row++ { //попробовать переписать через range и перекинуть в отдельные функции.

		for ind := 0; ind < sizeField; ind++ {

			num++
			getCoord = fillCoordinata(row, num, ind)

			fieldSlice = append(fieldSlice, getCoord) //из-за необходимости очистки, используем в данном случае append

			if num%sizeField == 0 {
				break
			}
		}

		rowsSlice[row] = fieldSlice //более быстрый способ заполнения

		fieldSlice = nil //чтобы не задублировать, очищаем слайс.

		if num == end {
			break
		}
	}

	var gotStruct Alldata
	gotStruct.Variants = rowsSlice

	testPosition, xchoice, ychoice := gotStruct.checkChoice(xNum, yNum)
	if !(testPosition) {
		fmt.Println("Введены не верные начальные координаты!") //проверка на соседство клеточек не пройдена
		return
	}

	hs := HorseStep{
		xchoice: xchoice,
		ychoice: ychoice,
		index:   0,
	}

	gotStruct.makeStep(&hs)

}

func (a *Alldata) makeStep(h *HorseStep) {

	var canMove bool

	a.findBordersForStep(h)

	for i := 0; i < len(a.borders); i++ {

		if h.xresult.cInd == 0 && h.xresult.cNum == 0 && h.xresult.cRow == 0 {

			h.xresult = fillCoordinata(a.borders[h.index].cRow, a.borders[h.index].cNum, a.borders[h.index].cInd)
			continue
		}

		if h.xresult.cNum == a.borders[i].cNum {
			continue
		}

		h.yresult = fillCoordinata(a.borders[i].cRow, a.borders[i].cNum, a.borders[i].cInd)
		canMove = checkStep(h)

		if canMove { //нашли устраивающий вариант.
			fmt.Println("step: ", h.xchoice.cNum, h.ychoice.cNum, h.xresult.cNum, h.yresult.cNum) //ход.
			a.prepareStep(h)
			break
		}
	}

	if !(canMove) {

		h.index++

		if h.index < len(a.borders) {
			h.xresult = fillCoordinata(a.borders[h.index].cRow, a.borders[h.index].cNum, a.borders[h.index].cInd) //сдвигаем координату на один.
			a.makeStep(h)
		} else { //с этим координатами нет возможности хода.

			if len(a.borders) != 0 {
				h.index = 0
				h.xresult.cInd = 0
				h.xresult.cNum = 0
				h.xresult.cRow = 0
				a.prepareStep(h)

			}

		}

	}

}

func (a *Alldata) prepareStep(h *HorseStep) {

	a.delCoordiantsFromAllData(h) //удаляем координаты из доступных.
	a.findNewXYChoice(h)          //определяем новые xchoice ychoice
	a.makeStep(h)                 //снова делаем ход.

}

func (a *Alldata) delCoordiantsFromAllData(h *HorseStep) {

	//удаляем выбранные координаты из слайса.
	for v := 0; v < len(a.Variants); v++ {

		thisRow := a.Variants[v]

		for ind := 0; ind < len(thisRow); ind++ {

			if thisRow[ind].cNum == h.xchoice.cNum || thisRow[ind].cNum == h.ychoice.cNum || thisRow[ind].cNum == h.xresult.cNum || thisRow[ind].cNum == h.yresult.cNum {
				thisRow[ind] = thisRow[len(thisRow)-1]
				thisRow = thisRow[:len(thisRow)-1]
				a.Variants[v] = thisRow //удаляем и получаем неотсортированный слайс. Нужно ли его сортировать?
				ind--
			}

		}

		//for ind, zn := range thisRow {

		//if zn.cNum == h.xchoice.cNum || zn.cNum == h.ychoice.cNum || zn.cNum == h.xresult.cNum || zn.cNum == h.yresult.cNum {
		//	thisRow[ind] = thisRow[len(thisRow)-1]

		//	thisRow = thisRow[:len(thisRow)-1]
		//	a.Variants[v] = thisRow //удаляем и получаем неотсортированный слайс. Нужно ли его сортировать?

		//}

	}

}

func (a *Alldata) findNewXYChoice(h *HorseStep) {

	var testPosition bool
	var xchoice, ychoice Coordinata

	x := make(map[string]int) //чтобы удобнее было осуществлять проверку на пустоту.
	y := make(map[string]int)

	//удаляем выбранные координаты из слайса.
	for v := 0; v < len(a.Variants); v++ {

		thisRow := a.Variants[v]
		for _, zn := range thisRow {

			if _, ok := x["num"]; !ok {
				x["num"] = zn.cNum
				continue
			}

			y["num"] = zn.cNum

			testPosition, xchoice, ychoice = a.checkChoice(x["num"], y["num"])
			if testPosition {
				break
			}
		}

		if testPosition {
			break
		}

	}

	h.xchoice = xchoice
	h.ychoice = ychoice

	//обнуляем
	h.xresult.cInd = 0
	h.xresult.cNum = 0
	h.xresult.cRow = 0

	h.yresult.cInd = 0
	h.yresult.cNum = 0
	h.yresult.cRow = 0
}

func checkStep(h *HorseStep) bool { //переделать на более понятный вариант - всю функцию.

	if h.xchoice.cInd == h.ychoice.cInd { //index - одинаковый

		if h.xresult.cInd == h.yresult.cInd {
			return false
		}

		if (h.xresult.cInd-h.yresult.cInd) > 1 || (h.yresult.cInd-h.xresult.cInd) > 1 {
			return false
		}

		if h.xresult.cRow != h.yresult.cRow {
			return false
		}

		if (h.xresult.cInd + h.yresult.cInd + 1) < h.xchoice.cInd {
			return false
		}

		if h.xresult.cInd != h.xchoice.cInd {

			//flag = false

			if h.yresult.cInd != h.xchoice.cInd {
				return false
			}
		}

	}

	if h.xchoice.cInd != h.ychoice.cInd { //row - одинаковый

		if h.xresult.cInd != h.yresult.cInd {
			return false
		}

		if (h.xresult.cRow-h.yresult.cRow) > 1 || (h.yresult.cRow-h.xresult.cRow) > 1 {
			return false
		}

		if (h.xchoice.cRow != h.xresult.cRow) && (h.xchoice.cRow != h.yresult.cRow) {

			if (h.xresult.cInd != h.xchoice.cInd) && (h.xresult.cInd != h.ychoice.cInd) {
				return false
			}
		}

	}

	return true
}

func (a *Alldata) findBordersForStep(h *HorseStep) {

	var borderSlice BorderForStep

	maxRow, minRow := findMaxMinRow(h.xchoice.cRow, h.ychoice.cRow, len(a.Variants)-1)       //определяем допустимые варианты по номерам
	maxIndex, minIndex := findMaxMinIndex(h.xchoice.cInd, h.ychoice.cInd, len(a.Variants)-1) //определяем допустимые варианты по колонкам

	for row := minRow; row <= (maxRow); row++ {
		thisRow := (a.Variants)[row]

		for _, num := range thisRow {

			if num.cNum == h.xchoice.cNum || num.cNum == h.ychoice.cNum { //2 следующие координаты не должны быть равны двум выбранным (от которых отталкиваемся)
				continue
			}

			if num.cInd < minIndex || num.cInd > maxIndex { //берём только близкие по индексу.
				continue
			}

			result := fillCoordinata(row, num.cNum, num.cInd) //заполняем координату
			borderSlice = append(borderSlice, result)         //можно переделать не через append

		}

	}

	a.borders = borderSlice
}

//ограничиваем поиск свободных клеток только в пределах близких индексов, номеров
func findMaxMinIndex(x, y, lenField int) (int, int) {
	return getBoundaryMaxMin(x, y, lenField)
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
		max, min = getBoundaryMaxMin(x, y, lenField)
	}

	return max, min
}

func getBoundaryMaxMin(x, y, lenField int) (int, int) {

	var max, min int

	if x > y {
		max = x
		min = y
	} else if x < y {
		max = y
		min = x
	} else {
		max = x
		return max, min - 2
	}

	if max < lenField {
		max++
	}

	if min != 0 {
		min--
	}

	return max, min
}

func fillCoordinata(row, num, ind int) Coordinata {

	cresult := Coordinata{ //заполняем пустую структуру
		cRow: row,
		cNum: num,
		cInd: ind,
	}

	return cresult
}

//сравниваем выбор пользователя с тем, что есть в таблице с номерами. И проверяем на соседство выбранных координат
func (a *Alldata) checkChoice(x1 int, x2 int) (bool, Coordinata, Coordinata) {

	var xchoice, ychoice Coordinata

	for _, thisRow := range a.Variants {

		for _, num := range thisRow {

			if num.cNum == x1 {
				xchoice = num
				continue
			}

			if num.cNum == x2 {
				ychoice = num
			}

		}
	}

	if xchoice.cRow == ychoice.cRow { //проверка на соседство двух клеточек шахматной доски. Если клетки находятся поодаль - (А1, B10) - то это ошибка и запускать программу нет смысла.
		if (xchoice.cNum-ychoice.cNum) == 1 || (ychoice.cNum-xchoice.cNum) == 1 {
			return true, xchoice, ychoice
		}

	} else {
		if xchoice.cInd == ychoice.cInd {
			return true, xchoice, ychoice
		}
	}

	return false, xchoice, ychoice
}

//проверяем, что ввели именно число
// func getNumInt() (int, error) {

// 	input := bufio.NewScanner(os.Stdin)
// 	input.Scan()

// 	num, err := strconv.Atoi(input.Text())
// 	return num, err

// }
