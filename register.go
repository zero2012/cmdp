package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"log"
	"encoding/json"
)

func Register(username string, password string) Respond  {
	result := Respond{}

	client := &http.Client{}

	req, err := http.NewRequest("POST", URL+"/user/register", strings.NewReader("Username="+username+"&Password="+password))
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

func ResetPassword(password string) Respond  {
	result := Respond{}

	client := &http.Client{}

	req, err := http.NewRequest("POST", URL+"/user/reset", strings.NewReader("password="+password))
	if err != nil {
		log.Println(err)
		return result
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	token, _ := ReadToken()//从文件取出数据
	req.Header.Set("Authorization", "Bearer "+string(token))


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
