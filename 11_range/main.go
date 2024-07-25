package main

import (
	"fmt"
)

func main() {
	// 	nums := []int{6, 7, 8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	// sum := 0
	// for i, num:= range nums{
	// 	fmt.Println(i,num)
	// sum = sum + num
	// fmt.Println(num)
	// fmt.Println(sum)
	// }

	// m := map[string]string{"fname": "mizanur", "lname": "rahman"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	//unicode code point rune
	//starting byte of rune
	//255 => 1 byte or more than byte
	for i, c := range "Golang" {
		fmt.Println(i, c)
	}
}
