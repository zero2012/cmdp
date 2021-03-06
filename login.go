package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"log"
	"encoding/json"
)

func Login(username string, password string) Respond  {
	result := Respond{}

	client := &http.Client{}

	req, err := http.NewRequest("POST", URL+"/user/login", strings.NewReader("Username="+username+"&Password="+password))
	if err != nil {
		log.Println(err)
		return result
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return result
	}

	json.Unmarshal(body,&result)
	return result
}
