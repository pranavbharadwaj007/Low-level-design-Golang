package bad

type ReadOnlyFile struct {
    File
}

func (r *ReadOnlyFile) Write() {
    panic("Can't write to a read only file")
}
