package entityS

type RPCResult struct {
	Id     int64       `json:"id"`
	Error  string      `json:"error"`
	Result interface{} `json:"result"`
}