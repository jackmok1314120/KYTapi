package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	type RequestBody struct {
		Asset          string `json:"asset"`
		Output_address string `json:"address"`
	}
	type ResultData []struct {
		Asset                     string      `json:"asset"`
		Address                   string      `json:"address"`
		Cluster                   interface{} `json:"cluster"`
		Rating                    string      `json:"rating"`
		CustomAddress             interface{} `json:"customAddress"`
		ChainalysisIdentification interface{} `json:"chainalysisIdentification"`
	}
	wallet_Id := "hoo_test01" //1
	kyt_url := "https://api.chainalysis.com/api/kyt/v1/users/" + wallet_Id + "/withdrawaladdresses"
	token := "8e283fdd013173ce7712d8e5ae9e9d21655274956bddbb40388a4a9229b135a5"
	asset := "BTC"                                  //2
	address := "359ditw3D7xyLneaV4oWunMoDy7dndVzTP" //3

	kyt := RequestBody{Asset: asset, Output_address: address} // data:byte
	requestBody, err := json.Marshal(&kyt)                    //json
	list := `[` + string(requestBody) + `]`                   //
	var data = []byte(list)
	fmt.Println("new_str", bytes.NewBuffer(data))

	req, err := http.NewRequest("POST", kyt_url, bytes.NewBuffer(data))
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
	var resul ResultData
	if err := json.Unmarshal(body, &resul); err == nil {
		//fmt.Println("response Body4:", resul[0])
		fmt.Println("risk:", resul[0].Rating)
		fmt.Println("cluster:", resul[0].Cluster)
	}
	//fmt.Println("response Body:", string(body))

}
