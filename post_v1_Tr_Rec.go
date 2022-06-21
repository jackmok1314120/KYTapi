package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	kyt_uesrId := "hoo_test01"
	kyt_url := "https://api.chainalysis.com/api/kyt/v1/users/" + kyt_uesrId + "/transfers/received"
	token := "8e283fdd013173ce7712d8e5ae9e9d21655274956bddbb40388a4a9229b135a5"
	kyt_tx := "18070ce7844135c256b215336348c5b4e8b6db723e10802115132cd86b9cabc6"
	kyt_out_addr := "3MnUVwd8tLMAVtjnWR2N6UrBaD3yhwt9iB"
	type RequestBody struct {
		Asset      string `json:"asset"`
		TransferRe string `json:"transferReference"`
	}
	asset := "BTC"
	transfer_Re := kyt_tx + ":" + kyt_out_addr

	kyt_data := RequestBody{Asset: asset, TransferRe: transfer_Re}

	requestBody, err := json.Marshal(&kyt_data)

	list1 := `[` + string(requestBody) + `]` //+[]
	var data_body = []byte(list1)
	//fmt.Println("new_str", bytes.NewBuffer(list))
	req, err := http.NewRequest("POST", kyt_url, bytes.NewBuffer(data_body))
	req.Header.Set("Token", token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("status", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
