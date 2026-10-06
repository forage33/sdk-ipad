> ← [返回文档索引](../INDEX.md) | [架构文档](../ARCHITECTURE.md)
## 八、联系人操作 (contact.go)

### 8.1 获取内部联系人列表 [已封装: GetInnerContacts]

- **请求URL**: `/wxwork/GetInnerContacts`
- **参数**:

| 参数名 | 必选 | 类型 | 说明 |
|--------|------|------|------|
| uuid | 是 | String | 实例唯一标识 |
| limit | 否 | int | 每次返回数量 |
| strSeq | 否 | String | 查询下标，填返回的seq进行翻页 |

### 8.2 获取外部联系人列表 [已封装: GetExternalContacts]

- **请求URL**: `/wxwork/GetExternalContacts`
- **参数**:

| 参数名 | 必选 | 类型 | 说明 |
|--------|------|------|------|
| uuid | 是 | String | 实例唯一标识 |
| limit | 否 | int | 每次返回数量 |
| strSeq | 否 | String | 查询下标 |

### 8.3 根据用户id批量获取详细信息 [已封装: GetUserInfoByVids]

- **请求URL**: `/wxwork/GetUserInfoByVids`
- **参数**:

| 参数名 | 必选 | 类型 | 说明 |
|--------|------|------|------|
| uuid | 是 | String | 实例唯一标识 |
| vids | 是 | []long | 用户ID列表 |

### 8.4 同意好友 [已封装: AgreeUser]


##### 简要描述

- 同意好友

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/AgreeUser
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |

##### 请求示例 

