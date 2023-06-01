package sharedtimer

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *TimerConfig) Bind(r *http.Request) error {
	var err error
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, c)
	if err != nil {
		return err
	}

	return nil
}
