package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// param is address
func zeroptr(iptr *int) {
	fmt.Println("*iptr before: ", *iptr)
	fmt.Println("iptr before: ", iptr)
	*iptr = 2 // get to the value at current address
	fmt.Println("*iptr after: ", *iptr)
	fmt.Println("iptr after: ", iptr)

}

func prtExample() {
	fmt.Println("---------prtExample---------")
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)

	var a int = 12

	*&a = 11

	fmt.Println(*&a, &a)
}
