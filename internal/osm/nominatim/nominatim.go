package nominatim

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var nominatimReverseGeocode = "https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f"

func ReverseGeocode(lat, lon float64) (*Address, error) {
	url := fmt.Sprintf(nominatimReverseGeocode, lat, lon)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Address Address `json:"address"`
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result.Address, nil
}
