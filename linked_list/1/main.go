package main

import (
	"fmt"
)

type Node struct {
	val int
	next *Node
}

type LinkedList struct {
	head *Node
}


func (l *LinkedList) add(v int) {
	newNode := &Node{val: v}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head 
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (l LinkedList) String() string {
	var result string
	curr := l.head

	for curr != nil {
		result += fmt.Sprintf("%d->", curr.val)
		curr = curr.next
	}
	result += "\n"

	return result
}

func (l *LinkedList) RemoveDuplicates() {
	hashMap := make(map[int]bool)

	curr := l.head

	for curr.next != nil {
		if hashMap[curr.next.val]{
			curr.next = curr.next.next
		} else {
			hashMap[curr.val] = true
			curr = curr.next
		}
	}
}

func main(){
	l := LinkedList{}
	l.add(4)
	l.add(6)
	l.add(7)
	l.add(9)
	l.add(5)
	l.add(4)
	l.add(8)
	l.add(7)
	
	fmt.Println(l)
	l.RemoveDuplicates()
	fmt.Println(l)
}
