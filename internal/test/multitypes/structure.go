package multitypes

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/gonstructor_test --type=AlphaStructure --type=BravoStructure --constructorTypes=allArgs,builder --withGetter --output=./alpha_and_bravo_gen.go"

type AlphaStructure struct {
	foo string
	bar int
}

type BravoStructure struct {
	buz string
	qux int
}
