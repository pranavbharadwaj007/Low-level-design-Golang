package main

import "myapp/bad"

func main() {
    var file *bad.File = &bad.ReadOnlyFile{}
    file.Read()  // works fine
    file.Write() // will panic - violates LSP
}
