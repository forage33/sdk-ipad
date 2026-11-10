> ← [返回文档索引](../../INDEX.md) | [架构文档](../../../ARCHITECTURE.md)
## 十一、控制台群发 (broadcast.go)

### 11.1 确认群发消息 ✅ `ConfirmGroupSend`
- **请求URL**: `/wxwork/oksendgroup`
- **参数**: uuid, vids([]long), msgid(long), isroom(int)

### 11.2 获取未发送群发列表 🔧 TODO `GetPendingGroupSendList`
- **请求URL**: 待确认
- **参数**: 待补充

### 11.3 客户群发送任务-获取要发送的客户群列表 🔧 TODO `GetGroupSendCustomerList`
- **请求URL**: 待确认
- **参数**: 待补充

### 11.4 获取群发任务详细信息 🔧 TODO `GetGroupSendTaskDetail`
- **请求URL**: 待确认
- **参数**: 待补充

### 11.5 朋友圈控制台任务执行 🔧 TODO `ExecSnsConsoleTask`
- **请求URL**: 待确认
- **参数**: 待补充

### 11.6 管理员获取企业成员客户列表 ❌ 未封装
- **请求URL**: 待确认
- **参数**: 待补充

---

## 十二、聊天标签 (chat_tag.go)

### 12.1 增删改聊天标签 ✅ `ModSessionTag`
- **请求URL**: `/wxwork/ModSessionTag`
- **参数**: uuid, cmd(int: 1=新增 2=修改 3=删除), tagName, tagID, globalVer

### 12.2 获取聊天标签列表 🔧 TODO `GetSessionTagList`
- **请求URL**: 待确认
- **参数**: 待补充

### 12.3 获取标签下的联系人 🔧 TODO `GetSessionTagMembers`
- **请求URL**: 待确认
- **参数**: 待补充

### 12.4 新增/移除联系人到聊天标签 🔧 TODO `ModSessionTagMembers`
- **请求URL**: 待确认
- **参数**: 待补充

---

