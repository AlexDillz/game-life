// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// type World struct {
// 	Height int
// 	Width  int
// 	Cells  [][]bool
// }

// // Функция для создания нового мира
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

// // Метод для подсчета живых соседей
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

// // Метод для вычисления следующего состояния клетки
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

// // Функция для вычисления следующего состояния мира
// func NextState(oldWorld, newWorld *World) {
// 	for i := 0; i < oldWorld.Height; i++ {
// 		for j := 0; j < oldWorld.Width; j++ {
// 			newWorld.Cells[i][j] = oldWorld.Next(i, j)
// 		}
// 	}
// }

// // Метод для случайной инициализации клеток
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
// 	// Создание локального генератора случайных чисел
// 	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

// 	height := 10
// 	width := 10
// 	currentWorld := NewWorld(height, width)
// 	nextWorld := NewWorld(height, width)

// 	// Устанавливаем начальное состояние с использованием локального генератора
// 	currentWorld.Seed(rng)

// 	for {
// 		// Выводим текущее состояние мира
// 		for _, row := range currentWorld.Cells {
// 			for _, cell := range row {
// 				if cell {
// 					fmt.Print("O ") // Живая клетка
// 				} else {
// 					fmt.Print(". ") // Мёртвая клетка
// 				}
// 			}
// 			fmt.Println()
// 		}
// 		fmt.Println()

// 		// Вычисляем следующее состояние
// 		NextState(currentWorld, nextWorld)

// 		// Обновляем текущее состояние
// 		*currentWorld = *nextWorld

// 		// Пауза
// 		time.Sleep(100 * time.Millisecond)

// 		// Очистка экрана для анимации
// 		fmt.Print("\033[H\033[2J")
// 	}
// }
