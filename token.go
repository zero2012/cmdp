package main

import (
	"runtime"
	"io/ioutil"
	"fmt"
)

func CreateToken(token interface{}) {
	GoRoot := runtime.GOROOT()
	if runtime.GOOS == "windows" {
		fmt.Println("create cmdp_token at :" + GoRoot + "\\cmdp_token")
		err := ioutil.WriteFile(GoRoot+"\\cmdp_token", []byte(token.(string)), 0644)
		if err!=nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("create cmdp_token at :" + GoRoot + "/cmdp_token")
		err := ioutil.WriteFile(GoRoot+"/cmdp_token", []byte(token.(string)), 0644)
		if err!=nil {
			fmt.Println(err)
		}
	}
}

func ReadToken() ([]byte, error) {
	GoRoot := runtime.GOROOT()
	if runtime.GOOS == "windows" {
		token, err := ioutil.ReadFile(GoRoot + "\\cmdp_token")
		if err!=nil {
			fmt.Println(err)
		}
		return token, err
	} else {
		token, err := ioutil.ReadFile(GoRoot + "/cmdp_token")
		if err!=nil {
			fmt.Println(err)
		}
		return token, err
	}
}
