# 04. 联系人回调

> 契约层参考：`callback_payload_account.go` · 回调数：7
>
> ← [回调总览](../README.md) · [文档索引](../../INDEX.md) · [架构文档](../../ARCHITECTURE.md)

好友申请/联系人变动/标签变动/已读/删除/拉黑/置顶免打扰回调。

## 回调列表

| # | 回调 | 类型码 | SDK 常量 | 文档 |
|---|------|--------|----------|------|
| 4.1 | 联系人添加申请通知 | `116001` | `ApplyFriend` | [4.1-ApplyFriend.md](./4.1-ApplyFriend.md) |
| 4.2 | 联系人变动通知 | `116002` | `FriendChange` | [4.2-FriendChange.md](./4.2-FriendChange.md) |
| 4.3 | 内部联系人标签描述变动通知 | `116003` | `InnerLabelChange` | [4.3-InnerLabelChange.md](./4.3-InnerLabelChange.md) |
| 4.4 | 当前联系人已读消息通知回调 | `102006` | `MsgRead` | [4.4-MsgRead.md](./4.4-MsgRead.md) |
| 4.5 | 联系人发消息提示删除通知 | `102001` | `HasDeleted` | [4.5-HasDeleted.md](./4.5-HasDeleted.md) |
| 4.6 | 联系人发消息提示拉黑通知 | `102002` | `HasBlacked` | [4.6-HasBlacked.md](./4.6-HasBlacked.md) |
| 4.7 | 置顶免打扰标记操作通知 | `116005` | `PinnedMsg` | [4.7-PinnedMsg.md](./4.7-PinnedMsg.md) |
