package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	"log"
	"strconv"
)

func Create(cmd string, comment string) Respond  {
	result := Respond{}

	client := &http.Client{}

	req, err := http.NewRequest("POST", URL+"/cmdp/create", strings.NewReader("Cmd="+cmd+"&Comment="+comment))
	if err != nil {
		log.Println(err)
		return result
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	token, _ := ioutil.ReadFile("token")//从文件取出数据
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

func Search(keyword string ) CmdsRespond  {
	result := CmdsRespond{}

	client := &http.Client{}

	req, err := http.NewRequest("POST", URL+"/cmdp/search", strings.NewReader("keyword="+keyword))
	if err != nil {
		log.Println(err)
		return result
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	token, _ := ioutil.ReadFile("token")//从文件取出数据
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

func Delete(id int ) CmdsRespond  {
	result := CmdsRespond{}

	client := &http.Client{}

	req, err := http.NewRequest("GET", URL+"/cmdp/delete/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Println(err)
		return result
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	token, _ := ioutil.ReadFile("token")//从文件取出数据
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

