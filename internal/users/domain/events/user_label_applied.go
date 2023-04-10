package events

type UserLabelAppliedEvent struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
