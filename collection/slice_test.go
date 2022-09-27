package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	slice := []string{"a", "b", "c"}
	slice = append(slice, "c")

	x := StringContain(slice, "a")
	assert.Equal(t, x, true)
}
