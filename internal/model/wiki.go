package model

type Wiki struct {
	Consumed      int            `json:"consumed"`
	DistinctUsers int            `json:"distinct_users"`
	NumBots       int            `json:"num_bots"`
	NumNonBots    int            `json:"num_nonbots"`
	DistinctUrl   map[string]int `json:"distinct_url"`
}
