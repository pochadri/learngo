package main

import "fmt"
//import "math"
import "time"
func main(){
	i:=2;
	// simple switch
	switch i {
	case 1:
		fmt.Println(" one ")
	case 2:
		fmt.Println( " two")
	case 3:
		fmt.Println("three")
	}

	// you can seperate multiple conditions in case using a comma
	// using default as well

	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday :
		fmt.Println("weekend");
	default:
		fmt.Println("weekday");
	}

	// switch without statment is alternative way to express if else logic

	t:=time.Now();
	switch  {
	case t.Hour()>12 :
		fmt.Println("after noon")
	default:
		fmt.Println("before noon")
	}

	// we can compare the type as well in switch , not doing it till i understand interfaces.

}