``` 
{
    "uuid":"3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "corpid":0,
    "vid":7881302555913738 //在回调消息中拿到申请人id
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.5 删除联系人 [已封装: DeleteContact]
##### 简要描述

- 删除联系人

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/DelContact
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |

##### 请求示例 

``` 
{
    "uuid":"3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "vid":7881302555913738
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.6 获取好友申请列表 [已封装: GetApplyUserList]
##### 简要描述

- 获取好友申请列表

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetApplyUserList
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |

##### 请求示例 

``` 
{
    "uuid":"3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "limit":100, 
    "seq":0 //页码第一次传0 后续查询传返回的seq
}
```

##### 返回示例 

``` 
{
    "data": {
        "list": [
            {
                "item_flag": 2,
                "unionid": "ozynqstMEN_38dTCLJIlwe73ZhB0",
                "create_time": 0,
                "sex": 1,
                "mobile": "",
                "company_remark": "",
                "acctid": "",
                "avatar": "http://wx.qlogo.cn/mmhead/WRSOYLaeNjHy6nWdWhbrLHZSfYcZWeXSqyWVwACZDWM/0",
                "english_name": "",
                "remark_phone": [],
                "realname": "",
                "apply_create_time": 1743647253,
                "real_remarks": "",
                "user_id": 7881302814222765,
                "apply_update_time": 1743647570,
                "nickname": "班长",
                "position": "",
                "state": 2,
                "corp_id": 1970325134026788,
                "source_info": {
                    "vid": 0,
                    "wx_friend_name": "",
                    "apply_mode": 2,
                    "mobile": "",
                    "source_roomid": 0,
                    "source_type": 54
                },
                "remarks": "",
                "seq": 7649727,
                "status": 1
            },
            {
                "item_flag": 3,//3 代表删除 申请列表里面不展示的那种，两种删除一种是左划删除 一种是 删除了他的好友
                "unionid": "ozynqskXWQK3wAWWKT_izAOfcIlY",
                "create_time": 0,
                "sex": 1,
                "mobile": "",
                "company_remark": "",
                "acctid": "",
                "avatar": "http://wx.qlogo.cn/mmhead/Q3auHgzwzM7BQVtj83EnfQcDMjz2hhrIvssZl9fv1H6rYUaN8VRibjg/0",
                "english_name": "",
                "remark_phone": [],
                "realname": "",
                "apply_create_time": 1743652237, //申请添加时间
                "real_remarks": "",
                "user_id": 7881302555913738,
                "apply_update_time": 1743661373,
                "nickname": "十年一刻",
                "position": "",
                "state": 2,
                "corp_id": 1970325134026788,
                "source_info": {
                    "vid": 0,
                    "wx_friend_name": "",
                    "apply_mode": 0,
                    "mobile": "",
                    "source_roomid": 0,
                    "source_type": 0
                },
                "remarks": "",
                "seq": 7649739,
                "status": 1
            }
        ],
        "seq": 7649739
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.7 获取置顶免打扰折叠标记列表 [已封装: GetManagedConversations]

##### 简要描述

- 获取置顶免打扰折叠标记列表

##### 请求URL
- ` http://127.0.0.1:8083/wxwork/getManagedConversations
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
##### 请求示例 

``` 
{
    "uuid": "c4bb91fa86e8739eb6e2b7431fa55ee7"
}
```

##### 返回示例 

``` 
{
    "data": {
        "top_list": [  //置顶列表
            {
                "messagetype": 0, //用户类型  0好友 1群聊 3app类型 6开放平台类型
                "id": 7881302555913738,
                "fw_id": 0,
                "state": 0
            }
        ],
        "shield_list": [ //免打扰列表 
            {
                "messagetype": 7,
                "id": 10059,
                "fw_id": 0,
                "state": 1 //1是未折叠
            },
            {
                "messagetype": 0,
                "id": 7881302555913738,
                "fw_id": 0,
                "state": 2
            },
            {
                "messagetype": 0,
                "id": 7881299726922560,
                "fw_id": 0,
                "state": 2 //折叠聊天
            },
            {
                "messagetype": 3,
                "id": 10031,
                "fw_id": 0,
                "state": 0
            }
        ],
        "mark_session_list": [ //标记列表
            {
                "messagetype": 0,
                "id": 7881302555913738,
                "fw_id": 0,
                "state": 0
            },
            {
                "messagetype": 0,
                "id": 7881299726922560,
                "fw_id": 0,
                "state": 0
            }
        ]
    },
    "errcode": 0,
    "errmsg": "ok"
}
```


### 8.8 获取会话列表 [已封装: GetSessionList]


##### 简要描述

- 获取会话列表

##### 请求URL
- ` http://127.0.0.1:8083//wxwork/GetSessionList
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
##### 请求示例 

``` 
{
    "uuid":"4b72858b-3b82-4439-8fd7-a4485c2bfd99",
    "limit":100, 
    "star_index":0 
}
```

##### 返回示例 

``` 
{
    "data": {
        "top_list": [  //置顶列表
            {
                "messagetype": 3,
                "id": 10067
            }
        ],
        "collect_list": [],
        "shield_list": [ //免打扰列表
            {
                "messagetype": 7,
                "id": 10059
            },
            {
                "messagetype": 3,
                "id": 10031
            }
        ],
        "seq": 0,
        "room_list": [ //会话列表
            {
                "unreadcnt": 0, //未读消息数量 小红点
                "sessionid": 10004, //用户id
                "beginmsgseq": 1006452,
                "msgtype": 3 //用户类型  0好友 1群聊 3app类型 6开放平台类型
            },
            {
                "unreadcnt": 1,
                "sessionid": 10024,//学习园地
                "beginmsgseq": 0,
                "msgtype": 3
            },
            {
                "unreadcnt": 0,
                "sessionid": 10049,
                "beginmsgseq": 0,
                "msgtype": 3
            },
            {
                "unreadcnt": 0,
                "sessionid": 10060,
                "beginmsgseq": 1006492,
                "msgtype": 3
            },
            {
                "unreadcnt": 0,
                "sessionid": 10067, //客户联系
                "beginmsgseq": 1006641,
                "msgtype": 3
            },
            {
                "unreadcnt": 0,
                "sessionid": 10097,
                "beginmsgseq": 0,
                "msgtype": 3
            },
            {
                "unreadcnt": 0,
                "sessionid": 10120,//行业资讯
                "beginmsgseq": 1006552,
                "msgtype": 3
            },
            {
                "unreadcnt": 0,
                "sessionid": 10124,
                "beginmsgseq": 1006554,
                "msgtype": 3
            },
            {
                "unreadcnt": 0,
                "sessionid": 10167,
                "beginmsgseq": 1006001,
                "msgtype": 3
            },
            {
                "unreadcnt": 0,
                "sessionid": 1688852792312821, //企业微信团队
                "beginmsgseq": 1006634,
                "msgtype": 0
            },
            {
                "unreadcnt": 0,
                "sessionid": 1688854256622049,
                "beginmsgseq": 1006253,
                "msgtype": 0
            },
            {
                "unreadcnt": 0,
                "sessionid": 1688856854669094,
                "beginmsgseq": 1006855,
                "msgtype": 0
            },
            {
                "unreadcnt": 0,
                "sessionid": 5629502477624082,
                "beginmsgseq": 1006556,
                "msgtype": 6
            },
            {
                "unreadcnt": 0,
                "sessionid": 7881302059975581,
                "beginmsgseq": 1006863,
                "msgtype": 0
            },
            {
                "unreadcnt": 0,
                "sessionid": 7881302555913738,
                "beginmsgseq": 1006802,
                "msgtype": 0
            },
            {
                "unreadcnt": 0,
                "sessionid": 10884170884206879,
                "beginmsgseq": 1006075,
                "msgtype": 1
            },
            {
                "unreadcnt": 0,
                "sessionid": 10935946535750967,
                "beginmsgseq": 1006077,
                "msgtype": 1
            }
        ]
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.9 修改联系人备注手机简介企业 [已封装: UpdateSetMyUserInfo]


##### 简要描述

- 修改联系人备注手机简介企业

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/UpdateSetMyUserInfo
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |
##### 请求示例 

``` 
{
    "uuid":"222222222222222222",
    "vid":7881302555913738,
    "remark":"备注啦啦啦",
    "compny":"你好企业",
	"labelid_list": [
        "14073750634828569",
        "14073750634828609"
    ],
    "phone":[
        "13533333"
    ],
    "des":"你好简介"
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.10 内部联系人备注修改 [已封装: SetColleagueRemark]

##### 简要描述

- 内部联系人备注修改

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SetColleagueRemark
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |
##### 请求示例 

``` 
{
    "uuid":"1688855749266556",
    "vid":1688856554448765,
    "remark":"备注啦啦啦22222",
    "des":"你好简介"
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.11 获取内部联系人备注列表 [已封装: GetColleaRemarkList]
##### 简要描述

- 获取内部联系人备注列表

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetColleaRemarkList
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |
##### 请求示例 

``` 

{
    "uuid": "1688855749266556", 
    "limit": 100, 
    "seq": 0
}
```

##### 返回示例 

``` 
{
    "data": {
        "list": [
            {
                "vid": 16888548534xxx90,
                "description": "2343243222滴滴滴",
                "remark": "3242342229低调低调滚滚滚",
                "timestamp": 1683268963
            },
            {
                "vid": 16888565xxx48765,
                "description": "你好简介",
                "remark": "备注啦啦啦22222",
                "timestamp": 1683274153
            }
        ],
        "seq": 15312891
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.12 置顶设置或者取消 [已封装: SetTop]
##### 简要描述

- 置顶设置或者取消

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SetTop
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |
##### 请求示例 

``` 

{
    "uuid":"31c16499bb967a0f5f2eec81af770066",
    "vid": 7881302555913666, //群id 或者联系人id等
    "isRoom":false, //vid是否为群类型 true是 false否
    "state": 0 //是否置顶 0取消置顶 1设置置顶
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.13 免打扰设置或者取消 [已封装: SetShield]
##### 简要描述

- 免打扰设置或者取消

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SetShield
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |
##### 请求示例 

``` 

{
    "uuid":"31c16499bb967a0f5f2eec81af770066",
    "vid": 7881302555913666, //群id 或者联系人id等
    "isRoom":false, //vid是否为群类型 true是 false否
    "state": 0 //是否设置免打扰 0取消 1设置
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.14 免打扰并且折叠设置或者取消 [已封装: SetTold]
##### 简要描述

- 设置免打扰并且折叠设置或者取消

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SetTold
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |
##### 请求示例 

- 折叠后的样子
![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=3fcf9bb924cd03e539f924f674fc18b0&file=file.png)

``` 

{
    "uuid":"31c16499bb967a0f5f2eec81af770066",
    "vid": 7881302555913666, //群id 或者联系人id等
    "isRoom":false, //vid是否为群类型 true是 false否
    "state": 0 //是否设置免打扰且折叠联系人 0取消 1设置
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.15 星星标记设置或者取消 [已封装: SetMark]
##### 简要描述

- 星星标记设置或者取消

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SetMark
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |
##### 请求示例 

``` 

{
    "uuid":"31c16499bb967a0f5f2eec81af770066",
    "vid": 7881302555913666, //群id 或者联系人id等
    "isRoom":false, //vid是否为群类型 true是 false否
    "state": 0 //是否标记 0取消 1设置
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.16 拉黑或者移除拉黑联系人 [已封装: AddSubContactBlack]
##### 简要描述

- 拉黑或者移除拉黑联系人

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/AddSubContactBlack
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |


##### 请求示例 

``` 
//拉黑后获取外部联系列表中 status=8的 是拉黑的联系人
{
    "uuid": "70ed16e18d60b83b659b699a47efdbd6",
    "vid":7881302555913738, //要拉黑的联系人id
    "type":2 //1拉黑 2移除拉黑
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.17 根据手机号搜索联系人 [已封装: SearchContactByPhone]


##### 简要描述

- 根据手机号搜索联系人

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SearchContact
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |

##### 请求示例 

``` 
{
    "uuid":"3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "phoneNumber":"1357xxxx"
}
```

##### 返回示例 

``` 
{
    "data": {
        "UserList": [
            {
                "headImg": "https://wework.qpic.cn/wwpic/117798_FP8D8xt3RcCNze9_1654278287/0",
                "ticket": "114A2A26E50DE4BC6351E5E2481F1FE6A544AC2A656CE72D2AC8474F8479A79DF7E9E87C93E85792D5A57C3DBAD6E4BDA1262D3CEAE434A0FC669A4301E653FF95870E989F50BDE208416535BDFF8AB1",
                "user_id": 1688xxx579393,
                "sex": 1,
                "name": "*试",
                "state": "1", //企微
                "corp_id": 1970325105140387
            },
            {
                "headImg": "http://wx.qlogo.cn/mmhead/PiajxSqBRaEKbb5Mic6z0NYdlP21NqUUHUpgVvn7VtLFpqcldicldJ0tw/0",
                "ticket": "114A2A26E50DE4BC6351E5E2481F1FE6A544AC2A656CE72D2AC8474F8479A79DF7E9E87C93E85792D5A57C3DBAD6E4BDA1262D3CEAE434A0FC669A4301E653FF95870E989F50BDE208416535BDFF8AB1",
                "user_id": 7881303xxx682,
                "openid": "orFrbsjn3TUv6HTzWlZqNQ54VOUM",
                "sex": 2,
                "name": "xxx",
                "state": "2" //个微
            }
        ]
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.18 搜索添加外部联系人 [已封装: AddSearch]



##### 简要描述

- 搜索添加外部联系人

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/AddSearch
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |

##### 请求示例 

``` 
{
    "uuid": "3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "vid": 7881303344205927,
    "optionid": "orFrbsmzGRd4ZYTNSceHtSz6NH9E",
    "phone": "135xxxxxxxx",
    "content": "你好",
    "ticket": "154701AA31F472DCF58E9C272B0B4FED753C8890153760F712AABE351500965E11B06233D3FB34BE19095F0FD4460F1A5DF9236B9B77824FDF92A09F04BA604556A228E187A4B9CDB3CB0D2Axxxxxx"
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.19 搜索添加企业用户 [已封装: AddSearchWxwork]



##### 简要描述

- 搜索添加企微用户

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/AddSearchWxwork
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |
##### 请求示例 

``` 
{
    "uuid":"4400e170-373e-4bae-9b72-83063ca0bd65",
    "vid": 16888556373xxxxx,
    "corpid": 1970326672xxxxx,
    "ticket":"154701AA31F472DCF58E9C272B0B4FED753C8890153760F712AABE351500965E11B06233D3FB34BE19095F0FD4460F1A5DF9236B9B77824FDF92A09F04BA604556A228E187A4B9CDB3CB0D2Axxxxxx",
    "content": "测试消息"
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.20 根据名片添加好友 [已封装: AddCard]



##### 简要描述

- 根据名片添加好友

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/AddBusinessCard
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |

##### 请求示例 

``` 
{
    "uuid":"3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "vid": 7881303344205927,//要添加的人id
    "recommendId":7881300180956131, //推荐人id 谁发你的名片填谁的id
	"roomid":10817805073014679,
    "content": "你好"
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.21 直接添加好友（被删除过才能用） [已封装: AddWxUser]


##### 简要描述

- 直接添加好友（被删除过才能用）

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/addWxUser
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |

##### 请求示例 

``` 
{
    "uuid": "5d5d05064419c38243bf91dcee3ac5d5", 
    "vid": 7881302555913738, 
    "content": "哈哈哈哈"
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

---

### 8.22 共享联系人

    
##### 简要描述

- 共享联系人到其他同事账号

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/OperCustomer
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|

##### 请求示例 

``` 
{
    "uuid":"4ba46fd93c1e1831bb2dbf52776f0455",
    "thisvid":1688854256622049, //当前账号vid
    "recommendvid":7881303393914243,//要共享的好友id
    "sendvid": 1688854256622049 //要共享给哪个同事的id
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```

### 8.23 添加共享联系人

    
##### 简要描述

- 添加同事共享过来用户好友

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/AddShareUser
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|

##### 请求示例 

``` 
{
    "uuid": "f703ca61f08b806159ff9ec4bf20ae96", 
    "vid": 7881302555913738, //同事共享过来的vid
    "content": "在吗，通过下"//添加内容
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "ok"
}
```