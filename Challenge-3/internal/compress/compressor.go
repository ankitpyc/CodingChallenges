package compress

type CompressedFile struct {
	OriginalFileUrl string
	NewFileUrl      string
	OriginalSize    uint64
	CompressedSize  uint64
}
