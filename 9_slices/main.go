package main

import (
	"fmt"
)

//slice -> dynamic arrays
//most used contruct in go
//+ useful methods
func main()  {
	// uninitialized slice is nil
	// var num []int

	// fmt.Println(num == nil)
	// var nums = make([]int,0,5)
	//capacity -> maximum number of elements can fit
	// fmt.Println(cap(nums))
	// fmt.Println(nums ==nil)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// nums :=[]int{}
	// nums = append(nums, 1)
	// nums = append(nums, 3)
	// nums[0]=1
	// fmt.Print(nums)

	// var nums = make([]int,0,5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// copy function
	
	// copy(nums2,nums)
	// fmt.Println(nums)
	// fmt.Println(nums2)

	//slice operator

	// var nums = []int{1,2,3}
	// fmt.Println(nums[0:2])
	// fmt.Println(nums[1:])

	// var nums = []int{1,2}
	// var num1 = []int{1,2}
	// fmt.Println(slices.Equal(nums,num1))
	var nums = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums)

}