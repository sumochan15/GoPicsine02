package piscine

const INT_MAX = int(^uint(0) >> 1)

func FactorialLimit() int {
		if INT_MAX == 2147483647 {
				return 12
		}
		return 20
}

func Calculation(nb int) int {
		if nb == 0 {
				return 1
		}
		answer := nb * Calculation(nb-1)
		return answer
}

func RecursiveFactorial(nb int) int {
		if nb < 0 || FactorialLimit() < nb {
				return 0
		}
		return Calculation(nb)
}