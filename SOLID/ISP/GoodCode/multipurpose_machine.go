package good

type MultiPurposeMachine struct {
    name string
}

func NewMultiPurposeMachine(name string) *MultiPurposeMachine {
    return &MultiPurposeMachine{name: name}
}

func (m *MultiPurposeMachine) Print(doc *Document) {
    println(m.name + ": Printing document...")
}

func (m *MultiPurposeMachine) Scan(doc *Document) {
    println(m.name + ": Scanning document...")
}

func (m *MultiPurposeMachine) Copy(doc *Document) {
    println(m.name + ": Copying document...")
}
