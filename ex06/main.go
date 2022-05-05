package main
import (
		"fmt"
		"piscine"
)
func main() {
		fmt.Println(piscine.IsPrime(5))
		fmt.Println(piscine.IsPrime(4))
		fmt.Println(piscine.IsPrime(0))
		fmt.Println(piscine.IsPrime(1))
		fmt.Println(piscine.IsPrime(97))
		fmt.Println(piscine.IsPrime(31))
		fmt.Println(piscine.IsPrime(42))
		fmt.Println(piscine.IsPrime(100))
		fmt.Println(piscine.IsPrime(101))
		fmt.Println(piscine.IsPrime(2147483647))
}