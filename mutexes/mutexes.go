package main

import (
	"fmt"
	"github.com/pborman/getopt"
	"runtime"
	"sync"
)

type RdWr struct{
	val		int
	rmtx	sync.Mutex
	wmtx	sync.Mutex
}

func (c *RdWr) gen(wg *sync.WaitGroup, num int) {
	c.wmtx.Lock()
	c.val++
	fmt.Printf("GEN: (%d)\n", num)
	c.rmtx.Unlock()
	wg.Done()
}

func (c *RdWr) post(wg *sync.WaitGroup, num int){
	c.rmtx.Lock()
	fmt.Printf("POST (%d): %d\n", num, c.val)
	c.wmtx.Unlock()
	wg.Done()
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

	c := RdWr{val: 0}
	c.rmtx.Lock()

	for i := 0; i < *optThreads; i++ {
		go c.gen(&w, i)
		go c.post(&w, i)
	}
	w.Wait()
}
