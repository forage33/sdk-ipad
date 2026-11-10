// sdk-ipad：企业微信 iPad 协议 SDK（独立 Go 模块）。
//
// 仅依赖标准库、gjson 与 zap，可独立编译、测试、打 tag 发布。
// 宿主工程通过 go.mod 的 require 引用本模块（联调期可用 replace 指向本地副本）。
module github.com/forage33/sdk-ipad

go 1.25.0

require (
	github.com/tidwall/gjson v1.17.3
	go.uber.org/zap v1.27.0
)

require (
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
)
