package entityS


type RPCRequest struct {
	Id      int64  `json:"id"`
	Method  string `json:"method"`
	Jsonrpc string `json:"jsonrpc"`
	Params  interface{} `json:"params"`
	//interface
	//在结构体后面加标签``  大写的字母就变成小写的了
}