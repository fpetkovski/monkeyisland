package ghosts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGhost(t *testing.T) {
	ghost1 := NewGhost("first ghost name")
	ghost2 := NewGhost("second ghost name")
	assert.Equal(t, uint64(1), ghost2.ID - ghost1.ID, "Ghost 2 ID must be greater than Ghost 1 ID")
}
