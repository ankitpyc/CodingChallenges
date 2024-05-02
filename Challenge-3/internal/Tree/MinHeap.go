package tree

import (
	treeNode "challenge3/internal/Tree/TreeNode"
	"fmt"
)

type Node treeNode.TreeNode

type MinHeap struct {
	heapsize int
	root     *Node
	arr      [](*CharNode)
	capacity int
}

func NewTree() *MinHeap {
	return &MinHeap{
		heapsize: 0,
		root:     nil,
	}
}

func NewTreeWithCapacity(capacity int) *MinHeap {
	return &MinHeap{
		heapsize: 0,
		root:     nil,
		arr:      make([]*CharNode, capacity),
	}
}
func (heap *MinHeap) Size() int {
	return heap.heapsize
}

func (heap *MinHeap) AddNode(char rune, count int) {
	fmt.Printf("adding node for %c", char)
	fmt.Println()
	heap.heapsize++
	i := heap.heapsize - 1
	heap.arr[i] = &CharNode{
		Char:  char,
		Count: count,
		Left:  nil,
		Right: nil,
	}
	for i != 0 && heap.arr[heap.parent(i)].Count > heap.arr[i].Count {
		heap.swap(heap.parent(i), i)
		i = heap.parent(i)
	}
}

func (heap *MinHeap) AddNodes(node *CharNode) {
	fmt.Printf("adding node for %c", node.Char)
	fmt.Println()
	heap.heapsize++
	i := heap.heapsize - 1
	heap.arr[i] = node
	for i != 0 && heap.arr[heap.parent(i)].Count > heap.arr[i].Count {
		heap.swap(heap.parent(i), i)
		i = heap.parent(i)
	}
}

func (heap *MinHeap) GetMinHeap() []*CharNode {
	return heap.arr
}

func (heap *MinHeap) ExtractMin() CharNode {
	minNode := heap.arr[0]
	heap.arr[0] = heap.arr[heap.heapsize-1]
	heap.heapsize--
	heap.MinHeapify(0)
	return *minNode
}

func (heap *MinHeap) MinHeapify(i int) {
	l := heap.left(i)
	r := heap.right(i)
	var smallest int = i
	if l < heap.heapsize && heap.arr[l].Count < heap.arr[i].Count {
		smallest = l
	}

	if r < heap.heapsize && heap.arr[r].Count < heap.arr[smallest].Count {
		smallest = r
	}
	if smallest != i {
		heap.swap(i, smallest)
		heap.MinHeapify(smallest)
	}
}
func (heap *MinHeap) getMin() *CharNode { return heap.arr[0] }

func (heap *MinHeap) parent(i int) int { return (i - 1) / 2 }

// to get index of left child of node at index i
func (heap *MinHeap) left(i int) int { return (2*i + 1) }

func (heap *MinHeap) swap(i, j int) {
	temp := heap.arr[i]
	heap.arr[i] = heap.arr[j]
	heap.arr[j] = temp
}

// to get index of right child of node at index i
func (heap *MinHeap) right(i int) int { return (2*i + 2) }
