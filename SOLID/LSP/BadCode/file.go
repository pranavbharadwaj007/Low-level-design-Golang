package bad

type File struct{}

func (f *File) Read() {
    println("Reading from file...")
}

func (f *File) Write() {
    println("Writing to file...")
}
