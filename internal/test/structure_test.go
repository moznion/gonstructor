package test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructureAllArgsConstructor(t *testing.T) {
	givenString := "givenstr"
	givenIOReader := strings.NewReader(givenString)
	givenChan := make(chan interface{}, 0)

	got := NewStructure(givenString, givenIOReader, givenChan)
	assert.IsType(t, &Structure{}, got)

	assert.EqualValues(t, givenString, got.foo)
	assert.EqualValues(t, givenIOReader, got.bar)
	assert.EqualValues(t, givenChan, got.Buz)
	assert.EqualValues(t, nil, got.qux)
}

func TestStructureBuilder(t *testing.T) {
	givenString := "givenstr"
	givenIOReader := strings.NewReader(givenString)
	givenChan := make(chan interface{}, 0)

	b := NewStructureBuilder()
	got := b.Foo(givenString).
		Bar(givenIOReader).
		Buz(givenChan).
		Build()
	assert.IsType(t, &Structure{}, got)

	assert.EqualValues(t, givenString, got.foo)
	assert.EqualValues(t, givenIOReader, got.bar)
	assert.EqualValues(t, givenChan, got.Buz)
	assert.EqualValues(t, nil, got.qux)
}

func TestChildStructureAllArgsConstructor(t *testing.T) {
	structure := NewStructure("givenstr", strings.NewReader("givenstr"), make(chan interface{}, 0))
	givenString := "foobar"

	got := NewChildStructure(structure, givenString)
	assert.IsType(t, &ChildStructure{}, got)

	assert.EqualValues(t, structure, got.structure)
	assert.EqualValues(t, givenString, got.foobar)
}
