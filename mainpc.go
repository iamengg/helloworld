package main

import (
	"fmt"
	"math"
)

/*
stack
push
pop
get max
get min
slice

Push element
	On pushing
		if min & max are getting affected update their values
		push node in stack
	On pop
		Not affects min, max for previous levels

*/

type node struct {
	val int
	min int
	max int
}

type stack []*node

func (st *stack) top() *node {
	if len(*st) != 0 {
		return (*st)[len(*st)-1]
	}

	return nil
}

func (st *stack) getMin() int{
	n := *st.top()
	if &n == nil{
		return -1
	}
	return n.min
}


func (st *stack) getMax() int{
	n := *st.top()
	if &n == nil{
		return -1
	}
	return n.max
}

func (st *stack) pop() *node {
	l := len(*st)
	if l != 0 {
		
		lastNode := (*st)[l-1]
		*st = (*st)[:l-1]
		//0 1 2 3 
		//0 1 2
		return lastNode
	}

	return nil
}

func (st *stack) push(val int) {
	min, max := 0, 0

	if len(*st) != 0 {
		n := *st.top()
		if n.max < val {
			max = val
		} else {
			max = n.max
		}

		if n.min > val {
			min = val
		} else {
			min = n.min
		}
	} else {
		min, max = val, val
	}

	ele := node{
		val: val,
		min: min,
		max: max,
	}

	*st = append(*st, &ele)
}

// [16:46] Anand Bhatt
//     anand.bhatt@go-mmt.com
//     9632312060

func mainStack() {
	//create stack
	//push
	//top
	//pop, if empty
	//check min, max updates

	var st stack

	st.push(1)
	st.push(10)
	fmt.Println(st)
	st.push(120)
	fmt.Println("Min is ", st.getMin())
	fmt.Println("Min is ", st.getMax())
	st.pop()
	fmt.Println("Min is ", st.getMin())
	fmt.Println("Min is ", st.getMax())

}

// 20 API 
// CPU increasing

/*
17.50

consumer x rate msg/sec 
producer y rate msg/sec

	control producing trigger
	producer want to keep on producing
		once channel is full 
			then stop until there is space for new values
		Once get signalled then start producing again
		How i know if it's full?_________

			Rather than just starting to write even if available space is small 
		Signal to start after 'X' space is available______________________channel signal
	
	Consumer starts consuming
		if no message 
			Sleep ________How i know how long to sleep, 
				OR
			Wait for signal from producer to start consuming 
		else{
			process it 
		}
	
*/

func main(){
	 a  := []int{1,2,3,1,55,1}

	 j :=0
	 for i:= 0; i< len(a);i++{
		 if a[i] > j{
			 j = a[i]
		 }
	 }

	 fmt.Println(j)
}

