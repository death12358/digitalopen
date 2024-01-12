package tree

// AddNode 新增節點
func (n *Node) AddNode(value int) {
	child := &Node{
		Site:     value,
		Children: make([]*Node, 0),
	}
	n.Children = append(n.Children, child)
}

// ReplaceNode 替換節點值
func (n *Node) ReplaceNode(site int, newValue string) {
	if n.Site == site {
		n.Value = newValue
	}

	for _, child := range n.Children {
		child.ReplaceNode(site, newValue)
	}
}

func CheckWild(line []string, index int) string {
	if index == 0 {
		return AllWild
	}
	if line[index-1] != Wild {
		return line[index-1]
	}

	return CheckWild(line, index-1)
}
