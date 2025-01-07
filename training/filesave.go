// Напишите новый метод func (w *World) SaveState(filename string) error для сохранения текущего состояния сетки в файл.
//Метод должен создавать новый файл и записать данные в бинарном виде, например:
// 110011
// 100101

// Примечания
// Код программы должен содержать описание струкрутры World:
// type World struct { Height int Width int Cells [][]bool }

package main

import (
	"os"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func (w *World) SaveState(filename string) error {
	// Создаём или перезаписываем файл
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Построчно формируем представление сетки
	for i, row := range w.Cells {
		for _, cell := range row {
			if cell {
				file.WriteString("1") // Живая клетка
			} else {
				file.WriteString("0") // Мёртвая клетка
			}
		}
		// Добавляем перенос строки только если это не последняя строка
		if i < len(w.Cells)-1 {
			file.WriteString("\n")
		}
	}
	return nil
}
