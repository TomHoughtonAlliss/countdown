## countdown
A simple app to generate countdown number puzzles, following the standard rules of show.

## Usage
### Configuration
Set your puzzle input in `puzzleInput` inside `main.go`.
 - `L` - a large number, one of [25, 50, 75, 100]. A maximum of four of these can be used.
 - `S` - a smaller number, from the interval [1, 10].

### Running
Run the app with `go run main.go`. It should output a puzzle instance.

Pressing Enter refreshes the puzzle input, with the same large-small structure.

## Plans
### Implement a checker
A user can input their solution and have it checked by the app. Most likely with an output similar to:
```
Target: 379. 
Available: 75, 5, 10, 1, 5, 10

> 75 * 5

Target: 379
Available: 375, 10, 1, 5, 10

> 5 - 1

Target: 379
Available: 375, 10, 10, 4

> 375 + 4

Target Reached!
```