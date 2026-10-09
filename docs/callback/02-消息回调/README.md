# 02. 消息回调

> 契约层参考：`callback_payload_msg.go` · 回调数：27
>
> ← [回调总览](../README.md) · [文档索引](../../INDEX.md) · [架构文档](../../ARCHITECTURE.md)

文本/GIF/小程序/链接/视频号/位置/名片/内外联系人媒体/红包/语音/通话/撤回/合并/系统/接龙等消息回调。

## 回调列表

| # | 回调 | 类型码 | msgtype | SDK 常量 | 文档 |
|---|------|--------|---------|----------|------|
| 2.1 | 文本消息接收 | `102000` | 2 | `ChatMsgText` | [2.1-ChatMsgText.md](./2.1-ChatMsgText.md) |
| 2.2 | gif表情消息接收 | `102000` | 104 | `ChatMsgGif` | [2.2-ChatMsgGif.md](./2.2-ChatMsgGif.md) |
| 2.3 | 小程序消息 | `102000` | 78 | `ChatMsgMiniProgram` | [2.3-ChatMsgMiniProgram.md](./2.3-ChatMsgMiniProgram.md) |
| 2.4 | 链接消息 | `102000` | 13 | `ChatMsgLink` | [2.4-ChatMsgLink.md](./2.4-ChatMsgLink.md) |
| 2.5 | 视频号消息 | `102000` | 141 | `ChatMsgVideoChannel` | [2.5-ChatMsgVideoChannel.md](./2.5-ChatMsgVideoChannel.md) |
| 2.6 | 名片消息 | `102000` | 41 | `ChatMsgCard` | [2.6-ChatMsgCard.md](./2.6-ChatMsgCard.md) |
| 2.7 | 位置消息 | `102000` | 6 | `ChatMsgLocation` | [2.7-ChatMsgLocation.md](./2.7-ChatMsgLocation.md) |
| 2.8 | 外部联系人图片消息（个v） | `102000` | 101 | `ChatMsgExternalImage` | [2.8-ChatMsgExternalImage.md](./2.8-ChatMsgExternalImage.md) |
| 2.9 | 外部联系人文件消息 | `102000` | 102 | `ChatMsgExternalFile` | [2.9-ChatMsgExternalFile.md](./2.9-ChatMsgExternalFile.md) |
| 2.10 | 外部联系人视频消息 | `102000` | 103 | `ChatMsgExternalVideo` | [2.10-ChatMsgExternalVideo.md](./2.10-ChatMsgExternalVideo.md) |
| 2.11 | 内部联系人图片消息（企业微信） | `102000` | 14 | `ChatMsgInternalImage` | [2.11-ChatMsgInternalImage.md](./2.11-ChatMsgInternalImage.md) |
| 2.12 | 内部联系人文件消息 | `102000` | 15 | `ChatMsgInternalFile` | [2.12-ChatMsgInternalFile.md](./2.12-ChatMsgInternalFile.md) |
| 2.12b | 内部联系人视频消息 | `102000` | 23 | `ChatMsgInternalVideo` | [2.12b-ChatMsgInternalVideo.md](./2.12b-ChatMsgInternalVideo.md) |
| 2.13 | 红包消息 | `102000` | 26 | `ChatMsgRedPacket` | [2.13-ChatMsgRedPacket.md](./2.13-ChatMsgRedPacket.md) |
| 2.14 | 语音消息 | `102000` | 16 | `ChatMsgVoice` | [2.14-ChatMsgVoice.md](./2.14-ChatMsgVoice.md) |
| 2.15 | 语音通话或者视频通话消息 | `102000` | 503 | `ChatMsgVoiceVideoCall` | [2.15-ChatMsgVoiceVideoCall.md](./2.15-ChatMsgVoiceVideoCall.md) |
| 2.16 | 语音或者视频通话，未接听消息 | `102000` | 503 | — | [2.16-ChatMsgVoiceVideoCallMissed.md](./2.16-ChatMsgVoiceVideoCallMissed.md) |
| 2.17 | 语音或者视频通话，在其他端接听消息 | `102000` | 503 | — | [2.17-ChatMsgVoiceVideoCallAnswered.md](./2.17-ChatMsgVoiceVideoCallAnswered.md) |
| 2.18 | 语音或者视频通话消息，对方取消通话 | `102000` | 40 | `ChatMsgCallCancel` | [2.18-ChatMsgCallCancel.md](./2.18-ChatMsgCallCancel.md) |
| 2.19 | 语音视频通话时长消息通知 | `102000` | 40 | — | [2.19-ChatMsgCallDuration.md](./2.19-ChatMsgCallDuration.md) |
| 2.20 | 撤回消息 | `102000` | 2063 | `ChatMsgRevoke` | [2.20-ChatMsgRevoke.md](./2.20-ChatMsgRevoke.md) |
| 2.21 | 合并消息 | `102000` | 4 | `ChatMsgMerge` | [2.21-ChatMsgMerge.md](./2.21-ChatMsgMerge.md) |
| 2.22 | 系统消息(发送失败封禁功能等提醒) | `102000` | 1011 | `ChatMsgSystem` | [2.22-ChatMsgSystem.md](./2.22-ChatMsgSystem.md) |
| 2.23 | 接龙消息接收 | `102000` | 213 | `ChatMsgChainRelay` | [2.23-ChatMsgChainRelay.md](./2.23-ChatMsgChainRelay.md) |
| 2.24 | 视频号直播消息 | `102000` | 146 | — | [2.24-ChatMsgVideoChannelLive.md](./2.24-ChatMsgVideoChannelLive.md) |
| 2.25 | 内部联系人大文件消息 | `102000` | 20 | — | [2.25-ChatMsgInternalBigFile.md](./2.25-ChatMsgInternalBigFile.md) |
| 2.26 | 新增客服聊天排队通知 | `102000` | 2314 | — | [2.26-ChatMsgCsQueueNotify.md](./2.26-ChatMsgCsQueueNotify.md) |
