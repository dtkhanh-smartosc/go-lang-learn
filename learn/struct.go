package main

import "fmt"

type person struct {
	name string
	age  int
}
type worker struct {
	person
	company string
	salary  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

// Go dont have inherit but it has embed, has-a
func structExample() {
	fmt.Println("---------structExample-----------")

	wk1 := worker{}
	wk1.name = "khanhdt"
	wk1.age = 22
	wk1.company = "smosc"
	wk1.salary = 1

	fmt.Println(wk1)
	wk1.person = person{name: "maidt", age: 12}
	fmt.Println(wk1)

	wk2 := worker{
		person:  person{name: "Duongdt", age: 12},
		company: "smosc",
		salary:  12,
	}
	fmt.Println("wk2", wk2)

	bob := person{"Bob", 20}
	fmt.Println(bob.name)
	fmt.Println(bob)
	fmt.Println(&bob)
	// can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})
	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(*newPerson("Jon"))

	// dont need to make new struct for one time struct
	aDoctor := struct{ name string }{name: "hello world"}
	fmt.Println("aDoctor", aDoctor)

	sp := bob
	// this is copy value from bob to sp
	fmt.Println("\n", sp)
	// struct fields with a dot.
	sp.age = 23
	fmt.Println(sp)
	fmt.Println(bob)

	prtsp := &bob
	// this make prtsp`s value = &bob => by changing prtsp result in bob too
	fmt.Println("\n", prtsp)
	// can also use dots with struct pointers - the pointers are automatically dereferenced.
	prtsp.age = 32
	fmt.Println(prtsp)
	fmt.Println(*prtsp)
	fmt.Println(bob)

}
