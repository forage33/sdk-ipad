> ← [返回文档索引](../INDEX.md) | [架构文档](../ARCHITECTURE.md)
## 四、群操作 (group.go)

### 4. 1 获取客户群列表 [已封装: GetChatroomMembers]

##### 简要描述

- 获取客户群列表

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetChatroomMembers
  
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
    "uuid":"3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "limit":5, //每次查询数量
    "star_index":0  //第一次默认传0 根据返回值查询下一次的类似页码
}
```

##### 返回示例 

``` 
{
    "data": {
        "total": 7,
        "next_start": 5,
        "roomList": [
            {
                "room_id": 10696052955013024,
                "create_user_id": 0,
                "infoticket": "9FE241BC08F50F9A2FCBA692759B6FD8A5EA88B107787838AC06D1FE8D5835E04813EA38915241BCC181108D88408912D2A1C9D559384C0BEB261275B6EB7591",
                "update_time": 1650135314,
                "total": 4,
                "create_time": 1648889598,
                "nickname": "dddddddddddd",
                "roomurl": "https://wework.qpic.cn/wwpic/699045_06DbkaBSSVeLfTg_1652088785/0"
            },
            {
                "room_id": 10696052955013038,
                "create_user_id": 0,
                "infoticket": "8501CD1288CE0BE28ACEB6B1711D4EE864CB477C179F78B67E799DF07FE6166A6FDBAAF1D6962F2B4C3066DEF1E0B891A6B6496F41F4A34E5E59BE26571174F1",
                "update_time": 1648918889,
                "total": 2,
                "create_time": 1648889513,
                "nickname": "",
                "roomurl": "https://wework.qpic.cn/wwpic/984831_vJ5pBSAJS1uIL6C_1648889499/0"
            },
            {
                "room_id": 10696052955012990,
                "create_user_id": 0,
                "infoticket": "CDDC4E02A82DD6A2D24D40701CD0F155D5ADD102C7F9D5EE189C7EB51D3F4630058E85DCB9FC6143F8D496B31A17526B7702EE4FA9B6A674B136D788DD023E96",
                "update_time": 1647968439,
                "total": 3,
                "create_time": 1647945500,
                "nickname": "",
                "roomurl": "https://wework.qpic.cn/wwpic/559949_Rm-uPI04Tn6s_pW_1647526323/0"
            },
            {
                "room_id": 10696052955016139,
                "create_user_id": 0,
                "infoticket": "4C29DACB7C81D790B6F25FDC3DE05049C3342D915DAD9C397A41539F7C9970C748BE6D51B82D45D9CE733B866A066F09C52A2C0B87528A23CD7C9C1BC70C24EE",
                "update_time": 1612542878,
                "total": 2,
                "create_time": 1600888512,
                "nickname": "",
                "roomurl": "https://wework.qpic.cn/wwpic/507283_C4iEBcybRuyTAdE_1645840879/0"
            },
            {
                "room_id": 10696052955013729,
                "create_user_id": 0,
                "infoticket": "CDD6706E9079E87BC2235C88B822FFC97E3FAC7E8114D6874D1F1DD1EDC69C20DE8D90D1AFA32460D66623899CCEFD2D3D591ACDAB5AB55321D27E8B3DB2281A",
                "update_time": 1650655257,
                "total": 3,
                "create_time": 1594170855,
                "nickname": "",
                "roomurl": "https://wework.qpic.cn/wwpic/683942_iMOh0nrgSByS0WN_1650641660/0"
            }
        ]
    },
    "errcode": 0,
    "errmsg": "ok"
}
```


### 4.2 获取群成员列表 [已封装: GetRoomUserList]


##### 简要描述

- 获取群成员列表

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetRoomUserList
  
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
    "uuid":"3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "roomid":10696052955016166
}
```

##### 返回示例 

