package main

import (
	"fmt"
	"time"
)

func main(){
	i := 2
	switch i {
	case 1 : 
		fmt.Println("one")
	case 2 :
		fmt.Println("two")
	case 3 :
		fmt.Println("three")
	}

	day := time.Now().Weekday()
	fmt.Println(day)
	
	switch day{
	case time.Saturday, time.Sunday :
		fmt.Println("It's a weekend")
	default : 
		fmt.Println("It's a weekday")
	}
	
	whatAmI := func(i interface{}){
		switch i.(type){
		case bool:
			fmt.Println("I'm bool")
		case int:
			fmt.Println("I'm int")
		default:
			fmt.Println("Don't know the type")
		}
	}

	whatAmI(true)
	whatAmI(2)
	whatAmI("hello")

}