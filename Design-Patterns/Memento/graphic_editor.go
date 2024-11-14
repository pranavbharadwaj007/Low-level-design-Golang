package main

import "fmt"

// GraphicEditor manages the properties of shapes
type GraphicEditor struct {
    shapeType string
    x, y      int
    color     string
    size      int
}

// NewGraphicEditor creates a new GraphicEditor
func NewGraphicEditor() *GraphicEditor {
    return &GraphicEditor{}
}

// Save creates and returns a new memento of the current state
func (e *GraphicEditor) Save() *EditorMemento {
    return NewEditorMemento(e.shapeType, e.x, e.y, e.color, e.size)
}

// Restore updates the editor's state from a memento
func (e *GraphicEditor) Restore(m *EditorMemento) {
    e.shapeType, e.x, e.y, e.color, e.size = m.GetState()
}

// SetShape updates all shape attributes
func (e *GraphicEditor) SetShape(shapeType string, x, y int, color string, size int) {
    e.shapeType = shapeType
    e.x = x
    e.y = y
    e.color = color
    e.size = size
}

// GetShape returns a string representation of the current shape
func (e *GraphicEditor) GetShape() string {
    return fmt.Sprintf("Shape: %s, Position: (%d, %d), Color: %s, Size: %d",
        e.shapeType, e.x, e.y, e.color, e.size)
}