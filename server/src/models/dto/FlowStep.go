package dto

type FlowStep struct {
	ID           int  `json:"id"`
	FlowID       int  `json:"flow_id"`
	PreviousStep *int `json:"previous_step"`
	NextStep     *int `json:"next_step"`
	ActionID     int  `json:"action_id"`
}

type FlowStepPost struct {
	FlowID       int  `json:"flow_id"`
	PreviousStep *int `json:"previous_step"`
	NextStep     *int `json:"next_step"`
	ActionID     int  `json:"action_id"`
}
