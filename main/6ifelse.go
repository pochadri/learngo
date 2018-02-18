package main

import "fmt"

func main(){

	// brackets around condition are not mandatory
	// flower bracket after condition is mandatory
	if 5%2==0 {
		fmt.Println("its true")
	}

	// A statement can preceed conditional
	// it would be available in the else blocks as well.
	if num:=-10; num>9 {
		fmt.Println(num," is double digits")
	} else if num <0{
		fmt.Println(num," is negative")
	} else {
		fmt.Println(num," is single digit")
	}
}