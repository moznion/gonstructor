package constructor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithPrefix(t *testing.T) {
	t.Run("apply prefix", func(t *testing.T) {
		got := withPrefix("Struct", "*", true)
		assert.Equal(t, "*Struct", got)
	})

	t.Run("ignore prefix", func(t *testing.T) {
		got := withPrefix("Struct", "*", false)
		assert.Equal(t, "Struct", got)
	})
}
