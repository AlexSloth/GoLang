package lesson2

import "math"

// FizzBuzz возвращает FizzBuzz, Fizz, Buzz или пустую строку для обычного числа.
// Для n <= 0 возвращает invalid.
func FizzBuzz(n int) string {
	// TODO: Реализовать через switch без выражения.
	switch {
	case n <= 0:
		return "invalid"
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return ""
	}
}

// Sum возвращает сумму. Для пустого среза и nil возвращает 0.
func Sum(values []int64) int64 {
	// TODO: Пройти по значениям циклом.
	var tmp int64
	tmp = 0
	for i := range values {
		tmp += values[i]
	}

	return tmp
}

// Min возвращает минимум. Для пустого среза и nil возвращает 0 по условию задачи.
func Min(values []int64) int64 {
	// TODO: Учесть пустой ввод, затем сравнить элементы.
	if len(values) == 0 {
		return 0
	}
	var tmp int64
	tmp = math.MaxInt64
	for _, i := range values {
		if tmp > i {
			tmp = i
		}
	}
	return tmp
}

// Max возвращает максимум. Для пустого среза и nil возвращает 0 по условию задачи.
func Max(values []int64) int64 {
	// TODO: Учесть пустой ввод, затем сравнить элементы.
	if len(values) == 0 {
		return 0
	}
	var tmp int64
	tmp = math.MinInt64
	for _, i := range values {
		if tmp < i {
			tmp = i
		}
	}
	return tmp
}

// Average возвращает среднее как float64. Для пустого среза и nil возвращает 0.
func Average(values []int64) float64 {
	// TODO: Выполнить преобразования до деления.
	if len(values) == 0 {
		return 0
	}
	var tmp int64
	tmp = Sum(values)
	var res float64
	var tmp1 float64
	var tmp2 float64
	tmp2 = float64(tmp)
	tmp1 = float64(len(values))
	res = tmp2 / tmp1
	return res
}

// Category возвращает invalid, small, regular или large по сумме в копейках.
func Category(amount int64) string {
	// TODO: Проверить границы через switch без выражения.
	switch {
	case amount < 0:
		return "invalid"
	case amount < 10000:
		return "small"
	case amount < 50000:
		return "regular"
	case amount >= 50000:
		return "large"
	default:
		return "invalid"
	}
}

// RuneCount считает кодовые точки Unicode в строке через range.
func RuneCount(text string) int {
	// TODO: Не использовать len(text) как количество кодовых точек.
	var res int
	res = 0
	for range text {
		res++
	}

	return res
}
