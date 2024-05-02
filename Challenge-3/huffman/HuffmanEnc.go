package algorithim

import (
	heap "challenge3/internal/Tree"
	treeNode "challenge3/internal/Tree/TreeNode"
	"fmt"
)

type HuffmanTree struct {
	root *treeNode.TreeNode
	heap *heap.MinHeap
}

func BuildHuffManTree(heap *heap.MinHeap) *HuffmanTree {
	return &HuffmanTree{
		root: nil,
		heap: heap,
	}
}

func (huffman *HuffmanTree) Encode() *heap.CharNode {
	for huffman.heap.Size() > 1 {
		node1 := huffman.heap.ExtractMin()
		node2 := huffman.heap.ExtractMin()
		MergeNode(huffman, &node1, &node2)
	}
	minNode := huffman.heap.ExtractMin()
	return &minNode
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
