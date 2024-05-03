package compress

import (
	"bufio"
	huffmanTree "challenge3/huffman"
	minheap "challenge3/internal/Tree"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

type FileDetails CompressedFile

type FileCompressor struct{}

func (f *FileCompressor) CompressFile(filepath string) map[rune]string {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("error while opening file ", err)
	}
	defer file.Close()
	if err != nil {
		log.Fatal("error reading file at the path ", err)
	}
	var freqMap map[rune]int = make(map[rune]int)
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		fmt.Println(line)
		if err != nil && err == io.EOF {
			break
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
	traverseTree(huff, node, "")
	return huff.Charmap
}

func (f *FileCompressor) WriteEncodedFile(filepath string, compressed_map map[rune]string) {
	file, _ := os.OpenFile(filepath, os.O_RDONLY, 0644)
	cwd, _ := os.Getwd()
	fmt.Println(cwd)
	outpath := path.Join(cwd, "output.txt")
	ofile, err := os.OpenFile(outpath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer ofile.Close()
	buf := make([]byte, 1)
	for {
		// Read one byte (one character) from the file
		n, err := file.Read(buf)

		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		if n == 0 {
			// Reached end of file
			break
		}

		// Process the character (in this example, just print it)
		chars := string(buf[0])
		val, ok := compressed_map[rune(buf[0])]
		if ok {
			_, err := ofile.WriteString(val)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		} else {
			ofile.WriteString(chars)
		}
	}
}

func (f *FileCompressor) BuildHuffManTree(heap *minheap.MinHeap) {
	var huff *huffmanTree.HuffmanTree = huffmanTree.BuildHuffManTree(heap)
	node := huff.Encode()
	traverseTree(huff, node, "")
}

func traverseTree(huff *huffmanTree.HuffmanTree, node *minheap.CharNode, str string) {
	if node != nil {
		traverseTree(huff, node.Left, str+"1")
		if node.Char != '-' {
			huff.Charmap[node.Char] = str
		}
		traverseTree(huff, node.Right, str+"0")
	}
}

func (f *FileCompressor) ProcessWordinLine(words []string, freqMap map[rune]int) {
	spaces_length := len(words) - 1
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
	_, ok := freqMap[' ']
	if ok {
		freqMap[' '] = freqMap[' '] + spaces_length
	} else {
		freqMap[' '] = spaces_length
	}

	_, ok = freqMap['\n']
	if ok {
		freqMap['\n'] = freqMap['\n'] + 1
	} else {
		freqMap['\n'] = 1
	}

}
