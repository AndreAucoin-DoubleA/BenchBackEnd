package status

import (
	"backend_bench/internal/repository"
	"encoding/json"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Lock the stats for thread safety
	repository.Stats.Lock()
	defer repository.Stats.Unlock()

	// Build a response object
	resp := struct {
		TotalChanges  int            `json:"total_changes"`
		DistinctUsers int            `json:"distinct_users"`
		NumBots       int            `json:"num_bots"`
		NumNonBots    int            `json:"num_nonbots"`
		DistinctUrl   map[string]int `json:"distinct_url"`
	}{
		TotalChanges:  repository.Stats.TotalChanges,
		DistinctUsers: len(repository.Stats.DistinctUsers),
		NumBots:       repository.Stats.NumBots,
		NumNonBots:    repository.Stats.NumNonBots,
		DistinctUrl:   repository.Stats.DistinctUrl,
	}

	// Marshal the response into JSON bytes
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to encode stats", http.StatusInternalServerError)
		return
	}

	// Write headers and JSON body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
