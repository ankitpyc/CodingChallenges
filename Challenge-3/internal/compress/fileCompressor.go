package compress

import (
	huffmanTree "challenge3/huffman"
	analyzer "challenge3/internal/TextAnalyzer"
	minheap "challenge3/internal/Tree"
	"fmt"
	"log"
	"os"
	"path"
)

type FileDetails CompressedFile

type FileCompressor struct{}

func (f *FileCompressor) CompressFile(filepath string) map[rune]string {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("error while opening file ", err)
	}
	defer file.Close()
	var freqCount map[rune]int = analyzer.AnalyzeEncode(file)
	heap := minheap.NewTreeWithCapacity(len(freqCount))
	heap.BuildTreeWithFreqCount(freqCount)
	var huff *huffmanTree.HuffmanTree = huffmanTree.NewHuffManTree(heap)
	huff.BuildHuffManTree().BuildEncodings()
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
