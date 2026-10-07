package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8} //automatically count the number of elements
	//   Fixed Size: Even though you use [...] to let Go
	//   calculate the size for you, the resulting array is
	//    still fixed in size once compiled. You cannot add
	//    or remove elements later.
	// array_name := [length]datatype{values} // here length is defined
	// or
	// array_name := [...]datatype{values} // here length is inferred

	fmt.Println(arr1)
	fmt.Println(arr2)
}