``` 
{
    "data": {
        "room_id": 10696052955012996,
        "create_user_id": 1688853790599424,
        "total": 6,
        "notice_sendervid": 10696052955012996,
	    "anti_spam_rules": [    //防骚扰规则id
            7485268685645994616
        ],
        "create_time": 1590005179,
        "notice_content": "群公告啊啊啊啊",
        "nickname": "外部群2",
		"flag": 270532609,//是否禁止群邀请 270532609 放开  270565505禁止
		"new_flag": 3,//是否允禁止修改群名 3禁止 2放开
		"is_out": 1,//是否退群了 1是 0否
        "member_list": [
            {
                "unionid": "ozynqskXWQK3wAWWKT_izAOfcIlY",
                "create_time": 0,
                "sex": 1,
                "mobile": "1xxxxx",
                "acctid": "xxxx",
                "join_scene": 1,
                "avatar": "https://wework.qpic.cn/wwpic/827947_xnqwcg3ATxmoUf-_1656095832/0",
                "english_name": "yingwenmingcheng",
                "realname": "xxxxx",
                "room_notes": "外部群2备注",
                "jointime": 1677418297,
                "nickname": "但是发射点",
                "room_nickname": "ccccccccccccc",
                "position": "232",
                "uin": 1688853790xxx,
                "invite_user_id": 16888572xxx,
                "corp_id": 1970325xxxx
            }
        ],
        "notice_time": 10696052955012996,
        "managers": []
    },
    "errcode": 0,
    "errmsg": "ok"
}
```


### 4.3 获取会话中的群列表 [已封装: GetSessionRoomList]


##### 简要描述

- 获取会话列表中的群聊

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetSessionRoomList
  
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
    "uuid": "7b486694-f58c-4259-b9fb-61acbeba19a2",
    "limit":10, 
    "star_index":0
}
```

##### 返回示例 

``` 
{
    "data": {
        "seq": 0,
        "room_list": [
            {
                "room_id": 10775490987050917,
                "create_user_id": 168877777777777,
                "total": 4,
                "flag": 128,
                "create_time": 1673335728,
                "nickname": "外部群测试2",
                "managers": []
            },
            {
                "room_id": 10779742636409552,
                "create_user_id": 168885877777777,
                "total": 3,
                "flag": 128,
                "create_time": 1673335408,
                "nickname": "测试",
                "managers": []
            }
        ]
    },
    "errcode": 0,
    "errmsg": "ok"
}
```


### 4.4 同意进群 [已封装: AgreeAddRoom]

##### 简要描述

- 同意进群

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/AgreeAddRoom
  
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
    "url":"https://work.weixin.qq.com/wework_admin/external_room/join/exceed?vcode=7537164e42c18af0b77deb05e1ba75fa" //邀请进群连接中的url
}
```

##### 返回示例 

``` 
{
    "data": {
        "data": {
            "vid": "1688853790599424",
            "status": "InviteTicketExpired"
        }
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 4.5 二维码进群 [已封装: AgreeCardAddRoom]


##### 简要描述

- 二维码进群

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/AgreeCardAddRoom
  
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
    "code":"7EefxHYdxxxxxx" //群二维码解析出来的连接结尾
}
```
![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=5d25e2a6ec987813c78fc1fe615d9ce3&file=file.png)
##### 返回示例 

``` 
{
    "data": {
        "data": {
            "vid": "1084618918xxxxxx",//群id
            "status": "Success"
        }
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 4.6 创建内部群聊 [已封装: CreateRoomNei]


##### 简要描述

- 创建内部群聊

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/CreateRoomNei
  
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
    "vids":[
        1688853801173786
    ],
    "roomName":"创建群聊了啊啊啊啊"
}
```

##### 返回示例 

``` 
{
    "data": {
        "roomname": "创建群聊了啊啊啊啊",
        "createTime": 1652186733,
        "createid": 1688853790599424,
        "roomid": 201446962331843
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 4.7 创建外部群聊 [已封装: CreateRoomWx]



##### 简要描述

- 创建外部群聊

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/CreateRoomWx
  
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
    "uuid":"3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "vids":[
        788xx2021632,
        1688xx798900687
    ],
    "roomName":"创建群聊了啊啊啊啊"
}
```

##### 返回示例 

``` 
{
    "data": {
        "roomname": "创建群聊了啊啊啊啊223332",
        "createTime": 1757407907,
        "createid": 168xx772759,
        "roomid": "1076666384084"
    },
    "errcode": 0,
    "errmsg": "ok"
}
{
    "data": null,
    "errcode": -24001048,
    "errmsg": "未验证企业仅可邀请200位外部联系人进群，验证后可继续邀请"
}
```


### 4.8 发送群公告 [已封装: SendNotice]


##### 简要描述

- 发送群公告

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendNotice
  
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
    "uuid":"3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "roomid":201446962331843, 
    "msg":"66666666"
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


