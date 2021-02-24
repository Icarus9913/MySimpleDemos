package tool

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var formData = map[string]interface{}{
	"jsonrpc": "2.0",
	"method":  "Filecoin.ChainHead",
	"id":      1,
}

type FilResp struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  Result `json:"result"`
	ID      int    `json:"id"`
}

type Result struct {
	Height float64 `json:"Height"`
}

func Height() int64 {
	by, err := json.Marshal(formData)
	if nil != err {
		panic(err)
	}

	req, err := http.NewRequest("POST", "http://192.168.1.234:1235/rpc/v0", bytes.NewReader(by))
	if nil != err {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if nil != err {
		panic(err)
	}
	defer resp.Body.Close()

	readAll, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}
	respMsg := FilResp{}
	err = json.Unmarshal(readAll, &respMsg)
	if nil != err {
		panic(err)
	}
	return int64(respMsg.Result.Height)
}
