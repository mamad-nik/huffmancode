package builder

// this package provides tools to build and work with a tree

import (
	"fmt"

	huffman "github.com/mamad-nik/huffman"
)

func swap(slice []huffman.Nodetype, s int) []huffman.Nodetype {
	if s == len(slice)-1 {
		return append(slice[:s-1], slice[s], slice[s-1])

	}
	temp := append(slice[:s], slice[s+1], slice[s])
	return append(temp, slice[s+2:]...)
	// OK
}
func remove(pq huffman.Pqueue, index int) (huffman.Pqueue, *huffman.Nodetype) {
	tmp := pq.Nodes[index]
	pq.Nodes = append(pq.Nodes[:index], pq.Nodes[index+1:]...)
	pq.Cap--
	return pq, &tmp
	//OK
}

func sorter(slice []huffman.Nodetype) []huffman.Nodetype {
	n := len(slice)
	for i := 1; i < n; i++ {
		key := slice[i]
		j := i - 1
		for j >= 0 && (slice[j].Frequency > key.Frequency) {
			swap(slice, j)
			j--
		}
		slice[j+1] = key
	}
	return slice
	//OK
}

func insert(pq huffman.Pqueue, node ...huffman.Nodetype) huffman.Pqueue {
	for _, v := range node {
		temp := append(pq.Nodes, v)
		pq = huffman.Pqueue{Nodes: sorter(temp), Cap: pq.Cap + 1}
	}
	return pq
	//OK
}
func printtree(node huffman.Nodetype) {
	fmt.Println(node.Symbole, node.Frequency)
	if node.Left != nil {
		printtree(*node.Left)
	}
	if node.Right != nil {
		printtree(*node.Right)
	}
}
func buildtree(pq huffman.Pqueue) *huffman.Nodetype {
	for pq.Cap > 1 {
		var p, q *huffman.Nodetype
		pq.Nodes = sorter(pq.Nodes)

		pq, p = remove(pq, 0)
		pq, q = remove(pq, 0)

		r := huffman.Nodetype{
			Symbole:   p.Symbole + q.Symbole,
			Frequency: p.Frequency + q.Frequency,
			Left:      p,
			Right:     q,
		}
		pq = insert(pq, r)
	}
	var r *huffman.Nodetype
	pq, r = remove(pq, 0)
	return r
	//OK
}
func Builder(pq *huffman.Pqueue, nodes ...huffman.Nodetype) (tree *huffman.Nodetype) {
	*pq = insert(*pq, nodes...)
	tree = buildtree(*pq)
	printtree(*tree)
	return tree
}
