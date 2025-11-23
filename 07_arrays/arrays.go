package main

import "fmt"

// numbered sequence of specific length
func main() {
	// zeroed values
	// int -> 0, string -> "", bool -> false

	// var nums [4]int

	// nums[0] = 1

	// // array length
	// fmt.Println(len(nums))
	// fmt.Println(nums[0])

	// var vals [4]bool
	// vals[1] = true
	// fmt.Println(vals)

	// var names [4]string
	// names[0] = "golang"
	// fmt.Println(names)

	// To declare it in single line
	// nums := [3]int{1, 2, 3}
	// fmt.Println(nums)

	// 2d arrays
	nums := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(nums)

	// - fixed size, that is predictable
	// - Memory optimization
	// - Constant time access

}
