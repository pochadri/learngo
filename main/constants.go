package main

import (
	"fmt"
);

//global variables

const a int = 100;
func main(){

	// type is not required for const, it is derived based on the context it is used.
	const x  = 10

	//calculated constant
	const calculated float64 = 0.000034/5.5

	// code is optimized by replacing all the const's with value during compile time.
	fmt.Println(calculated)

	fmt.Println(x)
	PrintValue(a)
}

/**
print a value given as input
 */
func PrintValue(value int){
	fmt.Println(value);
}

