package test

import "io"

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/bin/gonstructor --type=Structure --constructorTypes=allArgs,builder"
type Structure struct {
	foo string
	bar io.Reader
	Buz chan interface{}
	qux interface{} `gonstructor:"-"`
}

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/bin/gonstructor --type=ChildStructure"
type ChildStructure struct {
	structure *Structure
	foobar    string
}
