package main

import (
	_ "github.com/zzzhangjian/go_peatio/boot"
	_ "github.com/zzzhangjian/go_peatio/router"
	"github.com/gogf/gf/g"
)

func main() {
	g.Server().Run()
}
