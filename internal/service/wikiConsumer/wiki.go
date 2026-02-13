package wikiconsumer

import (
	"backend_bench/internal/model"
	"backend_bench/internal/repository"
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func StartWikiConsumer(streamURL string) {
	fmt.Println("Connecting to:", streamURL)

	req, err := http.NewRequest("GET", streamURL, nil)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}

	// Set headers for SSE and bot policy
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("User-Agent", "backend-bench-dev")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Connected! Reading stream...")

	seen := make(map[int64]struct{}) // track processed events
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data:") {
			continue
		}

		jsonData := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		var change model.RecentChange
		if err := json.Unmarshal([]byte(jsonData), &change); err != nil {
			continue
		}

		if _, ok := seen[change.ID]; ok { // skip duplicates
			continue
		}
		seen[change.ID] = struct{}{}

		updateStats(change)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
}

func updateStats(change model.RecentChange) {
	repository.Stats.Lock()
	defer repository.Stats.Unlock()

	repository.Stats.TotalChanges++
	repository.Stats.DistinctUsers[change.User] = struct{}{}
	repository.Stats.DistinctUrl[change.ServerURL]++
	if change.Bot {
		repository.Stats.NumBots++
	} else {
		repository.Stats.NumNonBots++
	}
}
