package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEntityKey(t *testing.T) {
	id := EntityID(1)
	assert.Equal(t, []byte{0x1}, id.GetKey())
}

func TestSetEntityKey(t *testing.T) {
	id := EntityID(1)
	id.SetKey([]byte{0x4})
	assert.Equal(t, []byte{0x4}, id.GetKey())
}

func TestGetEntityID(t *testing.T) {
	id := EntityID(1)
	id.SetKey([]byte{0x4})
	assert.Equal(t, EntityID(4), id.GetID())
}

func TestSetEntityID(t *testing.T) {
	id := EntityID(1)
	id.SetID(4)
	assert.Equal(t, []byte{0x4}, id.GetKey())
}

func TestGetUint64(t *testing.T) {
	id := EntityID(1)
	id.SetID(4)
	assert.Equal(t, uint64(4), id.GetUint64())
}

func TestGetCommonID(t *testing.T) {
	id := CommonID{}
	b := []byte("some text")
	b = append(b, PARENT_L1_BYTE)
	id.Set(string(b))

	assert.Equal(t, ID("some tex"), id.GetID())

}
