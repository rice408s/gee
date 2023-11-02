package gin

import "strings"

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查询
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)

	for i, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, n.children[i])
		}
	}
	return nodes
}

// recursive insert node
func (n *node) insert(pattern string, parts []string, height int) {
	// 递归终止条件
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)

	//insert node when child is nil
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	//recursion
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	// when search to the end of parts or part start with * ,return
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		//recursion search
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
