package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating maps

	// m := make(map[string]string)

	//setting an element
	// m["Name"] = "golang"
	// m := make(map[string]int)
	// m["Age"] = 30
	// m["Price"] = 23
	// fmt.Println(m["Age"])
	// // fmt.Print(len(m))

	// // delete(m,"Age")
	// clear(m)
	// fmt.Println(m)

	// m:=map[string]int{"price":23,"phone":3}

	// k, ok :=m["phone"]
	// fmt.Println(k)
	// if ok{
	// 	fmt.Println("All okay")
	// }else{
	// 	fmt.Println("Not okay")
	// }

	// fmt.Println(m)
	m1 :=map[string]int{"phone":3,"age":23}
	m2 :=map[string]int{"phone":3,"age":23}

	fmt.Println(maps.Equal(m1,m2))
}