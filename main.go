package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"github.com/fatih/color"
)

func main() {
	// 命令类型
	isRegister := flag.Bool("register", false, "register command")
	isLogin := flag.Bool("login", false, "login command")

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

		fmt.Println(result.Message)
		// 把token写入本地
		ioutil.WriteFile("token", []byte(result.Data.(string)), 0644)

		fmt.Println("register")
		fmt.Printf("username: %s\n", *username)
		fmt.Printf("password: %s\n", *password)
	}

	if (*isLogin) {
		result := Login(*username, *password)
		fmt.Println(result.Message)
		// 把token写入本地
		ioutil.WriteFile("token", []byte(result.Data.(string)), 0644)

		fmt.Println("login")
		fmt.Printf("username: %s\n", *username)
		fmt.Printf("password: %s\n", *password)
	}

	if (len(*cmd) > 0) {
		result := Create(*cmd, *comment)
		fmt.Println(result.Message)
		fmt.Printf("command: %s\n", *cmd)
		fmt.Printf("comment: %s\n", *comment)
	}


	keyword := flag.Arg(0)

	if len(keyword) > 0 {
		result := Search(keyword)

		green := color.New(color.FgGreen).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()


		fmt.Println(result.Message)
		cmds := result.Data
		len := len(cmds)
		for i := 0; i < len; i++ {
			// fmt.Printf("%s  %s id:%s\n", yellow(cmds[i].Cmd), blue(cmds[i].Comment), red(cmds[i].Id))
			// 支持windows
			fmt.Fprintf(color.Output,"%s  %s id:%s\n", green(cmds[i].Cmd), blue(cmds[i].Comment), red(cmds[i].Id))
		}
	}

	if(*id > 0){
		fmt.Println(*id)
		result := Delete(*id)
		fmt.Println(result.Message)
	}

	//fmt.Printf("success!")
}
