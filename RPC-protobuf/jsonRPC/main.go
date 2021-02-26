package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
//通过变量与bitcoind进行交互
func bitcoincode(s string) string {
	str := strings.Fields(s)
	url := "localhost:8332"
	curl1 := `{"jsonrpc":"1.0","id":"curltest","method":"`
	curl2 := `","params":[`
	curl3 := `]}`

	var quest string
	switch len(str) {
	case 1:
		quest = fmt.Sprintln(curl1 + str[0] + curl2 + curl3)
	case 2:
		quest = fmt.Sprintln(curl1 + str[0] + curl2 + "\"" + str[1] + "\"" + curl3)
	}

	fmt.Println(quest)
	var jsonStr = []byte(quest)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return string(body)
}


//实现web命令及暴露端口
func run() {
	http.HandleFunc("/block_chain/getbalance", blockChainGetBalance)
	http.HandleFunc("/block_chain/getwalletinfo", blockChainGetWalletInfo)
	http.HandleFunc("/block_chain/getbestblockhash", blockChainGetBestBlockHash)
	http.HandleFunc("/block_chain/getblock", blockChainGetBlock)
	http.ListenAndServe(":8332", nil)
}

//实现getbalance
func blockChainGetBalance(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, bitcoincode("getbalance"))
}


//实现getwalletinfo
func blockChainGetWalletInfo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, bitcoincode("getwalletinfo"))
}

//实现GetBestBlockHash
func blockChainGetBestBlockHash(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, bitcoincode("getbestblockhash"))
}

//实现GetBlock
func blockChainGetBlock(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	io.WriteString(w, bitcoincode(blockData))
}*/

const (
	rpcURL = "http://127.0.0.1:8332"
	JSONRPCVERSION = "2.0"
	method = "getblockcount"
)

var (
	//ID 自增id
	ID int
)

type btcinfo struct {
	Result	interface{}	`json:"result"`
	Error 	string		`json:"error"`
	Id		int64		`json:"id"`
}



func bitcoincode1(s string)  {

	var(
		//str string
		//btcheight string
		quest string
	)


	const(
		url string = "http://127.0.0.1:8332"
		curl1 = `{"jsonrpc":"2.0","id":"111","method":"`
		curl2 = `","params":[`
		curl3 = `]}`

	)

	quest = fmt.Sprintln(curl1 + s + curl2 + curl3)
	var jsonStr = []byte(quest)

	req,err:=http.NewRequest("POST",url,bytes.NewBuffer(jsonStr))
	if nil!=err{
		fmt.Println("request failed ",err)
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth("icarus", "keystore")

	client := &http.Client{}
	resp,err:=client.Do(req)
	if nil!=err{
		fmt.Println("get response failed",err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("直接返回body:",string(body))

	var	blockinfo btcinfo
	json.Unmarshal(body,&blockinfo)
	//btcheight1 := string(body)

	switch s {
	case "getblockcount":
		fmt.Println("\n当前区块高度:",blockinfo.Result.(float64))
	case "getdifficulty":
		fmt.Println("\n当前区块链难度:",blockinfo.Result.(float64))
	case "getnetworkhashps":
		fmt.Println("\n全网哈希生成速率:",blockinfo.Result.(float64))
	}

	//fmt.Println("当前区块高度:",blockinfo.Result)

}


func bitcoincode(s string)  {
	var (
		formData = map[string]interface{}{
			"jsonrpc": JSONRPCVERSION,
			"id":      ID + 1,
			"method":  s,
			//"params":  [""],

		}
		httpClient    = &http.Client{}
		reqJSON 	 []byte
		data			[]byte
		req           *http.Request
		resp          *http.Response

		err error
	)
	reqJSON, err = json.Marshal(formData)
	payloadBuffer := bytes.NewReader(reqJSON)
	req, err = http.NewRequest("POST", rpcURL, payloadBuffer)
	if err != nil {
		log.Print(err.Error())
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth("icarus", "keystore")
	resp, err = httpClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	btcheight := string(data)
	fmt.Println("当前区块信息：",btcheight)
}



func main() {
	//bitcoincode("getblockcount")
	bitcoincode1("getblockcount")
	bitcoincode1("getdifficulty")
	bitcoincode1("getnetworkhashps")

}