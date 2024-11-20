package dto

type SimpleFlow struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type SimpleFlowPost struct {
	Name      string `json:"name"`
	Active    *bool  `json:"active"`
	FirstStep int    `json:"first_step"`
	RunEvery  int    `json:"run_every"`
	NextRunAt int64  `json:"next_run_at"`
}
