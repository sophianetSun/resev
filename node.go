package main

import "fmt"

/*
	Node is element for reservation building
*/
type Node struct {
	TimeID uint // Time for ID
	Prev   *Node
	Next   *Node
	Elems  []string
}

func (o *Node) LinkLeft(n *Node) {
	o.Prev = n
	n.Next = o
}

func (o *Node) LinkRight(n *Node) {
	o.Next = n
	n.Prev = o
}

func (o *Node) UnlinkLeft() {
	o.Prev.Next = nil
	o.Prev = nil
}

func (o *Node) UnlinkRight() {
	o.Next.Prev = nil
	o.Next = nil
}

func Foo() {
	node := Node{1, nil, nil, []string{"a", "b"}}
	fmt.Println(node)
	newNode := Node{2, nil, nil, []string{"foo"}}
	node.LinkLeft(&newNode)
	fmt.Println(node, newNode)
	rightNode := Node{3, nil, nil, []string{"bar", "baz"}}
	node.LinkRight(&rightNode)
	fmt.Println(node, newNode, rightNode)
	node.UnlinkLeft()
	fmt.Println(node, newNode, rightNode)
}
