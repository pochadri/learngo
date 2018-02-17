package main

import(
	"fmt"
)

func main(){

	// loop with condition  and initialization
	for i:=0;i<5;i++{
		fmt.Println(i)
	}
	i := 0;
	// loop wihh only condition
	for i<3{
		fmt.Println(i);
		i++;
	}

	// infinite loop
	for {
		fmt.Println("infinite print")
		break;
	}
}
