package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	file, err := http.Get("https://s3.ap-southeast-1.amazonaws.com/utrading.io/utrading/install/android/install_google_v2.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Body.Close()
	body, err := ioutil.ReadAll(file.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
