package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// token:KYT API Key
	token := "8e283fdd013173ce7712d8e5ae9e9d21655274956bddbb40388a4a9229b135a5"
	hoo_userId := "hoo_test01" //hoo.com user ID.
	kyt_url_1 := "https://api.chainalysis.com/api/kyt/v1/users/" + hoo_userId + "/transfers/received"
	//kyt_api_get_user_url:"https://api.chainalysis.com/api/kyt/v1/users"
	//kyt_api_transfer_recevied:
	client := &http.Client{}
	// get请求
	req, err := http.NewRequest("GET", kyt_url_1, nil)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	// 在请求头中加入校验的token
	req.Header.Set("Token", token)
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

}
