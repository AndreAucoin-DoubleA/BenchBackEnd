package model

import "sync"

type RecentChange struct {
	ID        int64  `json:"id"`
	User      string `json:"user"`
	Bot       bool   `json:"bot"`
	ServerURL string `json:"server_url"`
}

type WikiStats struct {
	sync.Mutex                        // embed mutex
	TotalChanges  int                 `json:"total_changes"`
	DistinctUsers map[string]struct{} `json:"-"` // internal set for counting
	NumBots       int                 `json:"num_bots"`
	NumNonBots    int                 `json:"num_nonbots"`
	DistinctUrl   map[string]int      `json:"distinct_url"`
}
