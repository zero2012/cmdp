package main

import (
	"runtime"
	"io/ioutil"
)

func CreateToken(token interface{}) {
	GoRoot := runtime.GOROOT()
	if runtime.GOOS == "windows" {
		ioutil.WriteFile(GoRoot+"\\token", []byte(token.(string)), 0644)
	} else {
		ioutil.WriteFile(GoRoot+"/token", []byte(token.(string)), 0644)
	}
}

func ReadToken() ([]byte, error) {
	GoRoot := runtime.GOROOT()
	if runtime.GOOS == "windows" {
		return ioutil.ReadFile(GoRoot + "\\token")
	} else {
		return ioutil.ReadFile(GoRoot + "/token")
	}
}
