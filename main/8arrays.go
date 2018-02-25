package main

import "fmt"

func main(){

	// this is array of type [3]int, size is part of the array definition
	// all elements initialized to 0
	// array indexes start from 0 :) its obvious, no so obvious for people from Lua Pascal Fortran :-D
	var a [3]int

	// set value
	a[1] = 10;

	fmt.Println("set:",a)
	fmt.Println("element:",a[2])

	// initialize and define in one statement
	b:=[5]int{10,20,30,40,50}

	// we can slice the array like this.
	c:=b[1:2];
	fmt.Println("sliced:",c)

	// two dimenational array

	var twoD [2][3]int;

	for i:=0;i<2;i++ {
		for j:=0;j<3;j++ {
			twoD[i][j] = i+j;
			}
	}

	fmt.Println(twoD)


}
