package goow

import "strings"

type PatRoot struct {
	root *node
}

func NewPatRoot() *PatRoot {
	return &PatRoot{
		root: newNode(),
	}
}

func (p *PatRoot) insert(pattern string, parts []string) {
	p.root.insert(pattern, parts, 0)
}

func (p *PatRoot) search(parts []string) *node {
	return p.root.search(parts, 0)
}

type node struct {
	//待匹配路由
	pattern string
	// important fields
	part     string
	children []*node
	// node color (isKey)
	isWildcard bool
}

func newNode() *node {
	return &node{}
}

// matchChild is match a node.part in pattern
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWildcard {
			return child
		}
	}
	return nil
}

// matchChildren is match all nodes.part in pattern
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWildcard {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	//
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// current part
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		// create URL node
		child = &node{part: part, isWildcard: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// 下推
	child.insert(pattern, parts, height+1)

}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	// one by one
	part := parts[height]
	children := n.matchChildren(part)

	// 下推
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
