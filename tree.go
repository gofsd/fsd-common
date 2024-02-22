package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Pair struct {
	K uint16 `json:"i" validate:"min=1,max=65535"`
	V string `json:"v" validate:"min=1,max=255"`
}

type KV struct {
	Key   string `json:"key" validate:"omitempty,min=0,max=9"`
	Value string `json:"value" validate:"omitempty,min=0,max=256"`
}

type Command struct {
	Name  []string `json:"name" validate:"required,min=1,max=9,dive,min=2,max=16"`
	Args  []string `json:"args" validate:"omitempty,min=0,max=9,dive,min=1,max=64"`
	Flags []KV     `json:"flags" validate:"omitempty,min=0,max=9"`
}

type CommandResponse struct {
	Id      uint64 `json:"id"`
	Command Command
	Result  any
}

func (c *Client) TreeCmd(cmd Command) (*Pair, error) {
	rb, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/exec", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newPair := Pair{}
	err = json.Unmarshal(body, &newPair)
	if err != nil {
		return nil, err
	}

	return &newPair, nil
}
