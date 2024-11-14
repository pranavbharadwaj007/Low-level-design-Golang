# Understanding the Memento Pattern

## Pattern Structure
```mermaid
classDiagram
    note for GraphicEditor "ORIGINATOR"
    class GraphicEditor {
        -shapeType string
        -x int
        -y int
        -color string
        -size int
        +save() EditorMemento
        +restore(EditorMemento)
        +setShape(type, x, y, color, size)
        +getShape() string
    }
    
    note for EditorMemento "MEMENTO"
    class EditorMemento {
        -shapeType string
        -x int
        -y int
        -color string
        -size int
        +getState()
    }
    
    note for Caretaker "CARETAKER"
    class Caretaker {
        -history Stack~EditorMemento~
        +saveState(GraphicEditor)
        +undo(GraphicEditor)
    }
    
    GraphicEditor ..> EditorMemento : creates >
    Caretaker o--> EditorMemento : manages
```

## Sequence of Operations
```mermaid
sequenceDiagram
    participant C as Client
    participant G as GraphicEditor
    participant M as EditorMemento
    participant T as Caretaker
    
    Note over C,T: Save Operation
    C->>G: setShape(circle, 10, 20, red, 5)
    G->>G: Updates state
    C->>T: saveState(editor)
    T->>G: save()
    G->>M: create(state)
    M-->>G: memento
    G-->>T: memento
    T->>T: Add to history
    
    Note over C,T: Restore Operation
    C->>T: undo(editor)
    T->>T: Pop current state
    T->>T: Get previous state
    T->>G: restore(previousState)
    G->>M: getState()
    M-->>G: state values
    G->>G: Update state
```

## Components and Responsibilities

### 1. Originator (GraphicEditor)
- Contains the state that needs to be saved
- Creates mementos containing a snapshot of its current state
- Uses mementos to restore its state

### 2. Memento (EditorMemento)
- Stores the state of the Originator
- Provides immutable access to stored state
- Only Originator can access full state

### 3. Caretaker
- Keeps track of multiple mementos
- Never modifies mementos
- Manages when to save and restore Originator's state

## Example Usage

```go
    // Create the editor and caretaker
    editor := NewGraphicEditor()
    caretaker := NewCaretaker()

    // Example 1: Creating a circle
    fmt.Println("\n=== Creating Circle ===")
    editor.SetShape("circle", 10, 20, "red", 5)
    fmt.Println("Current State:", editor.GetShape())
    caretaker.SaveState(editor)
