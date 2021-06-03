package Models

type ModelResult struct {
	Code  int `json:"code"`
	Msg   string `json:"msg"`
	Model Address `json:"model"`
}
