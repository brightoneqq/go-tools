package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsExistInMap(t *testing.T) {
	m := map[interface{}]interface{}{
		"A": "1",
		"B": "2",
	}

	assert.Equal(t, IsExistInMap(m, "A"), true)
	assert.Equal(t, IsExistInMap(m, "C"), false)
}
