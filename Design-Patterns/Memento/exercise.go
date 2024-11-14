package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Exercise coordinates the shape editing operations
type Exercise struct {
    editor    *GraphicEditor
    caretaker *Caretaker
}

// NewExercise creates a new Exercise instance
func NewExercise() *Exercise {
    return &Exercise{
        editor:    NewGraphicEditor(),
        caretaker: NewCaretaker(),
    }
}

// Run executes the main program logic
func (e *Exercise) Run() {
    scanner := bufio.NewScanner(os.Stdin)

    // Get 3 sets of shape attributes
    for i := 0; i < 3; i++ {
        fmt.Println("Enter shape, x, y, color, size (space-separated):")
        scanner.Scan()
        input := strings.Fields(scanner.Text())

        if len(input) != 5 {
            fmt.Println("Invalid input")
            i--
            continue
        }

        x, _ := strconv.Atoi(input[1])
        y, _ := strconv.Atoi(input[2])
        size, _ := strconv.Atoi(input[4])

        e.editor.SetShape(input[0], x, y, input[3], size)
        e.caretaker.SaveState(e.editor)
    }

    // Perform undo and show result
    e.caretaker.Undo(e.editor)
    fmt.Println("After undo:", e.editor.GetShape())
}