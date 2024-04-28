package main

import (
	freader "Challenge-1/internal/filereader"
	file_details "Challenge-1/internal/filereader/dto"

	"fmt"

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
			filedetails := &file_details.FileDetails{}
			freader.ReadFile(args[0], filedetails)
			res := fmt.Sprintf("LC : %d , WC : %d , CC : %d", filedetails.LineCount, filedetails.WordCount, filedetails.CharCount)
			fmt.Println(res)
		},
	}

	fileReaderCmd.Flags().StringVarP(&filename, "filepath", "f", "hiname.txt", "name of the file to be read")
	rootCmd.MarkFlagRequired("filepath")
	rootCmd.AddCommand(fileReaderCmd)
	rootCmd.Execute()
}
