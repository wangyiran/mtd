package serializer

type Response struct {
	Status string      `json:"status"`
	Err    string      `json:"error"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}
