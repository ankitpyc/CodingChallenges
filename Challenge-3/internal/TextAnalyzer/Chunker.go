package analyzer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type TextAnalyzer interface {
	Analyze(file *os.File)
	CreateCunks(words []string, freqMap map[rune]int)
}

func AnalyzeEncode(file *os.File) map[rune]int {
	var freqMap map[rune]int = make(map[rune]int)
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		fmt.Println(line)
		if err != nil && err == io.EOF {
			break
		}
		splitWords := strings.Split(line, " ")
		CreateCunks(splitWords, freqMap)
	}
	return freqMap
}

func CreateCunks(words []string, freqMap map[rune]int) {
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
