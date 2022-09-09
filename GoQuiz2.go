package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
	// "sort"
)

func Problem1(x int) {
	exSlice := make([]int, x)
	// assign random ints to slice
	for i := 0; i < x; i++ {
		exSlice[i] = rand.Int()
	}
	sum := 0
	// start timing
	t0 := time.Now()
	// sum slice w/o concurrency
	for i := range exSlice {
		sum += exSlice[i]
	}
	// end timing
	t1 := time.Now()
	fmt.Printf("Summing a slice ***without*** concurrency of len %d took %d\n", len(exSlice), t1.Sub(t0))
	fmt.Printf("The sum of the slice is %d\n", sum)

	// make buffered channels, defer close channels until done using
	c1 := make(chan int, x/2)
	c2 := make(chan int, x/2)
	defer close(c1)
	defer close(c2)

	tmp1 := 0
	tmp2 := 0

	// go routine to sum first half of slice
	go func() {
		for i := 0; i < len(exSlice)/2; i++ {
			tmp1 = exSlice[i]
			c1 <- tmp1
		}
	}()

	// go routine to sum second half of slice
	go func() {
		for i := x - 1; i >= (len(exSlice) / 2); i-- {
			tmp2 = exSlice[i]
			c2 <- tmp2
		}
	}()

	sum = 0
	t0 = time.Now() // start timing
	// add to sum when channels are ready
	for range exSlice {
		select {
		case firsthalfsum := <-c1:
			sum += firsthalfsum
		case secondhalfsum := <-c2:
			sum += secondhalfsum
		}
	}
	t1 = time.Now() // end timing
	// fmt.Printf("first half == %d\n", firsthalfsum)
	// fmt.Printf("second half == %d\n", secondhalfsum)
	fmt.Printf("Summing a slice ***with*** concurrency of len %d took %d\n", len(exSlice), t1.Sub(t0))
	fmt.Printf("The sum of the slice is %d\n", sum)

}

// interface for sort.Sort()
type Interface []int

func (a Interface) Len() int           { return len(a) }
func (a Interface) Less(i, j int) bool { return a[i] < a[j] }
func (a Interface) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Problem2(x int) {
	exSlice := make([]int, x)
	for i := 0; i < x; i++ {
		exSlice[i] = rand.Int()
	}

	// sort using sort.Sort()
	t0 := time.Now()
	sort.Sort(Interface(exSlice))
	t1 := time.Now()
	fmt.Printf("sorting a slice of length %d using sort.Sort took %d\n", len(exSlice), t1.Sub(t0))

	// sort using sort.SliceStable()
	t0 = time.Now()
	sort.SliceStable(exSlice, func(i, j int) bool { return exSlice[i] < exSlice[j] })
	t1 = time.Now()
	fmt.Printf("sorting a slice using sort.SliceStable took %d\n", t1.Sub(t0))

}

func main() {
	// Get argument from command line
	x := os.Args[1]
	fmt.Printf("Your input is -> %s\n", x)
	y, _ := strconv.Atoi(x)
	// run programs
	Problem1(y)
	Problem2(y)
}
