package internal

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name  string
		input []*Node
	}{
		{
			name: "Basic",
			input: []*Node{
				{Ch: 'a', Weight: 5},
				{Ch: 'b', Weight: 9},
				{Ch: 'c', Weight: 12},
				{Ch: 'd', Weight: 13},
				{Ch: 'e', Weight: 16},
				{Ch: 'f', Weight: 45},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			hh := NewHuffmanHeap()
			sum := uint32(0)
			for _, node := range tt.input {
				heap.Push(hh, node)
				sum += node.Weight
			}
			root := hh.BuildTree()

			if root == nil {
				t.Fatal("Expected root node, got nil")
			}

			if root.Weight != sum {
				t.Fatalf("Expected root weight to be 100, got %d", root.Weight)
			}

			// Perform a simple check on the tree structure
			// The exact structure depends on the weights, so we need to check if the sum of weights is correct
			if !isValidHuffmanTree(root) {
				t.Fatal("The Huffman tree structure is invalid")
			}
		})
	}
}

// Helper function to validate Huffman tree structure
func isValidHuffmanTree(node *Node) bool {
	if node.IsLeaf() {
		return true
	}
	leftWeight := uint32(0)
	rightWeight := uint32(0)
	if node.Left != nil {
		leftWeight = node.Left.Weight
		if !isValidHuffmanTree(node.Left) {
			fmt.Println("left is not valid huffman tree", *node)
			return false
		}
	}
	if node.Right != nil {
		rightWeight = node.Right.Weight
		if !isValidHuffmanTree(node.Right) {
			fmt.Println("right is not valid huffman tree", node.Ch)
			return false
		}
	}
	fmt.Println("it doesn't add up", string(node.Ch), node.Weight, node.Left, node.Right)

	return node.Weight == leftWeight+rightWeight
}
