package test

import (
	"testing"

	"backend_bench/internal/model"
	"backend_bench/internal/repository"
	wikiconsumer "backend_bench/internal/service/wikiconsumer"
)

func TestUpdateStats(t *testing.T) {
	repository.Stats = &model.WikiStats{
		DistinctUsers: make(map[string]struct{}),
		DistinctUrl:   make(map[string]int),
	}

	change := model.RecentChange{
		User:      "Alice",
		Bot:       false,
		ServerURL: "https://en.wikipedia.org/wiki/OpenAI",
	}

	wikiconsumer.UpdateStats(change)

	if repository.Stats.TotalChanges != 1 {
		t.Errorf("Expected TotalChanges=1, got %d", repository.Stats.TotalChanges)
	}

	if len(repository.Stats.DistinctUsers) != 1 {
		t.Errorf("Expected 1 distinct user, got %d", len(repository.Stats.DistinctUsers))
	}

	if repository.Stats.NumBots != 0 {
		t.Errorf("Expected NumBots=0, got %d", repository.Stats.NumBots)
	}

	if repository.Stats.NumNonBots != 1 {
		t.Errorf("Expected NumNonBots=1, got %d", repository.Stats.NumNonBots)
	}

	if repository.Stats.DistinctUrl["https://en.wikipedia.org/wiki/OpenAI"] != 1 {
		t.Errorf("Expected 1 count for URL, got %d", repository.Stats.DistinctUrl["https://en.wikipedia.org/wiki/OpenAI"])
	}
}
