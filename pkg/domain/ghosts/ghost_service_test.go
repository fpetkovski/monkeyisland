package ghosts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateGhosts(t *testing.T) {
	ghosts := GenerateGhosts()
	assert.Equal(t, 10, len(ghosts))
}