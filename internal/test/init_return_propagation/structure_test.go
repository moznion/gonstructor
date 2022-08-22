package init_return_propagation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitFuncPropagation_InitializeWithNoReturnValue(t *testing.T) {
	givenString := "givenstr"

	got := NewStructureWithInitFoo(givenString)
	assert.IsType(t, &StructureWithInitFoo{}, got)
	assert.EqualValues(t, givenString, got.foo)
	assert.True(t, true, got.checked)

	got = NewStructureWithInitFooBuilder().Foo(givenString).Build()
	assert.IsType(t, &StructureWithInitFoo{}, got)
	assert.EqualValues(t, givenString, got.foo)
	assert.True(t, true, got.checked)
}

func TestInitFuncPropagation_InitializeWithActualValueReceiver(t *testing.T) {
	givenString := "givenstr"

	got := NewStructureWithInitBar(givenString)
	assert.IsType(t, &StructureWithInitBar{}, got)
	assert.EqualValues(t, givenString, got.foo)
	assert.True(t, true, got.checked)

	got = NewStructureWithInitBarBuilder().Foo(givenString).Build()
	assert.IsType(t, &StructureWithInitBar{}, got)
	assert.EqualValues(t, givenString, got.foo)
	assert.True(t, true, got.checked)
}

func TestInitFuncPropagation_InitializeWithError(t *testing.T) {
	givenString := "givenstr"

	got, err := NewStructureWithInitBuz(givenString)
	assert.IsType(t, &StructureWithInitBuz{}, got)
	assert.EqualValues(t, givenString, got.foo)
	assert.True(t, true, got.checked)
	assert.EqualError(t, err, "err")

	got, err = NewStructureWithInitBuzBuilder().Foo(givenString).Build()
	assert.IsType(t, &StructureWithInitBuz{}, got)
	assert.EqualValues(t, givenString, got.foo)
	assert.True(t, true, got.checked)
	assert.EqualError(t, err, "err")
}

func TestInitFuncPropagation_InitializeWithMultipleReturns(t *testing.T) {
	givenString := "givenstr"

	got, gotStr, err := NewStructureWithInitQux(givenString)
	assert.IsType(t, &StructureWithInitQux{}, got)
	assert.EqualValues(t, givenString, got.foo)
	assert.True(t, true, got.checked)
	assert.EqualValues(t, *gotStr, "str")
	assert.EqualError(t, err, "err")

	got, gotStr, err = NewStructureWithInitQuxBuilder().Foo(givenString).Build()
	assert.IsType(t, &StructureWithInitQux{}, got)
	assert.EqualValues(t, givenString, got.foo)
	assert.True(t, true, got.checked)
	assert.EqualValues(t, *gotStr, "str")
	assert.EqualError(t, err, "err")
}
