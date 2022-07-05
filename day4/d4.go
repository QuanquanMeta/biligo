package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//var wg sync.WaitGroup
var wg sync.WaitGroup
var once sync.Once

func main() {
	// mySyncTest()
	// time.Sleep(time.Second)
	// myGoMaxProcs()
	// myChan()
	// myChanTest()
	myCloseChan()
}

// concurrent: 1 goroutine 2 channel
// goroutine. Add go in front of a function
// GMP. G = goroutine M = machine thread P = processor
// m:n m of goroutine assgin to n thread
// a groutine starts with 2K

func hello(i int) {
	count := 100
	for i := 0; i < count; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func mySync() {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10) // 0 <= x< 10
		fmt.Println(-r1, -r2)
	}
}

func mySyncTest() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go mySync()
	}
	wg.Wait() // wait for the counter of wg to be 0
}

// GOMAXPROCS
func myGoMaxProcs() {
	runtime.GOMAXPROCS(4)
	fmt.Println(runtime.NumCPU())
	for i := 0; i < 2; i++ {
		wg.Add(2)
		go myA()
		go myB()
	}

	wg.Wait()
}

func myA() {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func myB() {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

// channel is type. it is FIFO
// chan <- // towards left
// 1. send ch1 <- 1
// 2. recieve x :=  <- ch1
// 3. close()
// channel is reference type. It needs to be initialised.
// make. slice/map/chan
// var variable chan element type
// go concurrent model is CSP (communicating Sequential Processes)
// use communication (channel) to share memory. not to share the memory to communicate

func myChan() {
	// var a []int    // slice
	// var b chan int // chan
	ch := make(chan int, 16) // no size. wait for receiver // normally small, if the value is large then use pointer
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-ch
		fmt.Printf("%d receive the value from ch\n", x)
	}()
	ch <- 10
	fmt.Println("10 send to the chan ch")
	c := make(chan int, 1) // with size. has 16 size
	fmt.Println(c)
	close(c)
	wg.Wait()
}

// start a goroutine, send 100 to a ch1
// start a goroutine, get avlue from ch1. Calculate the square, send it to ch2
func myCh1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func myCh2(ch1, ch2 chan int) {
	defer wg.Done()

	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() { close(ch2) })
}

func myChanTest() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	wg.Add(3)
	go myCh1(ch1)
	go myCh2(ch1, ch2)
	go myCh2(ch1, ch2)

	for ret := range ch2 {
		fmt.Println(ret)
	}
}

// close channel. for sending ch1 <-10 close(ch1)
//x, ok := <-ch1
func myCloseChan() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)

	// for x := range ch1 {
	// 	fmt.Println(x)
	// }

	x, ok := <-ch1
	fmt.Println(x, ok)
	x, ok = <-ch1
	fmt.Println(x, ok)
	x, ok = <-ch1
	fmt.Println(x, ok)
}

// oneway chan
// normally used with parameter to restric send only or receive only
func mySendOnlyChan(ch1 chan<- int) { // send-only, write only
	ch1 <- 10
}

func myReceivenlyChan(ch2 <-chan int) { // receive-only, read only
	x, ok := <-ch2
	fmt.Println(x, ok)
}

// worker goroutine pool
