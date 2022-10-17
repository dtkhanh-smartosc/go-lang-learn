package main

import (
	"fmt"
	"sample/function"
)

func main() {
	message := function.SayHello("123")
	fmt.Println(message)
}

/*
calling other function from other file

1. Create new module as normal
2. import module
	- publish new module to internet
	- or use: go mod edit replace for the purpose of locating the dependency
3. run: go mod tidy to to synchronize the module's dependencies
4. run: go run .


*/
