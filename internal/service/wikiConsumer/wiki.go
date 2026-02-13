package wikiconsumer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type RecentChange struct {
	ID        int64  `json:"id"`
	User      string `json:"user"`
	Bot       bool   `json:"bot"`
	ServerURL string `json:"server_url"`
}

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
		fmt.Println("Raw JSON:", jsonData)
		var change RecentChange
		if err := json.Unmarshal([]byte(jsonData), &change); err != nil {
			continue
		}

		if _, ok := seen[change.ID]; ok { // skip duplicates
			continue
		}
		seen[change.ID] = struct{}{}

		fmt.Printf("ID: %d | User: %s | Bot: %t | Server URL: %s\n",
			change.ID, change.User, change.Bot, change.ServerURL)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
}
