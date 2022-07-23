package main

import (
	"testing"
)

func TestLinkLeft(t *testing.T) {
	node := Node{1, nil, nil, []string{}}
	lNode := Node{2, nil, nil, []string{}}
	node.LinkLeft(&lNode)
	if node.Prev != &lNode && lNode.Next != &node {
		t.Error("lNode is node's prev")
	}
}
