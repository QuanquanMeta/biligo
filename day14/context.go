package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
// why we need context.
var wg sync.WaitGroup
var exitChan chan bool = make(chan bool, 1)

func f() {
	defer wg.Done()
	b := false
	for {

		if b {
			break
		}
		fmt.Println("hello world")
		time.Sleep(time.Microsecond * 500)

		select {
		case <-exitChan:
			b = true
			break
		default:
		}
	}
}

func test() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	// how to stop the goroutine? // use sync.WaitGroup
	// 1. global variable
	// 2. chan
	exitChan <- true
	wg.Wait()
}

func g2(ctx context.Context) {
	defer wg.Done()
	for {

		fmt.Println("hello asia")
		time.Sleep(time.Microsecond * 500)

		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func g(ctx context.Context) {
	defer wg.Done()
	go g2(ctx)
	for {

		fmt.Println("hello world")
		time.Sleep(time.Microsecond * 500)

		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func myContext() {
	start := time.Now()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(2)
	go g(ctx)
	time.Sleep(time.Second * 5)

	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

func connectWithDeadline() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("exit")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// withValue

func main() {
	//test()
	//myContext()
	connectWithDeadline()
}
