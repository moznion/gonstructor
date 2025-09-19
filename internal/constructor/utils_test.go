package constructor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithConditionalPrefix(t *testing.T) {
	t.Run("apply prefix", func(t *testing.T) {
		got := withConditionalPrefix("Struct", "*", true)
		assert.Equal(t, "*Struct", got)
	})

	t.Run("ignore prefix", func(t *testing.T) {
		got := withConditionalPrefix("Struct", "*", false)
		assert.Equal(t, "Struct", got)
	})
}
