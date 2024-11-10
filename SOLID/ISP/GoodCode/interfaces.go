package good

// Separated interfaces for different functionalities
type Printer interface {
    Print(doc *Document)
}

type Scanner interface {
    Scan(doc *Document)
}

type Copier interface {
    Copy(doc *Document)
}

// Optional: Combined interface for devices supporting multiple functions
type MultiFunctionDevice interface {
    Printer
    Scanner
    Copier
}
