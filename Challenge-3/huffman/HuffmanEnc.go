package algorithim

import (
	heap "challenge3/internal/Tree"
)

type HuffmanTree struct {
	root    *heap.CharNode
	heap    *heap.MinHeap
	Charmap map[rune]string
}

func NewHuffManTree(heap *heap.MinHeap) *HuffmanTree {
	return &HuffmanTree{
		root:    nil,
		heap:    heap,
		Charmap: make(map[rune]string),
	}
}

func (huffman *HuffmanTree) BuildHuffManTree() *HuffmanTree {
	for huffman.heap.Size() > 1 {
		node1 := huffman.heap.ExtractMin()
		node2 := huffman.heap.ExtractMin()
		huffman.MergeNode(node1, node2)
	}
	huffman.root = huffman.heap.ExtractMin()
	return huffman
}

func (huffman *HuffmanTree) MergeNode(node1, node2 *heap.CharNode) {
	node3 := &heap.CharNode{
		Count: node1.Count + node2.Count,
		Char:  '-',
		Left:  node1,
		Right: node2,
	}

	huffman.heap.AddNodes(node3)
}

func (huffman *HuffmanTree) BuildEncodings() map[rune]string {
	huffman.traverseTree(huffman.root, "")
	return huffman.Charmap
}

func (huffman *HuffmanTree) traverseTree(node *heap.CharNode, str string) {
	if node != nil {
		huffman.traverseTree(node.Left, str+"1")
		if node.Char != '-' {
			huffman.Charmap[node.Char] = str
		}
		huffman.traverseTree(node.Right, str+"0")
	}
}
