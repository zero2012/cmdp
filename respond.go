package main

type Respond struct {
	Status  string
	Message  string
	Data interface{}
}

type CmdsRespond struct {
	Status  string
	Message  string
	Data []Cmd
}