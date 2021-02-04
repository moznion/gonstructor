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

	// test for getters
	assert.EqualValues(t, givenString, got.GetFoo())
	assert.EqualValues(t, givenIOReader, got.GetBar())
	assert.EqualValues(t, givenChan, got.GetBuz())
	assert.EqualValues(t, nil, got.GetQux())
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

func TestStructureWithInitAllArgsConstructor(t *testing.T) {
	givenString := "givenstr"

	got := NewStructureWithInit(givenString)
	assert.IsType(t, &StructureWithInit{}, got)

	assert.EqualValues(t, givenString, got.foo)
	assert.EqualValues(t, "ok", got.status)
	assert.EqualValues(t, nil, got.qux)

	// test for getters
	assert.EqualValues(t, givenString, got.GetFoo())
	assert.EqualValues(t, "ok", got.GetStatus())
	assert.EqualValues(t, nil, got.GetQux())
}

func TestStructureWithInitBuilder(t *testing.T) {
	givenString := "givenstr"

	b := NewStructureWithInitBuilder()
	got := b.Foo(givenString).
		Build()
	assert.IsType(t, &StructureWithInit{}, got)

	assert.EqualValues(t, givenString, got.foo)
	assert.EqualValues(t, "ok", got.status)
	assert.EqualValues(t, nil, got.qux)
}

func TestStructureWithEmbeddingAllArgsConstructor(t *testing.T) {
	got := NewStructureWithEmbedding(Embedded{Bar: "bar"}, "foo")
	assert.IsType(t, &StructureWithEmbedding{}, got)

	assert.EqualValues(t, "foo", got.foo)
	assert.EqualValues(t, "bar", got.Bar)
	assert.EqualValues(t, "bar", got.Embedded.Bar)

	// test for getters
	assert.EqualValues(t, "foo", got.GetFoo())
	assert.EqualValues(t, "bar", got.GetEmbedded().Bar)
}

func TestStructureWithEmbeddingBuilder(t *testing.T) {
	b := NewStructureWithEmbeddingBuilder()
	got := b.Foo("foo").Embedded(Embedded{Bar: "bar"}).Build()
	assert.IsType(t, &StructureWithEmbedding{}, got)

	assert.EqualValues(t, "foo", got.foo)
	assert.EqualValues(t, "bar", got.Bar)
	assert.EqualValues(t, "bar", got.Embedded.Bar)
}

func TestStructureWithPointerEmbeddingAllArgsConstructor(t *testing.T) {
	got := NewStructureWithPointerEmbedding(Embedded{Bar: "bar"}, "foo")
	assert.IsType(t, &StructureWithPointerEmbedding{}, got)

	assert.EqualValues(t, "foo", got.foo)
	assert.EqualValues(t, "bar", got.Bar)
	assert.EqualValues(t, "bar", got.Embedded.Bar)

	// test for getters
	assert.EqualValues(t, "foo", got.GetFoo())
	assert.EqualValues(t, "bar", got.GetEmbedded().Bar)
}

func TestStructureWithPointerEmbeddingBuilder(t *testing.T) {
	b := NewStructureWithPointerEmbeddingBuilder()
	got := b.Foo("foo").Embedded(Embedded{Bar: "bar"}).Build()
	assert.IsType(t, &StructureWithPointerEmbedding{}, got)

	assert.EqualValues(t, "foo", got.foo)
	assert.EqualValues(t, "bar", got.Bar)
	assert.EqualValues(t, "bar", got.Embedded.Bar)
}
