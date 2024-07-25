package main

import (
	"fmt"
)

func main() {
	//simple switch

	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("one")

	// case 2:
	// 	fmt.Println("one")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	//mulipule condition switch

	// switch time.Now().Weekday(){
	// case time.Saturday,time.Sunday:
	// 	fmt.Println("it's weekend")
	// default:
	// 	fmt.Println("it's workday")
	// }


	//type switch

	whoAmI := func (i interface{})  {
		switch t :=i.(type) {
		case int:
			fmt.Println("It's integer")
		case string:
			fmt.Println("It's String")
		case bool:
			fmt.Println("It's Boolean")
		default:
			fmt.Println("other",t)
		}
	}

	whoAmI(true)
}