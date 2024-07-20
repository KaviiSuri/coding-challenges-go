package internal

import "fmt"

type Node struct {
	Ch     byte
	Weight uint32
	Right  *Node
	Left   *Node
}

func newNode(
	ch byte,
	weight uint32,
) Node {
	return Node{
		Ch:     ch,
		Weight: weight,
	}
}

func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

type Code map[byte][]bool

func (n *Node) BuildCode() Code {
	code := make(Code)
	// Special case: If the tree has only one node
	if n.Left == nil && n.Right == nil {
		code[n.Ch] = []bool{false} // Assigning a default code for single character
		return code
	}
	buildCode(n, []bool{}, code)
	return code
}

func buildCode(n *Node, currentCode []bool, c Code) {
	if n.IsLeaf() && n.Ch != 0 {
		c[n.Ch] = append([]bool{}, currentCode...)
	}

	if n.Left != nil {
		buildCode(n.Left, append(currentCode, false), c)
	}
	if n.Right != nil {
		buildCode(n.Right, append(currentCode, true), c)
	}
}

func PrettyPrentCode(c Code) string {
	s := "\n"
	for k, bs := range c {
		v := ""
		for _, b := range bs {
			if b {
				v += "1"
			} else {
				v += "0"
			}
		}
		s += fmt.Sprintf("%c: %v\n", k, v)
	}
	return s
}
