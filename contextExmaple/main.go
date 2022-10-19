package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	//  it’s important to know that the values stored in
	//  a specific context.Context are immutable,
	// 	meaning they can’t be changed. When you called the context.WithValue
	//  you passed in the parent context, and you also received a context back.
	//	You received a context back because the context.WithValue function
	//	didn’t modify the context you provided. Instead,
	//	it wrapped your parent context inside another one with the new value.

	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))

	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")
	doAnother(anotherCtx)

	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
}
func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("myKey"))
}

// a way to signal to any functions using it that the context has ended and
// should be considered complete. By signaling to these functions that
// the context is done, they know to stop processing any work related to
// the context that they may still be working on

func main() {
	// context.TODO function, one of two ways to
	// create an empty (or starting) context
	// ctx := context.TODO()

	// it’s designed to be used where you intend to start a known context
	ctx := context.Background()
	//doSomething(ctx)

	// One benefit of using context.Context in a program is
	// the ability to access data stored inside a context

	// By adding data to a context and passing the context
	// from function to function, each layer of a program
	// can add additional information about what’s happening
	ctx = context.WithValue(ctx, "myKey", "my-example-Value")
	doSomething(ctx)
	// fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-2))
}
