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
	baseURL := os.Getenv("BINDERBYTE_BASE_URL")

	if apiKey == "" || baseURL == "" {
		return nil, fmt.Errorf("env not set")
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

	var raw dto.TrackingRawResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	history := []dto.HistoryEntry{}
	_ = json.Unmarshal(raw.Data.History, &history)

	return &dto.TrackingResponse{
		Status:  raw.Status,
		Message: raw.Message,
		Data: dto.Data{
			Summary: raw.Data.Summary,
			Detail:  raw.Data.Detail,
			History: history,
		},
	}, nil
}
