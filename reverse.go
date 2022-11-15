package main

import "fmt"

type node struct {
	val        int
	prev, next *node
}

type ll struct {
	head, tail *node
}



func (l *ll) insertToList(val int) {
	if l.head == nil {
		l.head = &node{val: val, prev: nil, next: nil}
		l.tail = l.head
	} else {
		v := l.tail
		v.next = &node{val: val, prev: l.tail, next: nil}
		l.tail = v.next
	}
}


func (l *ll) reverseList() {

	var h, tail,tmp  *node
	h = nil
	tail = nil 
	tmp = nil 
 // 1 -> 2 -> 3-> 4->5 

	for i := l.head; i != nil; {
		if h == nil {
			h = i 
			tail = h 
			tmp  = i.next
		}else{			
			h.prev = i 
			tmp  = i.next
			i.next = h
			h = i 
		}

		i = tmp 
	}
	tail.next = nil 
	l.head = h 
	l.tail = tail
}

func ListDriver() {
	var l ll
	l.head = nil
	l.tail = nil

	l.insertToList(1)
	l.insertToList(2)
	l.insertToList(3)
	l.insertToList(4)

	for i := l.head; i != nil; i = i.next {
		fmt.Println(i.val)
	}

	l.reverseList()
	fmt.Println("'''''''''")
	for i := l.head; i != nil; i = i.next {
		fmt.Println(i.val)
	}
}

var tree *node 
var inOrderData = []int {}

func (t *node) insert(v int) *node{
	if t == nil{
		t =  &node{val:v, prev: nil , next: nil }
	}else{
		if v > t.val{
			t.next = t.next.insert(v)
			
		}else{
			t.prev = t.prev.insert(v)
			
		}
	}
	return t
	
}

func (t *node) inorder(){
	
	if t == nil{
		return 
	}
	
	t.prev.inorder()	
	inOrderData = append(inOrderData, t.val)
	t.next.inorder()
}

func main(){
	tree := tree.insert(5)
	tree.insert(3)
	tree.insert(4)
	tree.insert(6)
	tree.insert(7)
	
	
	fmt.Println(tree)
	tree.inorder()
	fmt.Println(inOrderData)
}