package repository

import "backend_bench/internal/model"

var Stats = &model.WikiStats{
	DistinctUsers: make(map[string]struct{}),
	DistinctUrl:   make(map[string]int),
}
