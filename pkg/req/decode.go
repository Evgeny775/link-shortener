package req

import (
	"encoding/json"
	"io"
)

func Decode[payloadType any](body io.ReadCloser) (*payloadType, error) {
	var payloadRequest payloadType
	err := json.NewDecoder(body).Decode(&payloadRequest)
	if err != nil {
		return &payloadRequest, err
	}
	return &payloadRequest, nil
}
