package main

import "fmt"

type rect struct {
	width, height int
}

// This area method has a receiver type of *rect.
// use a pointer receiver type to avoid copying on method calls or
// to allow the method to mutate the receiving struc
func (r *rect) area() int {
	fmt.Printf("address r, %v \n", r)
	fmt.Printf("r , %v \n", *r)

	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// Here’s an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// Go automatically handles conversion between values and pointers for method calls
func mothodExample() {
	fmt.Println("-----------mothodExample-------------")
	// defer execute when function finished all works and before close; run in LIFO order
	// params value declaired before defer will change
	// ex:
	// a : 10
	// defer fmr.println(a)
	// a = 20

	//  main function -> defer -> panic

	// recover used to recover from panics, only useful in deferred functions,
	// current function still stop but higher function will continue
	defer fmt.Println("----------end---------")
	r := rect{10, 20}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
