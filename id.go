package types

import "encoding/binary"

type ID struct {
	ID uint64 `json:"id"`
}

func (i *ID) SetID(id uint64) {
	i.ID = id
}

func (i *ID) GetID() uint64 {
	return i.ID
}

func (i *ID) SetKey(s []byte) {
	i.ID = binary.BigEndian.Uint64(s)
}

func (i *ID) GetKey() []byte {
	k := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(k, i.ID)
	return k
}

func (i *ID) Combine(user, entity uint32) {
	u, e := make([]byte, 4, 4), make([]byte, 4, 4)
	binary.BigEndian.PutUint32(u, user)
	binary.BigEndian.PutUint32(e, entity)
	res := append(u, e...)
	i.ID = binary.BigEndian.Uint64(res)
}
