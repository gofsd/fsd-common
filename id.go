package types

import (
	"encoding/binary"
	"sync/atomic"
)

type CommonID struct {
	ID uint64 `json:"id"`
}

var (
	RequestID atomic.Uint64
)

func (i *CommonID) SetID(id uint64) {
	i.ID = id
}

func (i *CommonID) GetID() uint64 {
	return i.ID
}

func (i *CommonID) SetKey(s []byte) {
	i.ID = binary.BigEndian.Uint64(s)
}

func (i *CommonID) GetKey() []byte {
	k := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(k, i.ID)
	return k
}

func (i *CommonID) Combine(user, entity uint32) {
	RequestID.Add(1)
	u, e := make([]byte, 4, 4), make([]byte, 4, 4)
	binary.BigEndian.PutUint32(u, user)
	binary.BigEndian.PutUint32(e, entity)
	res := append(u, e...)
	i.ID = binary.BigEndian.Uint64(res)
}
