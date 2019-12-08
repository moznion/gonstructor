package constructor

import g "github.com/moznion/gowrtr/generator"

type Generator interface {
	Generate() g.Statement
}
