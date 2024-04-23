package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"sync/atomic"
)

const (
	PARENT_LENGTH uint8 = iota
	PARENT_L1_BYTE
	PARENT_L2_BYTE
	PARENT_L3_BYTE
	PARENT_L4_BYTE
	PARENT_L5_BYTE
	PARENT_L6_BYTE
	PARENT_L7_BYTE
	PARENT_L8_BYTE
)
const (
	MAX_BYTE_ID_LENGTH   = 16
	MAX_STRING_ID_LENGTH = 40
)

var PARENT_FLAGS = []byte{PARENT_L1_BYTE, PARENT_L2_BYTE, PARENT_L3_BYTE, PARENT_L4_BYTE, PARENT_L5_BYTE, PARENT_L6_BYTE, PARENT_L7_BYTE, PARENT_L8_BYTE}

type ID[T any] struct {
	data T
}
type EntityID[T any] struct {
	parent, children T
}

type CommonID[T any] struct {
	entityID EntityID[uint64]
	ID[T]    `json:"id" mapstructure:"id" validate:"max=41"`
}

var (
	requestID atomic.Uint64
)
var def string
var _ Identifier[uint64] = &EntityID[uint64]{}
var _ Identifier[string] = &ID[string]{}

func (i *ID[uint64]) SetID(id uint64) {
	i.value = id
}

func (i *EntityID) GetID() uint64 {
	return i.value
}

func (i *EntityID) GetUint64() uint64 {
	return uint64(i.value)
}

func (i *EntityID) SetKey(s []byte) {
	buf := bytes.NewBuffer(s) // b is []byte

	idx, _ := binary.ReadUvarint(buf)
	i.value = idx
}

func (i *EntityID) GetKey() []byte {
	var k = make([]byte, 8)
	binary.BigEndian.PutUint64(k, i.value)
	return bytes.TrimLeft(k, "\x00")
}

func (i *EntityID) Combine(parent, entity uint32) {
	u, e := make([]byte, 4, 4), make([]byte, 4, 4)
	binary.BigEndian.PutUint32(u, parent)
	binary.BigEndian.PutUint32(e, entity)
	res := append(u, e...)
	i.value = binary.BigEndian.Uint64(res)
}

func (i *ID) Set(id string) error {
	if len(id) > MAX_STRING_ID_LENGTH {
		return Err("ELONG: %d/%d", len(id))
	}
	i.value = string(id)
	return nil
}

func (i *ID) Get() string {
	return i.value
}

func (i *ID) SetKey(key []byte) error {
	if len(key) > MAX_BYTE_ID_LENGTH {
		return fmt.Errorf("ID is too long")
	}
	i.value = string(key)
	return nil
}

func (i *ID) GetKey() []byte {
	return []byte(i.value)
}

func (i CommonID) SetParentID() (parent uint64) {
	k := make([]byte, 8, 8)
	k = append(k, i.parentID.GetKey()...)
	parent = binary.BigEndian.Uint64(k)
	return
}

func (i CommonID) GetParentID() (parent uint64) {
	k := make([]byte, 8, 8)
	k = append(k, i.parentID.GetKey()...)
	parent = binary.BigEndian.Uint64(k)
	return
}
func (i CommonID) GetParentKey() []byte {
	return []byte("ets")
}

func (i CommonID) GetID() (entity uint64) {
	k := make([]byte, 8, 8)
	k = append(k, i.entityID.GetKey()...)
	entity = binary.BigEndian.Uint64(k)
	return
}

func (i CommonID) GetKey() []byte {
	return []byte("ets")
}

func (i CommonID) SetKey() (entity uint64) {
	k := make([]byte, 8, 8)
	k = append(k, i.entityID.GetKey()...)
	entity = binary.BigEndian.Uint64(k)
	return
}

func (i CommonID) Increment() {
	entity := i.entityID.GetID()
	if entity < math.MaxUint32 {
		entity = entity + 1
		i.entityID.SetID(uint64(entity))
		requestID.Add(1)
	}
}

func SetRequestID(id uint64) {
	requestID.Store(id)
}

func GetRequestID() uint64 {
	return requestID.Load()
}
