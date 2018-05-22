package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"github.com/fatih/color"
	"strconv"
)

func main() {
	// 命令类型
	isRegister := flag.Bool("register", false, "register command")
	isLogin := flag.Bool("login", false, "login command")
	isReset := flag.Bool("reset", false, "reset password command")

	// 注册/登录用户
	username := flag.String("u", "", "register by username")
	password := flag.String("p", "", "your password")

	// 创建命令
	cmd := flag.String("c", "", "save your command prompt to web")
	comment := flag.String("m", "", "comment your command prompt")

	// 删除指定id的命令
	id := flag.Int("d", 0, "delete cmd by id")

	flag.Parse()

	if (*isRegister) {
		result := Register(*username, *password)
		// 把token写入本地
		if(result.Status!=SUCCESS){
			return
		}
		ioutil.WriteFile("token", []byte(result.Data.(string)), 0644)

		fmt.Println("register")
		printRespond(result)
	}

	if (*isLogin) {
		result := Login(*username, *password)
		// 把token写入本地
		if(result.Status!=SUCCESS){
			return
		}
		ioutil.WriteFile("token", []byte(result.Data.(string)), 0644)

		fmt.Println("login")
		printRespond(result)
	}


	if (*isReset) {
		result := ResetPassword(*password)
		// 把token写入本地
		if(result.Status!=SUCCESS){
			return
		}
		ioutil.WriteFile("token", []byte(result.Data.(string)), 0644)

		fmt.Println("reset password")
		printRespond(result)
	}

	if (len(*cmd) > 0) {
		result := Create(*cmd, *comment)
		fmt.Printf("command: %s\n", *cmd)
		fmt.Printf("comment: %s\n", *comment)
		printRespond(result)
	}


	keyword := flag.Arg(0)

	if len(keyword) > 0 {
		result := Search(keyword)

		green := color.New(color.FgGreen).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()

		cmds := result.Data
		len := len(cmds)
		for i := 0; i < len; i++ {
			// fmt.Printf("%s  %s id:%s\n", yellow(cmds[i].Cmd), blue(cmds[i].Comment), red(cmds[i].Id))
			// 支持windows
			fmt.Fprintf(color.Output,"%s  %s id:%s\n", green(cmds[i].Cmd), blue(cmds[i].Comment), red(cmds[i].Id))
		}
		if(result.Status == SUCCESS){
			color.Blue("results: "+strconv.Itoa(len))
		}else{
			color.Red(result.Message)
		}
	}

	if(*id > 0){
		result := Delete(*id)
		if(result.Status == SUCCESS){
			color.Green(result.Message)
		}else{
			color.Red(result.Message)
		}
	}

	//fmt.Printf("success!")
}

func printRespond(result Respond)  {
	if(result.Status == SUCCESS){
		color.Green(result.Message)
	}else{
		color.Red(result.Message)
	}
}
