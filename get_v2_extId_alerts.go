package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	type ResultData struct {
		Alerts []struct {
			AlertLevel   string      `json:"alertLevel"`
			Category     interface{} `json:"category"`
			Service      interface{} `json:"service"`
			ExternalID   string      `json:"externalId"`
			AlertAmount  float64     `json:"alertAmount"`
			ExposureType string      `json:"exposureType"`
		} `json:"alerts"`
	}
	// token:KYT API Key
	token := "8e283fdd013173ce7712d8e5ae9e9d21655274956bddbb40388a4a9229b135a5"
	hoo_external_Id := "7427b7ee-2abb-3d54-8665-44bba59f4b0e" //post_v2_tra_only_R.go 回传 external Id
	kyt_v2_url := "https://api.chainalysis.com/api/kyt/v2/transfers/" + hoo_external_Id + "/alerts"
	//kyt_api_get_user_url:"https://api.chainalysis.com/api/kyt/v1/users"
	client := &http.Client{}
	// get请求
	req, err := http.NewRequest("GET", kyt_v2_url, nil)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	// 在请求头中加入校验的token
	req.Header.Set("Token", token)
	//print(req.Body)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("status", resp.Status)
	fmt.Printf("%s\n", bodyText)
	var resul ResultData
	if err := json.Unmarshal(bodyText, &resul); err == nil {
		fmt.Println("risk:", resul.Alerts[0].AlertAmount)
		fmt.Println("AlertLevel:", resul.Alerts[0].AlertLevel)

		//fmt.Println("cluster:", resul[0].ChainalysisIdentification)
	}

}
