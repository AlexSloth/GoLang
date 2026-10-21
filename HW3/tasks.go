package homework

// PairInput соответствует объекту nums/target из условия контеста.
type PairInput struct {
	Nums   []int
	Target int
}

// longestPrefix возвращает самый длинный общий префикс строк.
func longestPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}
	prefix := words[0][:]
	for i, currentletter := range prefix {
		for _, word := range words[1:] {
			currentword := []rune(word)
			if currentletter != currentword[i] {
				return string(prefix[:i])
			}
		}
	}
	return string(prefix)
}

// pairIndices возвращает индексы двух разных элементов по возрастанию.
func pairIndices(input PairInput) [2]int {
	for i := 0; i < len(input.Nums)-1; i++ {
		for j := i + 1; j < len(input.Nums); j++ {
			if input.Nums[i]+input.Nums[j] == input.Target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
}

// hasDuplicate сообщает, встречается ли какое-либо значение повторно.
func hasDuplicate(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// compactSorted создаёт независимый слайс уникальных значений.
func compactSorted(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := make([]int, 0, len(nums))
	result = append(result, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			result = append(result, nums[i])
		}
	}
	return result
}

// balancedBrackets проверяет пары и порядок закрытия скобок.
func balancedBrackets(text string) bool {
	if len(text) == 0 {
		return true
	}
	if len(text)%2 == 1 {
		return false
	}
	stack := make([]rune, len(text))
	head := 0
	for _, i := range text {
		switch {
		case i == '(' || i == '[' || i == '{':
			stack[head] = i
			head++
		case head == 0:
			return false
		case i == ')' && stack[head-1] == '(':
			head--
		case i == ']' && stack[head-1] == '[':
			head--
		case i == '}' && stack[head-1] == '{':
			head--
		default:
			return false
		}
	}
	if head != 0 {
		return false
	}
	return true
}
