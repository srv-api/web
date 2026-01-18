package resi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	dto "github.com/srv-api/web/dto"
)

func (r *resiRepository) Track(courier, awb string) (*dto.TrackingResponse, error) {
	apiKey := os.Getenv("BINDERBYTE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("BINDERBYTE_API_KEY not set")
	}

	baseURL := os.Getenv("BINDERBYTE_BASE_URL")
	if baseURL == "" {
		return nil, fmt.Errorf("BINDERBYTE_BASE_URL not set")
	}

	url := fmt.Sprintf(
		"%s/track?api_key=%s&courier=%s&awb=%s",
		baseURL, apiKey, courier, awb,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.TrackingResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
