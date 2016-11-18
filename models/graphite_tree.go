package models

import (
	"fmt"
	"container/list"
	"strings"
)

type Node struct {
	parent *Node
	children map[string]*Node
}

func (node Node) is_leaf() bool {
	return len(node.children) == 0
}

func (node Node) insert(paths list.List) {
	if len(paths) == 0 {
		return
	}
	child_name := paths[0]
	paths.Remove(child_name)
	if ! node.children[child_name] {
		node.children[child_name] = Node(node)
	}
	node.children[child_name].insert(paths)
}

type MetricTree struct {
	index Node
}

func insert(metric_path string) {
	paths := strings.Split(metric_path, '.')

}

func main() {
	index := MetricTree{nil}
	fmt.Println(index)
}
