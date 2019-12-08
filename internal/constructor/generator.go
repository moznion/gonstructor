package constructor

import g "github.com/moznion/gowrtr/generator"

// Generator is an interface that has the responsibility to generate a constructor code.
type Generator interface {
	// Generate generates a constructor statement.
	Generate() g.Statement
}
