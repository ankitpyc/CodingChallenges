package main

import (
	fileCompressor "challenge3/internal/compress"

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
			compressed_map := fileCompressor.CompressFile(args[0])
			fileCompressor.WriteEncodedFile(args[0], compressed_map)

		},
	}
	fileReaderCmd.Flags().StringVarP(&filename, "filepath", "f", "hiname.txt", "name of the file to be read")
	rootCmd.MarkFlagRequired("filepath")
	rootCmd.AddCommand(fileReaderCmd)
	// Assign values to the array elements
	rootCmd.Execute()
}
