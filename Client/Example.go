package main

type node struct {
	value      int
	leftChild  *node
	rightChild *node
}

type tree struct {
	root node
}

func build(value int, nod *node) {
	n := nod
	if value > n.value && n.rightChild.value != 0 {
		build(value, n.rightChild)
	}
	if value > n.value && n.rightChild.value == 0 {
		n.rightChild = &node{value, new(node), new(node)}
	}
	if value < n.value && n.leftChild.value != 0 {
		build(value, n.leftChild)
	}
	if value < n.value && n.leftChild.value == 0 {
		n.leftChild = &node{value, new(node), new(node)}
	}
}

func main() {
	graf := tree{root: node{3, new(node), new(node)}}
	array := []int{2, 4, 6, 1, 5}
	for _, v := range array {
		build(v, &graf.root)
	}
}
