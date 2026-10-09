# sdk-ipad 回调消息文档

> 协议服务经 `/callback` 推送的全部回调消息：按域分目录组织，每个回调事件一个独立文档（契约层字段名的权威出处）。
> 数据来源：ShowDoc (https://www.showdoc.com.cn/1889765213630885)
>
> ← [文档索引](../INDEX.md) · [架构文档](../ARCHITECTURE.md) · [模块 README](../../README.md)

| 域 | 契约层参考 | 回调数 | 说明 |
|----|------------|--------|------|
| [01. 登录回调](./01-登录回调/README.md) | `callback_payload_account.go` | 6 | 扫码/取消/确认/登录成功/企业切换/验证码回调。 |
| [02. 消息回调](./02-消息回调/README.md) | `callback_payload_msg.go` | 27 | 文本/GIF/小程序/链接/视频号/位置/名片/内外联系人媒体/红包/语音/通话/撤回/合并/系统/接龙等消息回调。 |
| [03. 群回调](./03-群回调/README.md) | `callback_payload_account.go` | 9 | 群新增/移除/邀请/退出/改名/转让/管理员/解散/群变动回调。 |
| [04. 联系人回调](./04-联系人回调/README.md) | `callback_payload_account.go` | 7 | 好友申请/联系人变动/标签变动/已读/删除/拉黑/置顶免打扰回调。 |
| [05. 控制台下发任务回调](./05-控制台下发任务回调/README.md) | `—` | 3 | 控制台下发的群发客户群/客户/朋友圈任务回调。 |
| [06. 标签变动回调](./06-标签变动回调/README.md) | `callback_payload_account.go` | 4 | 个人/企业标签新增、删除回调。 |
| [07. 聊天标签 & 断开回调](./07-聊天标签与断开回调/README.md) | `callback_payload_msg.go + callback_payload_account.go` | 7 | 聊天标签变动、标签联系人变动、二次验证与断开类回调。 |

> 拆分前的按域合并分册已归档至 [_archive/](./_archive/README.md)，仅作历史留档。