### 4.9 修改群名 [已封装: UpdateRoomName]



##### 简要描述

- 修改群名

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/UpdateRoomName
  
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
    "roomid":10696052955012989, 
    "roomname":"3333"
}
```

##### 返回示例 

``` 
{
    "data": {
        "roomname": "3333",
        "roomid": 10696052955012989
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 4.10 发送链接的方式邀请成员进群 [已封装: InvitationToRoomLink]



##### 简要描述

- 发送链接的方式邀请成员进群

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/InvitationToRoomLink
  
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
    "roomid":10696052955013729, 
    "vids":[
        7881303344205927
    ]
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

### 4.11 直接邀请进群 [已封装: InvitationToRoom]



##### 简要描述

- 直接邀请进群

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/InvitationToRoom
  
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
    "roomid":10696052955013729, 
    "vids":[
        7881303344205927
    ]
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

### 4.12 移除群成员 [已封装: DelRoomUsers]



##### 简要描述

- 移除群成员

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/DelRoomUsers
  
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
    "roomid":10696052955016166, 
    "vids":[
        1688853931795099
    ]
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

### 4.13 设置群内昵称 [已封装: SetMyRoomName]

##### 简要描述

- 设置群内昵称

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SetMyRoomName
  
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
    "roomid":10696052955016166, 
    "content":"啊啊啊"
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



### 4.15 转让群主 [已封装: TransferChatroomOwner]

##### 简要描述

- 转让群主

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/TransferChatroomOwner
  
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
    "roomid":10696052955013038, 
    "vid":1688853792184982 //要转群主id
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

### 4.16 解散群 [已封装: DissolutionRoom]



##### 简要描述

- 解散群

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/DissolutionRoom
  
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
    "roomid":10696052955012990
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

### 4.17 获取群二维码 [已封装: WxRoomInvite]

##### 简要描述

- 获取群二维码

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/WxRoomInvite
  
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
    "roomid":10696052955016166
}
```

##### 返回示例 

``` 
{
    "data": {
        "QrCodePath": "http://47.94.7.218:8083/RoomQrCode/10696052955016166.jpg",
        "image_url": "",
        "roomid": 10696052955016166
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 4.18 获取欢迎语列表 [已封装: GetWelcomeMsgList]

##### 简要描述

- 获取欢迎语列表

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetWelcomeMsgList
  
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
    "offset":"",
    "limit":10
}
```

##### 返回示例 

``` 
{
    "data": {
        "offset": "xxxxxx",
        "WelocomeMsgList": [
            {
                "updateid": 1688853790599424,
                "createtime": 1649755966,
                "createid": 1688853790599424,
                "file_id": "",
                "aes_key": "",
                "id": 25,
				"type": 2,
                "img_size": 0,
                "updatetime": 1649755966,
                "content": "协议欢迎语嗷嗷#客户昵称##客户昵称#结束词嗷嗷",
                "WelcineMsgLink": "https://www.baidu.com/?tn=62095104_43_oem_dg",
                "md5": ""
            },
            {
                "updateid": 1688853790599424,
                "createtime": 1635458665,
                "createid": 1688853790599424,
                "file_id": "",
                "aes_key": "",
                "id": 24,
				"type": 2,
                "img_size": 0,
                "updatetime": 1635458665,
                "content": "哈哈哈哈#客户昵称#你好啊",
                "WelcineMsgLink": "https://www.baidu.com/",
                "md5": ""
            },
            {
                "updateid": 1688853790599424,
                "createtime": 1634843452,
                "createid": 1688853790599424,
                "file_id": "",
                "aes_key": "",
                "id": 22,
				"type": 2,
                "img_size": 0,
                "updatetime": 1635456204,
                "content": "你好啊#客户昵称#",
                "WelcineMsgLink": "https://easydoc.net/doc/22586813/HFiKj058/OXm2BSeu",
                "md5": ""
            },
            {
                "updateid": 1688853790599424,
                "createtime": 1633529586,
                "createid": 1688853790599424,
                "file_id": "",
                "aes_key": "",
                "id": 21,
				"type": 2,
                "img_size": 0,
                "updatetime": 1635453339,
                "content": "测试你好#客户昵称#你好啊",
                "WelcineMsgLink": "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",
                "md5": ""
            },
            {
                "updateid": 1688853790599424,
                "createtime": 1635452841,
                "createid": 1688853790599424,
                "file_id": "3069020102046230600201000204ea44290002030f55ca0204e25a44da0204617b0669042466346336663037632d343666652d343530662d616336302d39633438376639346330323502010002030484e00410c962f3f12c0cb03b080678d1dc006e370201010201000400",
                "aes_key": "32366230646636666462613465343534",
                "id": 23,
				"type": 2,
                "img_size": 296144,
                "updatetime": 1635452841,
                "content": "你好客户#客户昵称#",
                "WelcineMsgLink": "",
                "md5": "c962f3f12c0cb03b080678d1dc006e37"
            },
            {
                "updateid": 1688853790599424,
                "createtime": 1633528279,
                "createid": 1688853790599424,
                "file_id": "",
                "aes_key": "",
                "id": 20,
				"type": 2,
                "img_size": 0,
                "updatetime": 1633528279,
                "content": "测试欢迎语2222222222222222222",
                "WelcineMsgLink": "",
                "md5": ""
            },
            {
                "updateid": 1688853790599424,
                "createtime": 1633521415,
                "createid": 1688853790599424,
                "file_id": "",
                "aes_key": "",
                "id": 19,
				"type": 2,
                "img_size": 0,
                "updatetime": 1633521415,
                "content": "测试欢迎语2222222222222222222",
                "WelcineMsgLink": "",
                "md5": ""
            },
            {
                "updateid": 1688853790599424,
                "createtime": 1633521217,
                "createid": 1688853790599424,
                "file_id": "",
                "aes_key": "",
                "id": 18,
                "img_size": 0,
                "updatetime": 1633521217,
                "content": "测试欢迎语",
                "WelcineMsgLink": "",
                "md5": ""
            },
            {
                "updateid": 1688853790599424,
                "createtime": 1633520967,
                "createid": 1688853790599424,
                "file_id": "",
                "aes_key": "",
                "id": 17,
				"type": 2,
                "img_size": 0,
                "updatetime": 1633520967,
                "content": "欢迎你啊啊啊 啊啊啊",
                "WelcineMsgLink": "",
                "md5": ""
            },
            {
                "updateid": 1688853790599424,
                "createtime": 1633452337,
                "createid": 1688853790599424,
                "file_id": "",
                "aes_key": "",
                "id": 16,
				"type": 2,
                "img_size": 0,
                "updatetime": 1633452337,
                "content": "88888888888888888",
                "WelcineMsgLink": "",
                "md5": ""
            }
        ]
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 4.19 添加欢迎语 [已封装: AddWelcomeMsg]



##### 简要描述

- 添加欢迎语

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/AddWelcomeMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型   | 说明                                         |
| :----- | :--- | :----- | -------------------------------------------- |
| uuid   | 是   | String | 每个实例的唯一标识，根据uuid操作具体企业微信 |

##### 请求示例 

``` 
//list 和 连接可以为空 可以单独填list 其他的为空就是纯文字欢迎语 都填就是文字连接欢迎语
{
    "uuid":"3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "list":[
        "协议欢迎语嗷嗷",
        "#客户昵称#",
        "结束词嗷嗷"
    ],
    "url":"https://www.baidu.com/?tn=62095104_43_oem_dg", 
    "title":"百度啊啊啊",
    "content":"百度内容",
    "headImg":"https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png"
}
```

##### 返回示例 

``` 
{
    "data": {
        "id": 26
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 4.20 设置欢迎语 [已封装: SetWelcomeMsg]



##### 简要描述

- 设置欢迎语

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SetWelcomeMsg
  
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
    "indexid":25,
	"type": 2, //获取欢迎语列表中获取，不填默认是2
    "roomid":10696052955013038
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

### 4.21 取消欢迎语 [已封装: DelWelcomeMsg]



##### 简要描述

- 取消欢迎语

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/DelWelcomeMsg
  
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
    "roomid":10696052955013038
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

### 4.22 添加群成员好友 [已封装: AddRoomFriends]



##### 简要描述

- 添加群成员好友

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/AddRoomFriends
  
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
    "roomid":10696052955012990,
    "corpid":0,
    "vid":7881303344205927,
    "content":"你好啊"
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

### 4.23 退出群聊 [已封装: OutRoomReq]



##### 简要描述

- 推出群聊

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/OutRoomReq
  
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
    "uuid": "4db64991-8ce3-427b-8186-df8841216c4c",
    "roomid":10696052955016166
}
```

##### 返回示例 

``` 
{
    "data": {
        "room_id": 10696052955016166,
        "msg_id": 1092779
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 4.24 获取群成员列表(简洁版) [已封装: GetRoomDesUserList] 

##### 简要描述

- 获取群成员列表

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetRoomDesUserList
  
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
    "uuid":"3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "roomid":10696052955016166
}
```

##### 返回示例 

``` 
{
    "data": {
        "room_id": 10696052955012996,
        "create_user_id": 1688853790599424,
        "total": 6,
        "notice_sendervid": 10696052955012996,
        "create_time": 1590005179,
        "notice_content": "群公告啊啊啊啊",
        "nickname": "外部群2",
		"flag": 270532609,//是否禁止群邀请 270532609 放开  270565505禁止
		"new_flag": 3,//是否允禁止修改群名 3禁止 2放开
        "member_list": [
 {
                "room_notes": "",
                "jointime": 1600838896,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853787195331,
                "invite_user_id": 1688853790599424
            },
            {
                "room_notes": "",
                "jointime": 1600838953,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853787195332,
                "invite_user_id": 1688853790599424
            },
            {
                "room_notes": "",
                "jointime": 1600839131,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853787195334,
                "invite_user_id": 1688853790599424
            },
            {
                "room_notes": "",
                "jointime": 1600839021,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853788065293,
                "invite_user_id": 1688853790599424
            },
            {
                "room_notes": "",
                "jointime": 1622168755,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853789022262,
                "invite_user_id": 1688853790599424
            },
            {
                "room_notes": "",
                "jointime": 1589858817,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853790599424,
                "invite_user_id": 1688853790599424
            },
            {
                "room_notes": "",
                "jointime": 1600838997,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853792184982,
                "invite_user_id": 1688853790599424
            },
            {
                "room_notes": "",
                "jointime": 1600839106,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853793115324,
                "invite_user_id": 1688853790599424
            },
            {
                "room_notes": "",
                "jointime": 1600839254,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853793115330,
                "invite_user_id": 1688853790599424
            },
            {
                "room_notes": "",
                "jointime": 1600839058,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853794914376,
                "invite_user_id": 1688853790599424
            },
            {
                "room_notes": "",
                "jointime": 1600839215,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853794914379,
                "invite_user_id": 1688853790599424
            },
            {
                "room_notes": "",
                "jointime": 1600839300,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853796009464,
                "invite_user_id": 1688853790599424
            },
            {
                "room_notes": "",
                "jointime": 1600838842,
                "nickname": "",
                "join_scene": 0,
                "uin": 1688853798900687,
                "invite_user_id": 1688853790599424
            }
        ],
        "notice_time": 10696052955012996,
        "managers": []
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 4.25 获取群详细和头像 [已封装: GetRoomInfo]



##### 简要描述

- 获取群头像详细

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetRoomInfo
  
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
    "roomid":10696052955016166
}
```

##### 返回示例 

``` 
{
    "data": {
        "room_id": "10787964290465744",
        "create_user_id": 1688853787195332,
        "total": 19,
        "notice_sendervid": 10787964290465744,
        "flag": 129,
		"is_out": 0, //是否退群 1是 0否
        "create_time": 1649388196,
        "notice_content": "",
        "image_url": "https://wework.qpic.cn/wwpic/512746_S2SOrz0ZRFqKaLS_1678715585/0",
        "nickname": "外部群1",
        "new_flag": 0,
        "notice_time": 10787964290465744,
        "managers": []
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 4.26 获取群头像 [已封装: WxRoomHeader]



##### 简要描述

- 获取群头像

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/WxRoomHeader
  
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
    "roomid":10696052955016166
}
```

##### 返回示例 

``` 
{
    "data": {
        "image_url": "https://wework.qpic.cn/wwpic/23427_ROFv_3dGSyKNA6F_1697622635/0",
        "roomid": 10973555009565180
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

#### 群设置

####  4.27 设置群备注

#### 4.28 禁止修改群名字

#### 4.29 群邀请确认

#### 4.30 移除群管理

#### 4.31 设置群管理

#### 4.32 获取群防骚扰规则列表

#### 4.33 设置或者移除群防骚扰规则

#### 4.34 获取群聊黑名单列表

#### 4.35 添加或者移除群聊黑名单

#### 4.36 禁止群内添加和禁止修改群名

---

