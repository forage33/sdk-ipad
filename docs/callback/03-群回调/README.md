# 03. 群回调

> 契约层参考：`callback_payload_account.go` · 回调数：9
>
> ← [回调总览](../README.md) · [文档索引](../../INDEX.md) · [架构文档](../../ARCHITECTURE.md)

群新增/移除/邀请/退出/改名/转让/管理员/解散/群变动回调。

## 回调列表

| # | 回调 | 类型码 | SDK 常量 | 文档 |
|---|------|--------|----------|------|
| 3.1 | 群新增通知 | `115001` | `NewGroup` | [3.1-NewGroup.md](./3.1-NewGroup.md) |
| 3.2 | 移除群成员通知 | `115002` | `RemoveGroupMember` | [3.2-RemoveGroupMember.md](./3.2-RemoveGroupMember.md) |
| 3.3 | 邀请进群通知 | `115003` | `InviteInGroup` | [3.3-InviteInGroup.md](./3.3-InviteInGroup.md) |
| 3.4 | 自己主动退出群聊 | `115004` | `QuitGroup` | [3.4-QuitGroup.md](./3.4-QuitGroup.md) |
| 3.5 | 群名变更通知 | `115005` | `GroupNameChange` | [3.5-GroupNameChange.md](./3.5-GroupNameChange.md) |
| 3.6 | 群主转让 等其他通知 | `115006` | `GroupSuperAdminChange` | [3.6-GroupSuperAdminChange.md](./3.6-GroupSuperAdminChange.md) |
| 3.7 | 群管理员变动通知 | `115007` | `GroupAdminChange` | [3.7-GroupAdminChange.md](./3.7-GroupAdminChange.md) |
| 3.8 | 群解散通知 | `115008` | `GroupDisBand` | [3.8-GroupDisBand.md](./3.8-GroupDisBand.md) |
| 3.9 | 群变动通知 | `115009` | `GroupMemberChange` | [3.9-GroupMemberChange.md](./3.9-GroupMemberChange.md) |
