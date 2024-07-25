package main

import "fmt"
//any type interface
func sum(nums ...int) int{
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}
func main() {
	// fmt.Println(1,2,3,4,5) //variadic function
	// result := sum(3,6,2,8)
	nums :=[]int{2,3,5}
	result := sum(nums...)
	fmt.Println(result)
}
