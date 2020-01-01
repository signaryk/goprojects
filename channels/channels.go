package main

import (
	"fmt"
	"github.com/pborman/getopt"
	"runtime"
	"sync"
)

func gen(gench chan int, postch chan int, wg *sync.WaitGroup, num int){
	defer wg.Done()

	x := <- postch
	x++
	fmt.Printf("GEN: (%d)\n", num)
	gench <- x
}

func post(gench chan int, postch chan int, wg *sync.WaitGroup, num int){
	defer wg.Done()

	val := <- gench
	fmt.Printf("POST (%d): %d\n", num, val)
	postch <- val
}

func main() {
	optThreads := getopt.IntLong("threads", 't', 1, "Number of threads")
	optHelp := getopt.BoolLong("help", 'h', "Help")
	getopt.Parse()

	if *optHelp {
		getopt.Usage()
		return
	}

	fmt.Printf("GOMAXPROCS = %d\n", runtime.GOMAXPROCS(0))

	var w sync.WaitGroup
	w.Add(2 * *optThreads)

	gench := make(chan int, 1)
	postch := make(chan int, 1)

	postch <- 0

	for i := 0; i < *optThreads; i++ {
		go gen(gench, postch, &w, i)
		go post(gench, postch, &w, i)
	}
	w.Wait()
}
