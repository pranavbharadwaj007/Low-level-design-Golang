package files

type ReadOnlyFile struct {
    ReadableFile
}

func NewReadOnlyFile() *ReadOnlyFile {
    return &ReadOnlyFile{}
}
