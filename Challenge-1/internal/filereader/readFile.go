package filereader

import (
	f_reader "Challenge-1/internal/filereader/dto"
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFile(filepath string, fileDetails *f_reader.FileDetails) {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("read file line error: %v", err)
	}
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return
		}
		splitWords := strings.Split(line, " ")
		for _, word := range splitWords {
			fileDetails.CharCount = fileDetails.CharCount + len(word)
		}
		fileDetails.LineCount = fileDetails.LineCount + 1
		fileDetails.WordCount = fileDetails.WordCount + len(splitWords)
	}
}
