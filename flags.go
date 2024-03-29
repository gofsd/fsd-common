package types

type RootFlags struct {
	Action, Command, Output uint8
	UserID, EntityID        uint32
	Flag                    []string
}

func (rf *RootFlags) ForEach(callback func(key, value string)) (e error) {
	kv := KV{}
	for _, v := range rf.Flag {
		e = FromJson(&v, &kv)
		if e != nil {
			return e
		}
		callback(kv.Key, kv.Value)
	}
	return e
}
