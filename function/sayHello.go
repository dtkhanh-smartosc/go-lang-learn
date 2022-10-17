package function

import "fmt"

func SayHello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
