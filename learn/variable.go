package main

import (
	"fmt"
	"time"
)

var gbValue int = 999

/*
varable had initialized must be used : Golang
shadow between package and inner-function variable

	package var i = 100
	inner-function var i = 100
	after inner-func var was initialized function will
	use inner-func var instead of package var

	in the same block can not have them same var with same name

	inner-function/ inner-block var can`t be seen out side the block

var name convention stand for var time life, short as yot can

	Pascal or camelCase convention
	lower case first letter for package scope
	upper case first letter to export
	no private scope
*/
func varableExample() {

	fmt.Println("--------varable---------")

	fmt.Println("gb value", gbValue)

	var (
		myName string = "khanhdt"
		myAge  int    = 21
	)
	println("my name", myName)
	println("my age", myAge)

	var a = "initial"
	// print the value and type of variable
	fmt.Printf("%v, %T", a, a)
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	// short hand for var d int = 123
	d := 123

	fmt.Println(d)

	const a1 = 145
	fmt.Println(a1)

	// iota use as a counter  but only work in () case
	// if we do
	// const (
	//  c1 = iota
	//  c2 = iota
	// )
	// iota will not increase
	//  => scope ton constance block

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Printf("%v", c1)
	fmt.Printf("%v", c2)
	fmt.Printf("%v", c3)

}
func switchExample() {
	fmt.Println()
	fmt.Println("--------switchExample---------")
	var i int = 12
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("i dont know")
	}
	whatIsIt := func() {
		switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend")
		default:
			fmt.Println("It's a weekday")
		}
	}

	whatIsIt()

	whatTypeAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case float64:
			fmt.Println("float")
		default:
			fmt.Println("i dont know")
		}
	}
	whatTypeAmI(1)
	whatTypeAmI(true)
	whatTypeAmI(3.3)
}
func ifElseExample() {
	fmt.Println()
	fmt.Println("--------ifElseExample---------")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if num := 123; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
func sliceExample() {
	// slice and map , underlying data => a := b does not mean a copy the whole b,
	// works the same as pointer = pointing to the same memmory
	// vùng memory này cũng tương tự array = vùng nhớ liên tiếp
	fmt.Println("--------sliceExample---------")
	s := make([]string, 3)

	s[0] = "a"
	s[1] = "b1"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	// slices support several more that make them richer than arrays.
	// One is the builtin append, which returns a slice containing
	// one or more new values. Note that we need to accept a return
	// value from append as we may get a new slice value
	// may cause expensive copy operation if unerlying array is too small
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	c := make([]string, len(s))
	copy(c, s) // powerfull operation
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
func arrayExample() {
	fmt.Println()
	fmt.Println("--------arrayExample---------")
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	//  declare and initialize a variable for array in a single line.
	b := [5]int{1, 2, 3, 4, 5}
	//  [...] create an array with the size big enought to put elemnt I provide
	c := [...]int{1, 2, 3, 4}
	fmt.Println(c)
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
func forExample() {
	fmt.Println("--------forExample---------")
	// A classic initial/condition/after for loop.

	var i int = 1
	for i <= 3 {
		fmt.Print(i, " ")
		i++
	}

	fmt.Println()
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n, " ")
	}
	/*
		var devide func(int, int) (int, err)
		devide = func(a, b int) (int, err) {
			...
		}
	*/
}

func mapExample() {
	fmt.Println("--------mapExample---------")
	// thứ tự trong map là không thể đoán trước
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 16
	m["k3"] = 2001

	fmt.Println("map: ", m)
	fmt.Println("k1", m["k1"])
	fmt.Println("len:", len(m))

	delete(m, "k1")
	fmt.Println("map:", m)
	// This can be used to disambiguate between missing keys
	// and keys with zero values like 0 or "".
	// Here we didn’t need the value itself,
	// so we ignored it with the blank identifier _
	_, prs := m["k1"]
	_, prs1 := m["k2"]
	fmt.Println("prs:", prs)
	fmt.Println("prs1:", prs1)

	// short hand
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
func rangeExample() {
	fmt.Println("--------rangeExample---------")

	nums := []int{1, 2, 4, 5}
	sum := 0

	/*
		range on arrays and slices provides both the index
		and value for each entry. Above we didn’t need the index,
		so we ignored it with the blank identifier _.
		Sometimes we actually want the indexes though.
	*/

	for _, num := range nums {
		fmt.Print(num, " ")
		// _ will be the index of array
		sum += num
	}
	fmt.Println("sum: ", sum)

	for a, num := range nums {
		fmt.Print(a, "-", num, " ")
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	fmt.Println()
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}
	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func functionExample(a, b, c int) int {
	fmt.Println("--------functionExample---------")
	return a + b + c
}

// The (int, int) in this function signature shows that the function returns 2 ints.
func returnMultiple() (int, int) {
	fmt.Println("--------returnMultiple---------")
	var (
		i int = 111
		d int = 12
	)
	return i, d
}

func multipleAssigment() {
	fmt.Println("--------multipleAssigment---------")
	// Here we use the 2 different return values from the call with multiple assignment.
	a, b := returnMultiple()
	fmt.Printf("%v, %v", a, b)
}

// can be called with any number of trailing arguments
func variadicFuncExample(nums ...int) {
	fmt.Println("--------variadicFuncExample---------")
	// the type of nums is equivalent to []int
	// can call len(nums), iterate over it with range
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total: ", total)
}
