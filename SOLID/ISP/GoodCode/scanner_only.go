package good

type ScannerOnly struct {
    name string
}

func NewScannerOnly(name string) *ScannerOnly {
    return &ScannerOnly{name: name}
}

func (s *ScannerOnly) Scan(doc *Document) {
    println(s.name + ": Scanning document...")
}
