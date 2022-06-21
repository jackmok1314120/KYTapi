package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	token := "8e283fdd013173ce7712d8e5ae9e9d21655274956bddbb40388a4a9229b135a5"
	//hoo_userId := "hoo_test01"
	kyt_url := "https://api.chainalysis.com/api/kyt/v1/users/hoo_test01//withdrawaladdresses"
	//withdr := "/withdrawaladdresses"
	//kyt_asset := "BTC"
	//output_address := "31h1yc8MnKj49By5tSjApsaLUnEeJ1g4u4"
	req := url.Values{"asset": {"BTC"}, "address": {"31h1yc8MnKj49By5tSjApsaLUnEeJ1g4u4"}}
	//body := json.Marshaler(req)

	resp, err := http.PostForm(kyt_url, req)
	if err != nil {
		// handle error
	}
	resp.Header.Set("Token", token)
	//resp.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp.Header.Set("Accept", "application/json")
	resp.Header.Set("Content-Type", "application/json")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println("status", resp.Status)
	fmt.Println("response:", resp.Header)
	fmt.Println(string(body))

}
