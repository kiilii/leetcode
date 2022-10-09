package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3))    //III
	fmt.Println(intToRoman(4))    //IV
	fmt.Println(intToRoman(9))    //IX
	fmt.Println(intToRoman(58))   //LVIII
	fmt.Println(intToRoman(1994)) //MCMXCIV
}

// 执行用时：4 ms, 在所有 Go 提交中击败了 91.74% 的用户
// 内存消耗：3.1 MB, 在所有 Go 提交中击败了 49.95%的用户
func intToRoman(num int) string {
	var (
		values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		chars  = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	)

	var ans string

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			ans += chars[i]
		}
	}

	return ans
}
