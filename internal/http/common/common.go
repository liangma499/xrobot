package common

type Res struct {
	Code int32 `json:"code"`
	Data any   `json:"data,omitempty"`
}
