package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	type ResultData struct {
		UpdatedAt         interface{} `json:"updatedAt"`
		Asset             string      `json:"asset"`
		TransferReference string      `json:"transferReference"`
		Tx                int         `json:"tx"`
		Idx               interface{} `json:"idx"`
		UsdAmount         interface{} `json:"usdAmount"`
		AssetAmount       interface{} `json:"assetAmount"`
		Timestamp         int         `json:"timestamp"`
		OutputAddress     interface{} `json:"outputAddress"`
		ExternalID        string      `json:"externalId"`
	}
	wallet_Id := "hoo_test01"
	kyt_url := "https://api.chainalysis.com/api/kyt/v2/users/" + wallet_Id + "/transfers"
	token := "8e283fdd013173ce7712d8e5ae9e9d21655274956bddbb40388a4a9229b135a5"
	kyt_tx := "18070ce7844135c256b215336348c5b4e8b6db723e10802115132cd86b9cabc6"
	kyt_out_addr := "3MnUVwd8tLMAVtjnWR2N6UrBaD3yhwt9iB"
	asset := "BTC1"
	//external id:7427b7ee-2abb-3d54-8665-44bba59f4b0e
	type RequestBody struct {
		Asset      string `json:"asset"`
		TransferRe string `json:"transferReference"`
		Direction  string `json:"direction"`
	}

	transfer_Re := kyt_tx + ":" + kyt_out_addr

	kyt := RequestBody{Asset: asset, TransferRe: transfer_Re, Direction: "received"}

	//fmt.Println(kyt_url, "post", kyt)

	requestBody, err := json.Marshal(&kyt)

	req, err := http.NewRequest("POST", kyt_url, bytes.NewBuffer(requestBody))
	req.Header.Set("Token", token)
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
	var resul ResultData
	if err := json.Unmarshal(body, &resul); err != nil {

		fmt.Println("external:", resul.ExternalID)
		//return resul.ExternalID
	}

}
