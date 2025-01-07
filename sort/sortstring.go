// Напишите программу, содержащую функцию `SortNames(names []string)``,
// которая сортирует список имён в алфавитном порядке.
// Если первые символы совпадают, сортировать по вторым, и т.д.

// Примечания
// Пример отсортированного списка: Аксинья, Арина, Варвара, Есения

package main

import (
	"fmt"
	"slices"
)

func SortNames(names []string) {
	slices.SortFunc(names, func(a, b string) int {
		if a < b {
			return -1 // Сортировка в алфавитном порядке
		}
		if a > b {
			return 1
		}
		return 0
	})
}

func main() {
	names := []string{"Ангелина", "Игорь", "Алексей", "Йок", "Ян", "Яна"}
	SortNames(names) // Сортируем список
	fmt.Println(names)
}
