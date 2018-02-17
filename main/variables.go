package main

import "fmt"

func main()  {
	var x int;
	x = 11

	// := is the shorthand for defining and initializing a variable
	y:=10;
	// this is error since variable need to defined before using it.
	//z=10
	// no type is required, go will infer the type for initialized variables.
	var d = true;
	fmt.Println(d);
	fmt.Println(y);
	fmt.Println("value of x",x);

}
