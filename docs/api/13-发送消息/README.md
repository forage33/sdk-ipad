# 13. 发送消息

> 源文件：`message.go` · 接口数：20
>
> ← [API 总览](../README.md) · [文档索引](../../INDEX.md) · [架构文档](../../ARCHITECTURE.md)

文本/图片/文件/语音/视频/链接/小程序/名片/位置/表情/@消息/群发/撤回/已读等消息接口。

## 接口列表

| # | 接口 | SDK 方法 | 封装状态 | 文档 |
|---|------|----------|----------|------|
| 13.1 | 撤回消息 | `RevokeMsg` | ✅ 已封装 | [13.1-RevokeMsg.md](./13.1-RevokeMsg.md) |
| 13.2 | 发送文本消息 | `SendTextMsg` | ✅ 已封装 | [13.2-SendTextMsg.md](./13.2-SendTextMsg.md) |
| 13.3 | 发送CDN图片消息 | `SendCDNImage` | ✅ 已封装 | [13.3-SendCDNImage.md](./13.3-SendCDNImage.md) |
| 13.4 | 发送CDN文件消息 | `SendCDNFile` | ✅ 已封装 | [13.4-SendCDNFile.md](./13.4-SendCDNFile.md) |
| 13.5 | 发送CDN语音消息 | `SendCDNVoiceMsg` | ✅ 已封装 | [13.5-SendCDNVoiceMsg.md](./13.5-SendCDNVoiceMsg.md) |
| 13.6 | 发送CDN视频 | `SendCDNVideoMsg` | ✅ 已封装 | [13.6-SendCDNVideoMsg.md](./13.6-SendCDNVideoMsg.md) |
| 13.7 | 发送大视频消息 | `SendCDNBigVideoMsg` | ✅ 已封装 | [13.7-SendCDNBigVideoMsg.md](./13.7-SendCDNBigVideoMsg.md) |
| 13.8 | 发送大文件消息 | `SendCDNBigFileMsg` | ✅ 已封装 | [13.8-SendCDNBigFileMsg.md](./13.8-SendCDNBigFileMsg.md) |
| 13.9 | 发送链接卡片消息 | `SendLinkMsg` | ✅ 已封装 | [13.9-SendLinkMsg.md](./13.9-SendLinkMsg.md) |
| 13.10 | 发送小程序 | `SendAppMsg` | ✅ 已封装 | [13.10-SendAppMsg.md](./13.10-SendAppMsg.md) |
| 13.10b | 发送视频号消息 | `SendVideoNumber` | ✅ 已封装 | [13.10b-SendVideoNumber.md](./13.10b-SendVideoNumber.md) |
| 13.11 | 发送视频号直播消息 | `SendVideoNumberZhiBo` | ✅ 已封装 | [13.11-SendVideoNumberZhiBo.md](./13.11-SendVideoNumberZhiBo.md) |
| 13.12 | 发送gif表情 | `SendEmotionMessage` | ✅ 已封装 | [13.12-SendEmotionMessage.md](./13.12-SendEmotionMessage.md) |
| 13.13 | 发送名片消息 | `SendBusinessCardMsg` | ✅ 已封装 | [13.13-SendBusinessCardMsg.md](./13.13-SendBusinessCardMsg.md) |
| 13.14 | 发送位置消息 | `SendLocationMsg` | ✅ 已封装 | [13.14-SendLocationMsg.md](./13.14-SendLocationMsg.md) |
| 13.15 | 发送群@消息 | `SendTextAtMsg` | ✅ 已封装 | [13.15-SendTextAtMsg.md](./13.15-SendTextAtMsg.md) |
| 13.16 | 同步消息记录 | `SyncAllData` | ✅ 已封装 | [13.16-SyncAllData.md](./13.16-SyncAllData.md) |
| 13.17 | 群发消息 | `SendGroupsMsg` | ✅ 已封装 | [13.17-SendGroupsMsg.md](./13.17-SendGroupsMsg.md) |
| 13.18 | 语音转文字 | `SpeechToTextEntity` | ✅ 已封装 | [13.18-SpeechToTextEntity.md](./13.18-SpeechToTextEntity.md) |
| 13.19 | 消息已读 | `MarkAsRead` | ✅ 已封装 | [13.19-MarkAsRead.md](./13.19-MarkAsRead.md) |
