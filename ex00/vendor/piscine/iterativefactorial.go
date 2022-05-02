package piscine

func IterativeFactorial(nb int) int {
	if nb < 0{
		return 0
	} 

	answer := 1

	for i:=nb; i > 0; i--{
		answer *= i
		if answer > 9223372036854775807/i{
			return 0
		}
	}
	return answer
}

//intMaxの設定はどうする？環境による？？