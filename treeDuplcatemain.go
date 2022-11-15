/*
resMap(key,val) = inorderTraversal, frequency
create tree
For each node in tree
	Do inOrder traversal & store it

	       3
          / \
         2   6
        /   / \
       9   2   8
          /
         9

*/

package main

import "fmt"

type node struct {
	lft, rt *node
	val     int
}

type tree struct {
	root *node
}

var inorderFreq map[int]int

//inorder for root
//inorder for each node
func (t *node) inorder() int {
	if t == nil {
		return -1
	}
	lft := t.lft.inorder()
	val := t.val
	rt := t.rt.inorder()

	path := 0
	if lft != -1 {
		path = lft
		k, v := inorderFreq[path]
		if v {
			inorderFreq[path] = k + 1
		}else{
			inorderFreq[path] = 1
		}
	}
	path = path*10 + val

	if rt != -1 {
		path = path*10 + rt
		k, v := inorderFreq[rt]
		if v {
			inorderFreq[rt] = k + 1
		}else{
			inorderFreq[rt] = 1
		}
	}

	return path
}

func (t *node) inorderTraversal() {
	traversalPath := t.inorder()
	k, v := inorderFreq[traversalPath]
	if v {
		inorderFreq[traversalPath] = k + 1
	}
}

func (t *node) PrintDuplicateTree() {
	for k, v := range inorderFreq {
		if v > 1 {
			fmt.Println(k,",", v)
		}
	}
}

func main() {
	//create tree
	//insert data
	//inorder for each node
	//search for freq > 1 in map

	var t *node
	t = &node{val: 3, lft: nil, rt: nil}
	t.lft = &node{val: 2, lft: nil, rt: nil}
	t.rt = &node{val: 6, lft: nil, rt: nil}
	t.lft.lft = &node{val: 9, lft: nil, rt: nil}
	t.rt.lft = &node{val: 2, lft: nil, rt: nil}
	t.rt.rt = &node{val: 8, lft: nil, rt: nil}

	t.rt.lft.lft = &node{val: 9, lft: nil, rt: nil}
	inorderFreq = make(map[int]int)
	t.inorderTraversal()
	t.PrintDuplicateTree()
}
