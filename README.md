## countdown
A simple app to generate countdown number puzzles and check a user's process.

## Usage
### Configuration
Set your puzzle input in `puzzleInput` inside `main.go`. The length of this is recommended to be 6.
 - `L` - a large number, one of [25, 50, 75, 100]. A maximum of four of these can be used.
 - `S` - a smaller number, from the interval [1, 10].

Here you can also change the target interval, but without carefully considering the puzzle input this will quickly become impossible.

### Running
Run the app with `go run main.go`. It should output a puzzle instance based on your configuration.

Then, an input `>` appears indicating that you can enter your first equation.

Here you can input some commands or an equation. The app checks for commands first.
 - `reset` - this returns the puzzle to its initial configuration, as a restart.
 - `next` - this skips the current puzzle.

Otherwise, the equation must be of the form `x o y`, where `x` and `y` are integers and `o` is an operator from this list:
 - `*` - multiplication.
 - `/` - division.
 - `+` - addition.
 - `-` - subtraction.

 **If your input deviates from this structure the app will error and need to be re-run.**

The result of your equation will be calculated, the numbers used removed, and your result added to your usable values.

If your equation reaches the target value then it is considered solved. You get your full equation outputted, and a new puzzle is generated.