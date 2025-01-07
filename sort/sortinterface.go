//На предприятии работают несколько cсотрудников. Каждый из них имеет свою должность,
// фиксированную месячную заработную плату, и стаж работы.
// Напишите программу в котором тип Company реализует следующий интерфейс:
// type CompanyInterface interface{
// AddWorkerInfo(name, position string, salary, experience uint) error
// SortWorkers() ([]string, error)
// }

// AddWorkerInfo - метод добавления информации о новых сотрудниках, где name - имя сотрудника, position - должность, salary - месячная зароботная плата, experience - стаж работы (месяцев).
// SortWorkers - метод сортировки, который сортирует список сотрудников по следующим свойствам: доход за время работы на предприятии(по убыванию), должность (от высокой) и возвращает слайсл формата: имя - доход - должность.
// Допустимые должности в порядке убывания: "директор", "зам. директора", "начальник цеха", "мастер", "рабочий".

// Примечания
// Пример отсортированного вывода:
// Михаил - 12000 - директор
// Андрей - 11800 - мастер
// Игорь - 11000 - зам. директора

package main

import (
	"errors"
	"fmt"
	"sort"
)

// Интерфейс компании
type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

// Структура сотрудника
type Worker struct {
	Name       string
	Position   string
	Salary     uint
	Experience uint
}

// Структура компании
type Company struct {
	workers []Worker
}

// Порядок должностей
var positionPriority = map[string]int{
	"директор":       5,
	"зам. директора": 4,
	"начальник цеха": 3,
	"мастер":         2,
	"рабочий":        1,
}

// Реализация метода AddWorkerInfo
func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	if _, exists := positionPriority[position]; !exists {
		return errors.New("недопустимая должность")
	}
	c.workers = append(c.workers, Worker{Name: name, Position: position, Salary: salary, Experience: experience})
	return nil
}

// Реализация метода SortWorkers
func (c *Company) SortWorkers() ([]string, error) {
	if len(c.workers) == 0 {
		return nil, errors.New("нет сотрудников для сортировки")
	}

	// Сортировка сотрудников
	sort.SliceStable(c.workers, func(i, j int) bool {
		worker1, worker2 := c.workers[i], c.workers[j]

		// Сравнение по доходу
		income1 := worker1.Salary * worker1.Experience
		income2 := worker2.Salary * worker2.Experience
		if income1 != income2 {
			return income1 > income2
		}

		// Сравнение по должности
		return positionPriority[worker1.Position] > positionPriority[worker2.Position]
	})

	// Форматирование результата
	var result []string
	for _, worker := range c.workers {
		income := worker.Salary * worker.Experience
		result = append(result, fmt.Sprintf("%s - %d - %s", worker.Name, income, worker.Position))
	}
	return result, nil
}
