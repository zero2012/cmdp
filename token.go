package main

import (
	"runtime"
	"io/ioutil"
)

func CreateToken(token interface{}) {
	GoRoot := runtime.GOROOT()
	if runtime.GOOS == "windows" {
		ioutil.WriteFile(GoRoot+"\\cmdp_token", []byte(token.(string)), 0644)
	} else {
		ioutil.WriteFile(GoRoot + "/cmdp_token", []byte(token.(string)), 0644)
	}
}

func ReadToken() ([]byte, error) {
	GoRoot := runtime.GOROOT()
	if runtime.GOOS == "windows" {
		return ioutil.ReadFile(GoRoot + "\\cmdp_token")
	} else {
		return ioutil.ReadFile(GoRoot + "/cmdp_token")
	}
}
