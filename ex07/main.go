package main
import (
		"fmt"
		"piscine"
)
func main() {
		fmt.Println(piscine.FindNextPrime(5))
		fmt.Println(piscine.FindNextPrime(4))
		fmt.Println(piscine.FindNextPrime(0))
		fmt.Println(piscine.FindNextPrime(1))
		fmt.Println(piscine.FindNextPrime(17))
		fmt.Println(piscine.FindNextPrime(18))
		fmt.Println(piscine.FindNextPrime(100))
		fmt.Println(piscine.FindNextPrime(2147483646))
		fmt.Println(piscine.FindNextPrime(2147483647))
}