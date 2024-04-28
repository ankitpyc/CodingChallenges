package filereader

type FileDetails struct {
	WordCount   int
	LineCount   int
	CharCount   int
	MaxFrequent []MaxFrequentWord
}

type MaxFrequentWord struct {
	Word  string
	Count int
}
