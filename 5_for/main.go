package main

import "fmt"

// for -> only const in go looping
func main()  {
	//while loop

	// i :=1
	// for i <= 3{
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	//infinity loop
	// for {
	// 	fmt.Println("i")
	// }

	//classic for loop
	// for i:=0; i<= 3; i++{
	// 	// break
		
	// 	if i == 2{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//range new 1.22

	for i:= range 3{
		fmt.Println(i)
	}
}