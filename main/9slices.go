package main

import (
	"fmt"
)

func main(){

	// slice doesn't need length, whereas arrays need it
	// length and capacity are same for this case
	s:=[]int{1,2,3};
	fmt.Println(s);
	fmt.Println("len:",len(s));
	fmt.Println("cap:",cap(s));

	// this is another way to create if you dont want to initialize slice
	// second parameter is length and third parameter is capacity
	s1 := make([]int,10,20);
	s1[0] = 100;
	fmt.Println(s1);
	fmt.Println("len:",len(s1));
	fmt.Println("cap:",cap(s1));

	// append more elements to create a new slice
	// as we appended length modified by capacity is not changed
	s2:= append(s1, 1,2,3 )
	fmt.Println(s2);
	fmt.Println("len:",len(s2));
	fmt.Println("cap:",cap(s2));
	// old slice is also still available
	fmt.Println(s1)

	// modifying s1 slice
	s1[2] = 90;

	// s2 internally points to the same array, so modifying the content will change the view to s1 and s2
	fmt.Println(s2);
	// old slice is also still available
	fmt.Println(s1)

	// this is the real slice, we have created the slice from existing slice
	s3:= s1[2:8];
	fmt.Println("sliced:",s3)
	// resliced to include all the elements till capacity
	s3 =s1[2:cap(s1)]
	fmt.Println(s3)


}