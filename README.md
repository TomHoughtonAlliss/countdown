# countdown
A simple app to generate countdown number puzzles

# Usage
Set your puzzle input in `puzzleInput` inside `main.go`.
 - `L` - a large number, one of [25, 50, 75, 100]. A maximum of four of these can be used.
 - `S` - a smaller number, from the interval [1, 10].

Run the app with `go run main.go`. It should output a puzzle instance.

Pressing Enter refreshes the puzzle input, with the same large-small structure.
