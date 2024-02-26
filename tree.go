package types

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

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
