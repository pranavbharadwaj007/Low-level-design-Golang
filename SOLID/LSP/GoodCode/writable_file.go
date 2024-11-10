package files

type WritableFile struct {
    ReadableFile
}

func NewWritableFile() *WritableFile {
    return &WritableFile{}
}

func (w *WritableFile) Write() {
    println("Writing to a file...")
}
