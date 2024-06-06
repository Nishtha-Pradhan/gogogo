package main

import (
	"fmt"
	"context"
	"time"
)

func doSomething(ctx context.Context) {
	fmt.Println("Doing Something")
}

func doSomethingWithValue(ctx context.Context) {
	fmt.Printf("doSomethingWithValue: myKey's value is %s\n", ctx.Value("myKey"))

	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")
	doAnother(anotherCtx)

	fmt.Printf("doSomethingWithValue: myKey's value is %s\n", ctx.Value("myKey"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("myKey"))
}

func endDoSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)
	
	printCh := make(chan int)
	go endDoAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		printCh <- num
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("endDoSomething: finished\n")
}

func endDoAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("endDoAnother err: %s\n", err)
			}
			fmt.Printf("endDoAnother: finished\n")
			return
		case num := <-printCh:
			fmt.Printf("endDoAnother: %d\n", num)
		}
	}
}

func deadlineOfDoSomething(ctx context.Context) {
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	defer cancelCtx()

	printCh := make(chan int)
	go endDoAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("deadlineOfDoSomething: finished\n")
}

func timeLimitOfDoSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer cancelCtx()

	printCh := make(chan int)
	go endDoAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("timeLimitOfDoSomething: finished\n")
}

func main() {

	fmt.Println("************* Create empty context **************")

	ctx1 := context.TODO()
	// The context.TODO function is one of two ways to create an empty (or starting) context.
	fmt.Println("ctx1 using TODO method:")
	doSomething(ctx1)

	ctx2 := context.Background()
	// The context.Background function is the other way to create empty context.
	fmt.Println("ctx2 using Background method:")
	doSomething(ctx2)

	fmt.Println("************* Use value with context **************")

	ctxV := context.Background()
	// First you initialize an empty context and then add value to it.
	ctxV = context.WithValue(ctxV, "myKey", "myValue")
	doSomethingWithValue(ctxV)

	/**
	Done channel checks whether a context has ended or not.
	The periodic checking of the Done channel while processing work in-between is done using the Select statement.
	*/
	fmt.Println("************* Working with 'DONE' channel and 'SELECT' statement **************")

	ctxD := context.Background()
	endDoSomething(ctxD)

	/**
	Setting a deadline for when the context needs to be finished, 
	and it will automatically end when that deadline passes.
	*/

	fmt.Println("************* Giving a Context a Deadline **************")

	ctxDl := context.Background()
	deadlineOfDoSomething(ctxDl)

	/**
	Setting a time limit is more useful to use then setting a deadline as this is what we look for in most cases.
	Using the context.WithTimeout function you only need to provide a time.Duration value for how long you want the context to last.
	*/

	fmt.Println("************* Giving a Context a Time Limit **************")

	ctxTl := context.Background()
	timeLimitOfDoSomething(ctxTl)

}
