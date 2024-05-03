package algorithim

import (
	heap "challenge3/internal/Tree"
	"fmt"
)

type HuffmanTree struct {
	root    *heap.CharNode
	heap    *heap.MinHeap
	Charmap map[rune]string
}

func BuildHuffManTree(heap *heap.MinHeap) *HuffmanTree {
	return &HuffmanTree{
		root:    nil,
		heap:    heap,
		Charmap: make(map[rune]string),
	}
}

func (huffman *HuffmanTree) Encode() *heap.CharNode {
	for huffman.heap.Size() > 1 {
		node1 := huffman.heap.ExtractMin()
		node2 := huffman.heap.ExtractMin()
		MergeNode(huffman, node1, node2)
	}
	huffman.root = huffman.heap.ExtractMin()
	return huffman.root
}

func MergeNode(huffman *HuffmanTree, node1, node2 *heap.CharNode) {
	fmt.Printf("node 1 %c : %d node 2  %c : %d", node1.Char, node1.Count, node2.Char, node2.Count)
	fmt.Println()
	node3 := &heap.CharNode{
		Count: node1.Count + node2.Count,
		Char:  '-',
		Left:  node1,
		Right: node2,
	}

	huffman.heap.AddNodes(node3)
}
