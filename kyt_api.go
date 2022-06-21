package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*func Check user{
	client := &http.Client{}
	// get请求
	req, err := http.NewRequest("GET", "https://api.chainalysis.com/api/kyt/v1/users/"+ hoo_userId,nil)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	// 在请求头中加入校验的token
	req.Header.Set("Token",token)
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
	fmt.Printf("%s\n", bodyText)
}
type Payload []struct {
	Asset   string `json:"asset"`
	Address string `json:"address"`
}
	data := Payload{
		Asset = "BTC"
		Address = ""
// fill struct
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
	// handle err
	}
	body := bytes.NewReader(payloadBytes)*/

func main() {
	// token:KYT API Key
	token := "8e283fdd013173ce7712d8e5ae9e9d21655274956bddbb40388a4a9229b135a5"
	hoo_userId := "hoo_test01" //hoo.com user ID.
	kyt_v1_url := "https://api.chainalysis.com/api/kyt/v1/users/" + hoo_userId + "/withdrawaladdresses"
	//output_address := "31h4qstPuuYe26HAQpiWhCj4upQtSxDbsG"
	//kyt_api_get_user_url:"https://api.chainalysis.com/api/kyt/v1/users"
	//kyt_api_transfer_recevied:
	/*postData := map[string]interface{}{
		"asset" : "BTC",
		"address" : output_address,

	}*/
	client := &http.Client{}
	// get请求
	req, err := http.NewRequest("POST", kyt_v1_url, nil)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	// 在请求头中加入校验的token
	req.Header.Set("Token", token)
	//req.Header.Set("data",)
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
	fmt.Printf("%s\n", bodyText)

}
