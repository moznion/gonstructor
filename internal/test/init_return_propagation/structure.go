package init_return_propagation

import "errors"

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/gonstructor_test --type=StructureWithInitFoo --constructorTypes=allArgs,builder --init initialize -propagateInitFuncReturns"
type StructureWithInitFoo struct {
	foo     string
	checked bool `gonstructor:"-"`
}

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/gonstructor_test --type=StructureWithInitBar --constructorTypes=allArgs,builder --init initializeWithActualValueReceiver -propagateInitFuncReturns"
type StructureWithInitBar struct {
	foo     string
	checked bool `gonstructor:"-"`
}

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/gonstructor_test --type=StructureWithInitBuz --constructorTypes=allArgs,builder --init initializeWithError -propagateInitFuncReturns"
type StructureWithInitBuz struct {
	foo     string
	checked bool `gonstructor:"-"`
}

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/gonstructor_test --type=StructureWithInitQux --constructorTypes=allArgs,builder --init initializeWithMultipleReturns -propagateInitFuncReturns"
type StructureWithInitQux struct {
	foo     string
	checked bool `gonstructor:"-"`
}

func (s *StructureWithInitFoo) initialize() {
	s.checked = true
}

func (s StructureWithInitBar) initializeWithActualValueReceiver() {
}

func (s *StructureWithInitBuz) initializeWithError() error {
	s.checked = true
	return errors.New("err")
}

func (s *StructureWithInitQux) initializeWithMultipleReturns() (*string, error) {
	s.checked = true

	str := "str"
	return &str, errors.New("err")
}
