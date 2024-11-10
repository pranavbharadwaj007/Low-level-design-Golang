package good

type SimplePrinter struct {
    name string
}

func NewSimplePrinter(name string) *SimplePrinter {
    return &SimplePrinter{name: name}
}

func (s *SimplePrinter) Print(doc *Document) {
    println(s.name + ": Printing document...")
}
