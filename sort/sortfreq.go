// Дана строка с символами из набора алфавита. Напишите программу с функцией SortByFreq(str string) string,
// которая будет сортировать символы из строки по возрастанию с учетом частоты повторения.
// Символы с наименьшим количеством вхождений должны идти в начале, а символы с наибольшей частотой - в конце.
// В случае одинаковой частоты символов, они должны быть отсортированы в алфавитном порядке.

// Примечания
// Пример:
// Вход: "abbbzzzat"
// Выход: "taabbbzzz"

package main

import (
	"sort"
	"strings"
)

func SortByFreq(str string) string {
	freqmap := make(map[rune]int)
	for _, char := range str {
		freqmap[char]++
	}
	// Создаем слайс для хранения символов
	type charFreq struct {
		char rune
		freq int
	}
	var charFreqs []charFreq

	// Заполняем слайс структур charFreq
	for char, freq := range freqmap {
		charFreqs = append(charFreqs, charFreq{char, freq})
	}

	// Сортируем слайс по частоте, затем по алфавиту
	sort.Slice(charFreqs, func(i, j int) bool {
		if charFreqs[i].freq == charFreqs[j].freq {
			return charFreqs[i].char < charFreqs[j].char
		}
		return charFreqs[i].freq < charFreqs[j].freq
	})

	// Формируем итоговую строку
	var result strings.Builder
	for _, cf := range charFreqs {
		result.WriteString(strings.Repeat(string(cf.char), cf.freq))
	}

	return result.String()
}
