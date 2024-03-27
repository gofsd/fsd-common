package types

var (
	_ ICrud      = &Default{}
	_ IRootFlags = &Default{}
)

type Default struct {
	ID
	RootFlags
}

// func (def *Default) FromString(s string) error {
// 	return json.Unmarshal([]byte(s), def)
// }

// func (def *Default) String() string {
// 	data, e := json.Marshal(def)
// 	if e != nil {
// 		return ""
// 	}
// 	return string(data)
// }

// func (def *Default) FromJson(s []byte) error {
// 	return json.Unmarshal(s, def)
// }

// func (def *Default) Json() ([]byte, error) {
// 	return json.Marshal(def)
// }
