package main

import (
	"fmt"
	"net/http"
) 

func get(test, url string, expectedCode int){
	resp, err := http.Get(url)
	switch {
	case err != nil:
		fmt.Printf("[ERROR] : %s failed with %s : %s\n", test, url, err)
	case resp.StatusCode != expectedCode:
		fmt.Printf("[ERROR] : %s : %s has response code of %d\n", test, url, resp.StatusCode)
	default:
		fmt.Printf("[PASSED] : %s ok with %s\n", test, url)
		//fmt.Println(resp.Body)	
	}
	defer resp.Body.Close()
}

func main(){
	host := "localhost"
	port := 2020
	baseUrl := fmt.Sprintf("http://%s:%d", host, port)

	get("listing tags", baseUrl + "/v1/tags", 200)
	get("listing entries", baseUrl + "/v1/entries", 200)
	get("listing known entry", baseUrl + "/v1/entries/id1", 200)
	get("listing unknown entry", baseUrl + "/v1/entries/idx", 404)
}
