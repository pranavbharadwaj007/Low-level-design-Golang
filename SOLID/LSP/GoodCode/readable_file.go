package files

type ReadableFile struct{}

func NewReadableFile() *ReadableFile {
    return &ReadableFile{}
}

func (r *ReadableFile) Read() {
    println("Reading from a file")
}
