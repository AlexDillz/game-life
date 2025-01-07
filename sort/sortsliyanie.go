// Даны два слайса. Напишите программу, содержащую функцию SortAndMerge(left, right []int) []int,
// которая объединит слайсы в один отсортированный в два этапа: - отсортировать каждый слайс - объединить полученные слайсы в один.
// Кстати, именно так работает алгоритм сортировки слиянием (merge sort)

// Примечания
// Ообъединять слайсы до сортировки не допустимо.

package main

import (
	"slices"
)

func SortAndMerge(left, right []int) []int {
	slices.Sort(left)
	slices.Sort(right)

	NewSlice := make([]int, 0, len(left)+len(right))

	// Слияние двух отсортированных слайсов
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			NewSlice = append(NewSlice, left[i])
			i++
		} else {
			NewSlice = append(NewSlice, right[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы
	NewSlice = append(NewSlice, left[i:]...)
	NewSlice = append(NewSlice, right[j:]...)

	return NewSlice
}
