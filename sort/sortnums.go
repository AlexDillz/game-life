//Напишите программу, содержащую функцию SortNums(nums []uint),
//которая сортирует слайс nums по возрастанию

package main

import (
	"slices"
)

func SortNums(nums []uint) {
	slices.Sort(nums)
}
