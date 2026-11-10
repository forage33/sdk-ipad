# 04. 开放平台

> 源文件：`api_open_platform_*.go`（每接口一文件） · 接口数：5
>
> ← [API 总览](../README.md) · [文档索引](../../INDEX.md) · [架构文档](../../../ARCHITECTURE.md)

企业用户 id 与开放平台 id（external_userid）互转、群 id 与开放平台 chatid 互转。

## 接口列表

| # | 接口 | SDK 方法 | 封装状态 | 文档 |
|---|------|----------|----------|------|
| 4.1 | 企业用户id转开放平台id | `Vid2UserId` | ✅ 已封装 | [4.1-企业用户id转开放平台id.md](./4.1-企业用户id转开放平台id.md) |
| 4.2 | 开放平台id转企业用户id | `UserId2Vid` | ✅ 已封装 | [4.2-开放平台id转企业用户id.md](./4.2-开放平台id转企业用户id.md) |
| 4.3 | 企业用户id转开放平台id（代开发应用版本） | `Vid2OpenUserTripartiteId` | ✅ 已封装 | [4.3-企业用户id转开放平台id（代开发应用版本).md](./4.3-企业用户id转开放平台id（代开发应用版本).md) |
| 4.4 | 群id转chatid | `RoomIdToChatId` | ✅ 已封装 | [4.4-群id转chatid.md](./4.4-群id转chatid.md) |
| 4.5 | chatid转群id | `ChatIdToRoomId` | ✅ 已封装 | [4.5-chatid转群id.md](./4.5-chatid转群id.md) |

> 封装说明：`scorpID` 为空 / `corpID`、`appID` 为 0 时对应参数不下发（协议侧默认转当前账号登录企业）；返回统一为 `VidOpenIdPair` 或 `RoomChatIdPair`（见 `api_common_types.go`）。
