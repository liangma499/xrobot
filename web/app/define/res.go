package define

type Res struct {
	Code int `json:"code"`
	Data any `json:"data,omitempty"`
}
