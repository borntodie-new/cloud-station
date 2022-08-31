package main

import (
	"fmt"
	"gitee.com/go-course/cloud-station-g7/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println("执行出错：", err.Error())
	}
}
