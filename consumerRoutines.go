package main

import (
	"fmt"
	"sync"
)

func show(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Show method")

}

func wgMain() {
	//show()
	var wg sync.WaitGroup
	wg.Add(1)
	go show(&wg)
	wg.Wait()
}

func showRout(ch chan<- string) {
	//fmt.Println("Showrout")
	for i := 0; i < 5; i++ {
		ch <- "From showRout"
	}
	close(ch)
}

func mainRoutines() {
	ch := make(chan string, 1)
	go showRout(ch)

	// for msg := range ch{
	// 	fmt.Println(msg)
	// }
	keepReading := true
	for keepReading {
		select {
		case m, open := <-ch:
			if open == false {
				keepReading = false
			}
			fmt.Println(m)

		default:
			//	fmt.Println("No msg")

		}
	}

	m := map[int]string{}
	var mtx sync.Mutex

	for i := 0; i < 10; i++ {
		go mWriter(i, m, &mtx)
	}

	fmt.Println(m)
	var arr = make([]int, 5)
	updateArr(arr, 111)
	fmt.Println(arr)

	s := arr[:5]
	updateArr(s, 5)
	fmt.Println(s)

}

func mWriter(i int, m map[int]string, mtx *sync.Mutex) {
	mtx.Lock()
	defer mtx.Unlock()
	m[i] = fmt.Sprintf("%d", i)
}

func updateArr(a []int, mul int) {
	for i := 0; i < 5; i++ {
		a[i] = i * mul
	}
}

func cmain() {
	//mainRoutines()
	vv := make([]int, 5)
	vv[1] = 1
	vv[2] = 5
	v := vv[:]
	f(v)
	fmt.Println(v)
}

func f (s []int){
	
	(s)[0] = 1
	(s)[3] = 3
}