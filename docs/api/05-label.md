> ← [返回文档索引](../INDEX.md) | [架构文档](../ARCHITECTURE.md)
## 五、开放平台

> 文档页面未成功抓取，具体接口待确认

---

## 六、标签 (label.go)

### 6.1 获取标签列表 ✅ `GetLabelList`

- **请求URL**: `/wxwork/GetLabelListReq`
- **请求方式**: POST
- **参数**:

| 参数名 | 必选 | 类型 | 说明 |
|--------|------|------|------|
| uuid | 是 | String | 实例唯一标识 |
| index | 是 | int | 分页索引 |
| sync_type | 是 | int | 同步类型：1=企业标签, 2=个人标签 |

### 6.2 添加标签 ✅ `AddLabel`

- **请求URL**: `/wxwork/AddLabelReq`
- **参数**:

| 参数名 | 必选 | 类型 | 说明 |
|--------|------|------|------|
| uuid | 是 | String | 实例唯一标识 |
| labelName | 是 | String | 标签名称 |
| data_type | 是 | int | 数据类型 |
| label_type | 是 | int | 标签类型 |
| groupid | 是 | long | 标签组ID |

### 6.3 修改标签 ✅ `EditLabel`

- **请求URL**: `/wxwork/EditLabelReq`
- **参数**:

| 参数名 | 必选 | 类型 | 说明 |
|--------|------|------|------|
| uuid | 是 | String | 实例唯一标识 |
| labelName | 是 | String | 新标签名称 |
| labelid | 是 | long | 标签ID |
| groupid | 是 | long | 标签组ID |

### 6.4 删除标签 ✅ `DelLabel`

- **请求URL**: `/wxwork/DelLabelReq`
- **参数**:

| 参数名 | 必选 | 类型 | 说明 |
|--------|------|------|------|
| uuid | 是 | String | 实例唯一标识 |
| labelid | 是 | long | 标签ID |
| groupid | 是 | long | 标签组ID |

### 6.5 为标签添加用户 ✅ `LabelAddUsers`

- **请求URL**: `/wxwork/LabelAddUserReq`
- **参数**:

| 参数名 | 必选 | 类型 | 说明 |
|--------|------|------|------|
| uuid | 是 | String | 实例唯一标识 |
| labelid | 是 | long | 标签ID |
| userid_list | 是 | []long | 要添加的用户ID列表 |
| del_userid | 是 | []long | 要删除的用户ID列表 |

### 6.6 为用户添加标签 ✅ `UserAddLabels`

- **请求URL**: `/wxwork/UserAddLabelsReq`
- **参数**:

| 参数名 | 必选 | 类型 | 说明 |
|--------|------|------|------|
| uuid | 是 | String | 实例唯一标识 |
| userid | 是 | long | 用户ID |
| labelid_list | 是 | []long | 标签ID列表 |

---

