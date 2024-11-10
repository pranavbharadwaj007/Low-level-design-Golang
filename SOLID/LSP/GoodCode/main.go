package main

import "files"

// ReadAnyFile can take any file that implements Readable
func readAnyFile(file files.Readable) {
    file.Read()
}

func main() {
    // Create instances
    readOnlyFile := files.NewReadOnlyFile()
    writableFile := files.NewWritableFile()

    // Use the files
    readOnlyFile.Read()
    
    writableFile.Read()
    writableFile.Write()

    // Both can be used where Readable is expected
    readAnyFile(readOnlyFile)
    readAnyFile(writableFile)
}
