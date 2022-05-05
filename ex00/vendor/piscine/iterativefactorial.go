package piscine

const INT_MAX = int(^uint(0) >> 1)

func IterativeFactorial(nb int) int {
		if nb < 0 {
				return 0
		}

		factorial := 1

		for i := nb; i > 0; i-- {
				factorial *= i
				if factorial > INT_MAX/i {
						return 0
				}
		}
		return factorial
}