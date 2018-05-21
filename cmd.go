package main

import (
	"time"
)

// 所有命令,代码 都保存在这个表中
type Cmd struct {
	Id        int
	UserId    int
	Cmd       string
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
