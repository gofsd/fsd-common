package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	def := Testing{}
	var str string
	Json(def, &str)
	assert.Equal(t, str, "")
}
