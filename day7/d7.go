package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

var (
	wg     sync.WaitGroup
	x      int64 = 0
	lock   sync.Mutex
	rwlock sync.RWMutex // only w lock will actually lock the resource. It can be used when read times are a lot larger than write times
	once   sync.Once
	m2     sync.Map
)

func main() {
	// mySelect()
	// testAdd()
	// testRWlock()
	// testMap()
	// testSyncMap()
	testAtomicAdd()
}

// goroutine workerpool
func worker(id int, jobs <-chan int /*read only*/, results chan<- int /*write only*/) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}

}

func workerTest() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// start 3 goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5 tasks
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	// output result
	for a := 1; a <= 5; a++ {
		x := <-results
		fmt.Printf("resut:%d\n", x)
	}
}

// use goroutine to calc the sume of each digits

func myRand(jc chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		jc <- newJob
	}
}

func getValue(jc <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	// get value from job chan
	for {
		job := <-jc
		sum := calcDigits(job.value)
		newResult := &result{
			job: job,
			sum: sum,
		}

		resultChan <- newResult
	}
}

func calcDigits(n int64) (dSum int64) {
	for n > 10 {
		r := n % 10
		n /= 10
		dSum += r
	}
	return dSum
}

func myCalc() {
	wg.Add(1)
	go myRand(jobChan)

	// strat 24 goroutine
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go getValue(jobChan, resultChan)
	}

	for result := range resultChan {
		fmt.Printf("value: %d sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()
}

// select 多路复用. randomly select from channels

func mySelect() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

// sync
// lock mutex

func add() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func testAdd() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

func read() {
	defer wg.Done()
	// lock.Lock()
	rwlock.RLock()
	defer rwlock.RUnlock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	// lock.Unlock()

}

func write() {
	defer wg.Done()
	// lock.Lock()
	rwlock.Lock()
	defer rwlock.Unlock()
	x += 1
	time.Sleep(time.Millisecond * 5)
	// lock.Unlock()

}

func testRWlock() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	time.Sleep(time.Second)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

// sync.Once // delay load. singleton

// sync.Map. Map is not thread safe
// Store, Load, LoadOrStore, Delete, Range

var m = make(map[string]int)

func getMap(key string) int {
	return m[key]
}

func setMap(key string, value int) {
	m[key] = value
}

func testMap() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			setMap(key, n)
			fmt.Printf("key=:%v, v:=%v\n", key, getMap(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func testSyncMap() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			v, _ := m2.Load(key)
			fmt.Printf("key=:%v, v:=%v\n", key, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// atomic

func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func testAtomicAdd() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go atomicAdd()
	}
	wg.Wait()
	fmt.Println(x)
	fmt.Println(time.Now().Sub(start))
}
