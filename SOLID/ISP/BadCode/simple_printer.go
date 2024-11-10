package bad

// SimplePrinter is forced to implement methods it doesn't support
type SimplePrinter struct {
    name string
}

func NewSimplePrinter(name string) *SimplePrinter {
    return &SimplePrinter{name: name}
}

func (s *SimplePrinter) Print(doc *Document) {
    println(s.name + ": Printing document...")
}

func (s *SimplePrinter) Scan(doc *Document) {
    panic("Scan operation not supported")
}

func (s *SimplePrinter) Copy(doc *Document) {
    panic("Copy operation not supported")
}
