package piscine

func Sqrt(nb int) int {

	if nb <= 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		if i > nb/i {
			return 0
		} else if i == nb/i {
			return i
		}
	}
	return 0
}
