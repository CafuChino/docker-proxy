package utils

import (
	"bytes"
	"encoding/json"
	"io"
)

func ParseJsonReqBody(jsonBody io.ReadCloser)(parsedBody map[string]interface{}, err error) {
	bodyBuf := new(bytes.Buffer)
	_, _err := bodyBuf.ReadFrom(jsonBody)
	if _err != nil {
		bodyBuf = nil
	}
	bodyString := bodyBuf.String()
	err = json.Unmarshal([]byte(bodyString), &parsedBody)
	return
}
