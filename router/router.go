package router

import (
    "github.com/zzzhangjian/go_peatio/app/api/hello"
    "github.com/gogf/gf/g"
)

// 统一路由注册.
func init() {
    g.Server().BindHandler("/", hello.Handler)
}
