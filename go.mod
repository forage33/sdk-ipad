// sdk-ipad 独立 Go 子模块：企业微信 iPad 协议 SDK。
//
// 位于仓库主目录（与 ocean-scrm-server 平级），与主工程解耦，仅依赖标准库与 zap，
// 可独立编译、测试、打 tag 发布。主工程通过 ocean-scrm-server/go.mod 的
// require + replace（指向 ../sdk-ipad）引用本模块，import 路径保持
// ocean-scrm-server/sdk-ipad 不变。
module ocean-scrm-server/sdk-ipad

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
