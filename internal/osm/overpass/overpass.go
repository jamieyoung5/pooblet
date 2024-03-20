// Interacts with the overpass API

package overpass

import (
	"bytes"
	"io"
	"net/http"
)

const overpassInterpreter = "http://overpass-api.de/api/interpreter"

func Query(query string) (response []byte, err error) {
	resp, err := http.Post(overpassInterpreter, "application/x-www-form-urlencoded", bytes.NewBufferString("data="+query))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
