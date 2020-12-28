package RpcServer

import (
	"../entityS"
	"../utils"

	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Cremap() map[string]string {
	Map := make(map[string]string)
	Map["Encoding"] = "UTF-8"
	Map["Content-Type"] = "application"
	Map["Authorization"] = "Basic " + utils.Base64Str(entityS.RPCUSER+":"+entityS.RPCPASSWORD)
	return Map
}

func PareJSON(method string, Paramsl interface{}) string {

	rpcReq := entityS.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  method,
		Jsonrpc: "2.0",
		Params:  Paramsl,
	}

	reqBytes, err := json.Marshal(&rpcReq)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(reqBytes)
}

func Dopost(reqBytes string, header map[string]string, url string) *entityS.RPCResult{

	//client： 客户端 ，客户端用于发起请求 http.client
	//method 方法  指的请求类型 是post 还是 get
	//client : 客户端 客户端用于发送请求
	client := http.Client{} //
	//
	request, err := http.NewRequest("POST", url, bytes.NewReader([]byte(reqBytes)))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	//给post请求添加请求头设置信息
	//key -->value
	//request.Header.Add("Encoding","UTF-8")
	//request.Header.Add("Content-Type","application")
	//request.Header.Add("Authorization","Basic "+utils.Base64Str(entityS.RPCUSER+":"+entityS.RPCPASSWORD))
	for k, v := range header {
		request.Header.Add(k, v)

	}
	//
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	//
	code := response.StatusCode
	if code == 200 {
		fmt.Println("请求成功")
		resultBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		fmt.Println("rpc调用的响应结果:" + string(resultBytes))
		//json的反序列化
		var result entityS.RPCResult
		err = json.Unmarshal(resultBytes, &result)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		//反序列化正常，没有出现错误
		return &result
	} else {
		fmt.Println("请求失败")
		return nil
	}
	return nil

}
