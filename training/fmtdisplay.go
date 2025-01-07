// В приведенном коде игры текущее состояние сетки выводится на экран следующим образом: fmt.Println(currentWorld), что не совсем наглядно. Чтобы получить наглядное изображение нашей сетки мы можем отрисовывать клетки разного цвета, например:
// brownSquare := "\xF0\x9F\x9F\xAB"
// greenSquare := "\xF0\x9F\x9F\xA9"
// Для этого, созданный нами тип type World struct {
// Height int // высота сетки
// Width int // ширина сетки
// Cells [][]bool
// }
// должен иметь метод String(), который будет формировать строку для отображения. Напишите данный метод и попробуйте запустить всю программу, поэкспериментируйте с разными символами и цветами.

// Примечания
// Код программы должен содержать описание струкрутры World:
// type World struct { Height int Width int Cells [][]bool }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // World представляет сетку клеток
// type World struct {
// 	Height int
// 	Width  int
// 	Cells  [][]bool
// }

// // String формирует строковое представление мира для отображения
// func (w *World) String() string {
// 	brownSquare := "\xF0\x9F\x9F\xAB" // Символ для живых клеток
// 	greenSquare := "\xF0\x9F\x9F\xA9" // Символ для мёртвых клеток

// 	result := ""
// 	for _, row := range w.Cells {
// 		for _, cell := range row {
// 			if cell {
// 				result += brownSquare + " "
// 			} else {
// 				result += greenSquare + " "
// 			}
// 		}
// 		result += "\n"
// 	}
// 	return result
// }

// // NewWorld создаёт новый мир заданных размеров
// func NewWorld(height, width int) *World {
// 	cells := make([][]bool, height)
// 	for i := range cells {
// 		cells[i] = make([]bool, width)
// 	}
// 	return &World{
// 		Height: height,
// 		Width:  width,
// 		Cells:  cells,
// 	}
// }

// // Neighbors возвращает количество живых соседей клетки
// func (w *World) Neighbors(x, y int) int {
// 	directions := [][2]int{
// 		{-1, -1}, {-1, 0}, {-1, 1},
// 		{0, -1}, {0, 1},
// 		{1, -1}, {1, 0}, {1, 1},
// 	}

// 	count := 0
// 	for _, dir := range directions {
// 		nx, ny := x+dir[0], y+dir[1]
// 		if nx >= 0 && nx < w.Height && ny >= 0 && ny < w.Width {
// 			if w.Cells[nx][ny] {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

// // Next возвращает следующее состояние клетки
// func (w *World) Next(x, y int) bool {
// 	n := w.Neighbors(x, y)
// 	alive := w.Cells[x][y]
// 	if (n == 2 || n == 3) && alive {
// 		return true
// 	}
// 	if n == 3 && !alive {
// 		return true
// 	}
// 	return false
// }

// // NextState обновляет состояние мира
// func NextState(oldWorld, newWorld *World) {
// 	for i := 0; i < oldWorld.Height; i++ {
// 		for j := 0; j < oldWorld.Width; j++ {
// 			newWorld.Cells[i][j] = oldWorld.Next(i, j)
// 		}
// 	}
// }

// // Seed случайно заполняет клетки мира
// func (w *World) Seed(rng *rand.Rand) {
// 	for _, row := range w.Cells {
// 		for i := range row {
// 			if rng.Intn(10) == 1 { // С вероятностью 10% клетка будет живой
// 				row[i] = true
// 			}
// 		}
// 	}
// }

// func main() {
// 	// Создание генератора случайных чисел
// 	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

// 	// Создание мира
// 	height := 10
// 	width := 10
// 	currentWorld := NewWorld(height, width)
// 	nextWorld := NewWorld(height, width)

// 	// Инициализация начального состояния
// 	currentWorld.Seed(rng)

// 	// Анимация игры
// 	for {
// 		// Вывод текущего состояния мира
// 		fmt.Print(currentWorld.String())

// 		// Вычисление следующего состояния
// 		NextState(currentWorld, nextWorld)

// 		// Смена миров
// 		currentWorld, nextWorld = nextWorld, currentWorld

// 		// Пауза
// 		time.Sleep(500 * time.Millisecond)

// 		// Очистка экрана
// 		fmt.Print("\033[H\033[2J")
// 	}
// }

package main

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func (w *World) String() string {
	brownSquare := "\xF0\x9F\x9F\xAB" // Символ для живых клеток
	greenSquare := "\xF0\x9F\x9F\xA9" // Символ для мёртвых клеток

	result := ""
	for _, row := range w.Cells {
		for _, cell := range row {
			if cell {
				result += brownSquare + " "
			} else {
				result += greenSquare + " "
			}
		}
		result += "\n"
	}
	return result
}

func NewWorld(height, width int) *World {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
}

func (w *World) Neighbors(x, y int) int {
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < w.Height && ny >= 0 && ny < w.Width {
			if w.Cells[nx][ny] {
				count++
			}
		}
	}
	return count
}
