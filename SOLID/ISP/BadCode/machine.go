package bad

// Machine - Fat interface that violates ISP
type Machine interface {
    Print(doc *Document)
    Scan(doc *Document)
    Copy(doc *Document)
}
