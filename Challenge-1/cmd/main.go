package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var filename string
	var rootCmd = &cobra.Command{Use: "app"}
	var fileReaderCmd = &cobra.Command{
		Use:   "file [filename.txt]",
		Short: "prints the no of lines,words , characters to the terminal",
		Long:  `prints the no of lines,words , characters to the terminal`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				lineCount int64
				wordCount int
				charCount int
			)

			f, err := os.OpenFile(args[0], os.O_RDWR|os.O_CREATE, 0644)
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
					charCount = charCount + len(word)
				}
				lineCount = lineCount + 1
				wordCount = wordCount + len(splitWords)
			}
			res := fmt.Sprintf("LC : %d , WC : %d , CC : %d", lineCount, wordCount, charCount)
			fmt.Println(res)
		},
	}

	fileReaderCmd.Flags().StringVarP(&filename, "filepath", "f", "hiname.txt", "name of the file to be read")
	rootCmd.MarkFlagRequired("filepath")
	rootCmd.AddCommand(fileReaderCmd)
	rootCmd.Execute()
}
