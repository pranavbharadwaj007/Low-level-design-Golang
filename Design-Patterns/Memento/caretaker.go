package main

// Caretaker manages the history of editor states
type Caretaker struct {
	history []*EditorMemento
}

// NewCaretaker creates a new Caretaker
func NewCaretaker() *Caretaker {
	return &Caretaker{
		history: make([]*EditorMemento, 0),
	}
}

// SaveState adds current editor state to history
func (c *Caretaker) SaveState(editor *GraphicEditor) {
	c.history = append(c.history, editor.Save())
}

// Undo reverts to the previous state if available
func (c *Caretaker) Undo(editor *GraphicEditor) {
	if len(c.history) < 2 {
		return
	}
	// Remove current state
	c.history = c.history[:len(c.history)-1]
	// Get previous state
	previousState := c.history[len(c.history)-1]
	editor.Restore(previousState)
}
