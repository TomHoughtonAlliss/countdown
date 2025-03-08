package numbers

// Config stores values needed to construct a puzzle.
type Config struct {
	Input string
	Lower int
	Upper int
}

func NewConfig(input string, l, u int) Config {
	return Config{
		Input: input,
		Lower: l,
		Upper: u,
	}
}
