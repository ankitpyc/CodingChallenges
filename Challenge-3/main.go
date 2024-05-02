package main

import (
	fileCompressor "challenge3/internal/compress"
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var filename string
	var fileCompressor fileCompressor.FileCompressor
	var rootCmd = &cobra.Command{Use: "app"}
	var fileReaderCmd = &cobra.Command{
		Use:   "file [filename.txt]",
		Short: "Compresses the filename using Huffman coding Algorithim",
		Long:  `prints the no of lines,words , characters to the terminal`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			stringmap := fileCompressor.CompressFile(args[0])
			for key, value := range stringmap {
				fmt.Printf("Key: %c, Value: %d\n", key, value)
				fmt.Println()
			}
		},
	}
	fileReaderCmd.Flags().StringVarP(&filename, "filepath", "f", "hiname.txt", "name of the file to be read")
	rootCmd.MarkFlagRequired("filepath")
	rootCmd.AddCommand(fileReaderCmd)
	// Assign values to the array elements
	rootCmd.Execute()
}
