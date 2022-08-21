package multitypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFooAndBarStructureAllArgsConstructor(t *testing.T) {
	givenFooString := "foo"
	givenBarInt := 12345
	givenBuzString := "buz"
	givenQuxInt := 54321

	gotAlpha := NewAlphaStructure(givenFooString, givenBarInt)
	gotBravo := NewBravoStructure(givenBuzString, givenQuxInt)

	assert.IsType(t, &AlphaStructure{}, gotAlpha)
	assert.IsType(t, &BravoStructure{}, gotBravo)

	assert.EqualValues(t, givenFooString, gotAlpha.foo)
	assert.EqualValues(t, givenBarInt, gotAlpha.bar)

	assert.EqualValues(t, givenBuzString, gotBravo.buz)
	assert.EqualValues(t, givenQuxInt, gotBravo.qux)

	// test for getters
	assert.EqualValues(t, givenFooString, gotAlpha.GetFoo())
	assert.EqualValues(t, givenBarInt, gotAlpha.GetBar())

	assert.EqualValues(t, givenBuzString, gotBravo.GetBuz())
	assert.EqualValues(t, givenQuxInt, gotBravo.GetQux())
}

func TestStructureBuilder(t *testing.T) {
	givenFooString := "foo"
	givenBarInt := 12345
	givenBuzString := "buz"
	givenQuxInt := 54321

	alphaBuilder := NewAlphaStructureBuilder()
	gotAlpha := alphaBuilder.Foo(givenFooString).
		Bar(givenBarInt).
		Build()
	assert.IsType(t, &AlphaStructure{}, gotAlpha)

	assert.EqualValues(t, givenFooString, gotAlpha.foo)
	assert.EqualValues(t, givenBarInt, gotAlpha.bar)

	bravoBuilder := NewBravoStructureBuilder()
	gotBravo := bravoBuilder.Buz(givenBuzString).
		Qux(givenQuxInt).
		Build()
	assert.IsType(t, &BravoStructure{}, gotBravo)

	assert.EqualValues(t, givenBuzString, gotBravo.buz)
	assert.EqualValues(t, givenQuxInt, gotBravo.qux)
}
