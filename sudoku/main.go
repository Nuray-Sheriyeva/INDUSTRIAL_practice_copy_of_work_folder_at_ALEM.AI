package main

import (
	"fmt"
	"os"
)

// grid - это наша доска судоку 9x9.
// 0 означает пустую клетку, числа 1-9 - заполненную клетку.
var grid [9][9]int

// solvedGrid хранит первое найденное решение,
// чтобы его не испортил дальнейший backtracking при поиске второго решения.
var solvedGrid [9][9]int

func main() {
	// Судоку должно прийти ровно 9 строк
	args := os.Args[1:]
	if len(args) != 9 {
		fmt.Println("Error")
		return
	}

	// Заполняем grid из входных строк
	for row := 0; row < 9; row++ {
		line := args[row]

		// Каждая строка должна быть ровно 9 символов
		if len(line) != 9 {
			fmt.Println("Error")
			return
		}

		for col := 0; col < 9; col++ {
			char := line[col]

			if char == '.' {
				grid[row][col] = 0
			} else if char >= '1' && char <= '9' {
				grid[row][col] = int(char - '0')
			} else {
				// Любой другой символ - это ошибка
				fmt.Println("Error")
				return
			}
		}
	}

	// Проверяем, что начальные числа не конфликтуют друг с другом
	if !startIsValid() {
		fmt.Println("Error")
		return
	}

	// Считаем, сколько решений у этого судоку
	count := 0
	count = countSolutions(count)

	// Нам нужно ровно одно решение
	if count != 1 {
		fmt.Println("Error")
		return
	}

	printGrid()
}

// startIsValid проверяет изначально расставленные числа на дубли
func startIsValid() bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			value := grid[row][col]
			if value == 0 {
				continue
			}
			// Временно убираем число, чтобы isValidMove не сравнило его с самим собой
			grid[row][col] = 0
			valid := isValidMove(row, col, value)
			grid[row][col] = value

			if !valid {
				return false
			}
		}
	}
	return true
}

// isValidMove проверяет, можно ли поставить число value в клетку (row, col)
func isValidMove(row int, col int, value int) bool {
	// Проверка строки и столбца
	for i := 0; i < 9; i++ {
		if grid[row][i] == value {
			return false
		}
		if grid[i][col] == value {
			return false
		}
	}

	// Проверка блока 3x3
	blockRowStart := (row / 3) * 3
	blockColStart := (col / 3) * 3

	for r := blockRowStart; r < blockRowStart+3; r++ {
		for c := blockColStart; c < blockColStart+3; c++ {
			if grid[r][c] == value {
				return false
			}
		}
	}

	return true
}

// findEmptyCell ищет первую свободную клетку.
// Если свободных клеток нет, возвращает -1, -1
func findEmptyCell() (int, int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

// countSolutions - это основной алгоритм (backtracking).
// Он пробует все возможные числа в пустых клетках и считает решения.
// Если решений уже 2 - дальше не ищет, чтобы не тратить время.
func countSolutions(count int) int {
	if count >= 2 {
		return count
	}

	row, col := findEmptyCell()

	// Если свободных клеток нет - значит доска полностью и правильно заполнена
	if row == -1 {
		if count == 0 {
			solvedGrid = grid
		}
		return count + 1
	}

	// Пробуем поставить в эту клетку числа от 1 до 9
	for value := 1; value <= 9; value++ {
		if isValidMove(row, col, value) {
			grid[row][col] = value

			count = countSolutions(count)

			// Откатываем назад (backtrack), чтобы попробовать другое число
			grid[row][col] = 0

			if count >= 2 {
				return count
			}
		}
	}

	return count
}

// printGrid печатает решённую доску в требуемом формате
func printGrid() {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if col > 0 {
				fmt.Print(" ")
			}
			fmt.Print(solvedGrid[row][col])
		}
		fmt.Println()
	}
}
