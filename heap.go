package main

/**
 * @input A : Integer array
 * @input B : Integer
 *
 * @Output Integer array.
 */
import (
	"container/heap"
	"fmt"
)

// func solve(A []int , B int )  ([]int) {
//     sort.Ints(A)
//     return A[len(A)-B:]
// }

type intNum []int

func (i intNum) Len() int               { return len(i) }
func (i intNum) Less(a int, b int) bool { return i[a] > i[b] }
func (i intNum) Swap(a int, b int)      { i[a], i[b] = i[b], i[a] }

func (i *intNum) Push(num interface{}) {
	*i = append(*i, num.(int))
}

func (i *intNum) Pop() interface{} {
	t := *i
	ele := t[len(t)-1]
	*i = t[:len(t)-1]
	return ele
}

func solve(A []int, B int) []int {
	var h = make(intNum,0, len(A))

	heap.Init(&h)
	for i := range A{
		heap.Push(&h, A[i])
	}
	//fmt.Println(h)
	// fmt.Println("pop ",heap.Pop(&h))
	// fmt.Println("pop ",heap.Pop(&h))
	// fmt.Println("pop ",heap.Pop(&h))
	// fmt.Println("pop ",heap.Pop(&h))
	// fmt.Println("pop ",heap.Pop(&h))

	if len(h)-B > 0{
		return h[len(h)-B:]
	}

	return intNum{}
	
}

func hmain() {
	A := []int{1000, 12, 3, 4, 50}
	out := solve(A, 2)
	fmt.Println(out)
}
