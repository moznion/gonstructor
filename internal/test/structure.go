package test

import "io"

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/gonstructor_test --type=Structure --constructorTypes=allArgs,builder --withGetter"
type Structure struct {
	foo string
	bar io.Reader
	Buz chan interface{}
	qux interface{} `gonstructor:"-"`
}

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/gonstructor_test --type=ChildStructure --output=./super_duper_child_structure_gen.go"
type ChildStructure struct {
	structure *Structure
	foobar    string
}

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/gonstructor_test --type=StructureWithInit --constructorTypes=allArgs,builder --withGetter --init initialize"
type StructureWithInit struct {
	foo    string
	status string      `gonstructor:"-"`
	qux    interface{} `gonstructor:"-"`
}

func (structure *StructureWithInit) initialize() {
	structure.status = "ok"
}
