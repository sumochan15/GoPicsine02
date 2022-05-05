package piscine

func IterativePower(nb int, power int) int {
		if power < 0 {
				return 0
		} else if power == 0 {
				return 1
		} else {
				answer := 1
				for i := power; i > 0; i-- {
						answer *= nb
				}
				return answer
		}
}