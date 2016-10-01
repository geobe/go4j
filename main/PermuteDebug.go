package main

import (
	"fmt"
	"github.com/geobe/go4j/permute"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

func main() {
	size := int(11)
	cpus := runtime.NumCPU()
	fmt.Printf("Prozessoren: %d\n", cpus)
	runtime.GOMAXPROCS(cpus)
	fmt.Printf("GOGC was %d%%\n", debug.SetGCPercent(-1))
	perm := permute.NewPermutation(size - 1)
	fmt.Println(perm)
	start := time.Now()
	for {
		_, done := perm.Next()
		//fmt.Println(seq)
		if done {
			break
		}
	}
	done := time.Now()
	fmt.Printf("\nproduce all permutations of %d took %v \n", size-1, done.Sub(start))
	start = time.Now()
	parTestPermuteTail(size, cpus)
	done = time.Now()
	fmt.Printf("\nproduce all tail permutations of %d took %v \n", size, done.Sub(start))
}

func parTestPermuteTail(size, cpus int) {
	//fmt.Println("running TestPermuteTail")
	perm := permute.NewPermutation(size)
	var head permute.Permutation
	var wg sync.WaitGroup
	var waiting int = 0
	//wg.Add(size)
	for i := 0; i < len(perm); i++ {
		wg.Add(1)
		waiting++
		head = perm.PermuteFirst(int(i))
		go permuteTail(&wg, head)
		if waiting == cpus {
			wg.Wait()
			waiting = 0

		}
	}
	wg.Wait()
}

func permuteTail(wg *sync.WaitGroup, perm permute.Permutation) {
	defer wg.Done()
	fmt.Println(perm)
	start := time.Now()
	for {
		//fmt.Println(perm)
		_, done := perm.NextTail()
		if done {
			end := time.Now()
			fmt.Printf("tail permutations of %d took %v \n", len(perm), end.Sub(start))
			break
		}
	}
}