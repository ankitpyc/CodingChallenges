package compress

import (
	"bufio"
	huffmanTree "challenge3/huffman"
	minheap "challenge3/internal/Tree"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type FileDetails CompressedFile

type FileCompressor struct{}

func (f *FileCompressor) CompressFile(filepath string) map[rune]int {
	// var fileDetails = &FileDetails{}
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("error reading file at the path ", err)
	}
	var freqMap map[rune]int = make(map[rune]int)
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read file line error: %v", err)
		}
		splitWords := strings.Split(line, " ")
		f.ProcessWordinLine(splitWords, freqMap)
	}

	heap := minheap.NewTreeWithCapacity(len(freqMap))
	for key, val := range freqMap {
		heap.AddNode(key, val)
	}
	var huff *huffmanTree.HuffmanTree = huffmanTree.BuildHuffManTree(heap)

	node := huff.Encode()
	traverseTree(node, "")
	fmt.Print("node is ", node.Count)
	fmt.Println()
	size := heap.Size()
	for j := 0; j < size; j++ {
		no := heap.ExtractMin()
		fmt.Printf("min  %c -- %d", no.Char, no.Count)
		fmt.Println()
	}

	return freqMap
}

func traverseTree(node *minheap.CharNode, str string) {
	if node != nil {
		traverseTree(node.Left, str+"1")
		if node.Char != '-' {
			fmt.Printf("%c  : %s ", node.Char, str)
			fmt.Println()
		}
		fmt.Println()
		traverseTree(node.Right, str+"0")
	}
}

func (f *FileCompressor) ProcessWordinLine(words []string, freqMap map[rune]int) {
	for _, word := range words {
		trimmedword := strings.TrimSpace(word)
		var chars []rune = []rune(trimmedword)
		for _, val := range chars {
			_, ok := freqMap[val]
			if ok {
				freqMap[val]++
			} else {
				freqMap[val] = 1
			}
		}
	}
}
