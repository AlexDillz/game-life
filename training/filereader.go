// // Напишите новый метод func (w *World) LoadState(filename string) error для структуры World,
//  который будет считывать из файла исходное состояние и устанавливать размерность сетки в соответствии с данным файлом. Файл должен содержать бинарные данные, например:
// // 110011
// // 100101
// // где 1 - живая клетка, 0 - мертвая. Размерность после чтения данного файла: width: 6, height: 2.
// В случае нарушения размерности в файле (разное количество элементов в строках) должна возвращаться ошибка.

// // Примечания
// // Код программы должен содержать описание струкрутры World:
// // type World struct { Height int Width int Cells [][]bool }

package main

import (
	"fmt"
	"os"
	"strings"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func (w *World) LoadState(filename string) error {
	// Чтение данных из файла
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Разделение данных на строки
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	// Проверка размерности строк
	width := len(lines[0])
	for _, line := range lines {
		if len(line) != width {
			return fmt.Errorf("inconsistent row lengths in file")
		}
	}

	// Инициализация сетки с новой размерностью
	height := len(lines)
	cells := make([][]bool, height)
	for i, line := range lines {
		cells[i] = make([]bool, width)
		for j, char := range line {
			if char == '1' {
				cells[i][j] = true
			} else if char == '0' {
				cells[i][j] = false
			} else {
				return fmt.Errorf("invalid character '%c' in file", char)
			}
		}
	}

	// Установка новых значений для структуры World
	w.Height = height
	w.Width = width
	w.Cells = cells
	return nil
}
