package main

// EditorMemento stores the state of a shape
type EditorMemento struct {
	shapeType string
	x, y      int
	color     string
	size      int
}

// NewEditorMemento creates a new memento with the given state
func NewEditorMemento(shapeType string, x, y int, color string, size int) *EditorMemento {
	return &EditorMemento{
		shapeType: shapeType,
		x:         x,
		y:         y,
		color:     color,
		size:      size,
	}
}

// GetState returns all state values
func (m *EditorMemento) GetState() (string, int, int, string, int) {
	return m.shapeType, m.x, m.y, m.color, m.size
}
