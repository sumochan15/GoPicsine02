package piscine

const INT_MAX = int(^uint(0) >> 1)

func optimized(count, arg, num1, num2 int) int {
		if num1 > INT_MAX-num2 {
				return -2
		} else if count == arg {
				return num2
		}
		return optimized(count+1, arg, num2, num1+num2)
}

func Fibonacci(index int) int {
		if index < 0 {
				return -1
		} else if index <= 1 {
				return index
		}
		return optimized(1, index, 0, 1)
}

