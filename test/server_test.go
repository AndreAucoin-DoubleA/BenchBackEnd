package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend_bench/internal/handler/status"
	"backend_bench/internal/model"
	"backend_bench/internal/repository"
)

func TestStatsHandler(t *testing.T) {
	repository.Stats = &model.WikiStats{
		TotalChanges:  2,
		DistinctUsers: map[string]struct{}{"Alice": {}, "Bob": {}},
		NumBots:       1,
		NumNonBots:    1,
		DistinctUrl:   map[string]int{"url1": 1, "url2": 1},
	}

	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	w := httptest.NewRecorder()

	status.StatusHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var statsResp struct {
		TotalChanges  int            `json:"total_changes"`
		DistinctUsers int            `json:"distinct_users"`
		NumBots       int            `json:"num_bots"`
		NumNonBots    int            `json:"num_nonbots"`
		DistinctUrl   map[string]int `json:"distinct_url"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&statsResp); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	if statsResp.TotalChanges != 2 {
		t.Errorf("Expected TotalChanges=2, got %d", statsResp.TotalChanges)
	}
	if statsResp.DistinctUsers != 2 {
		t.Errorf("Expected DistinctUsers=2, got %d", statsResp.DistinctUsers)
	}
	if statsResp.NumBots != 1 || statsResp.NumNonBots != 1 {
		t.Errorf("Expected NumBots=1 and NumNonBots=1")
	}
	if len(statsResp.DistinctUrl) != 2 {
		t.Errorf("Expected 2 URLs, got %d", len(statsResp.DistinctUrl))
	}
}
