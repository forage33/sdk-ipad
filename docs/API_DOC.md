# 企业微信协议服务 API 接口文档

> 数据来源：ShowDoc (https://www.showdoc.com.cn/1889765213630885)
> 生成时间：2026-04-16
> 说明：本文档汇总了协议服务的全部 API 接口，并标注了 SDK 封装状态

## 封装状态说明

| 标记 | 含义 |
|------|------|
| ✅ 已封装 | SDK 中已有完整实现 |
| 🔧 TODO | SDK 中已有 stub，待补全参数和实现 |
| ❌ 未封装 | SDK 中完全没有，需要新增 |

---

## 一、初始化 (wx_init.go)

### 1.1 初始化企业微信 ✅ `InItWXWork`

##### 简要描述

- 初始化企业微信

##### 请求URL

- ` http://47.94.7.218:8083/wxwork/init

##### 请求方式

- POST
- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----|:---|:-----|-----|
|vid |否  |long |第一次登录可以传空字符串,登陆成功后下次初始化填登录账号的id，用来自动登录操作，这个值代表了这个号的设备绑定，首次不传值，会默认生成新设备信息，然后扫码登录了，登陆了A账号成功后，后续在重新登录的时候，vid参数就传A账号的账号id 16888开头的就是账号id。   |
|proxySituation |否  |int |是否全局代理(必须设置代理的时候他才生效，不设置代理默认就是0，设置了代理不传这个字段默认就是全局代理)，1是全局代理代表设置上就是长效的代理不能取消，0代表不是全局代理，可以设置上登录成功后用代理设置接口去取消代理。  |

##### 请求示例

```
{
    "vid":"", 
    "ip":"", //代理ip
    "port":"",//代理端口
    "proxyType":"", //http 代理类型
    "userName":"", //代理账号 没有不传
	"passward":"", //代理密码 没有不传
	"proxySituation":0,  
	"deverType":"ipad"
}
```

##### 返回示例

```
{
    "data": {
        "uuid": "427d7ee5-3a1c-4183-a83b-532ba1e71a1e", //uuid是这个账号生命周期内一直用到的，用来操作哪个账号执行接口功能的。
		"is_login": "false"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```



### 1.2 设置消息回调地址 ✅ `SetCallBackURL`

##### 简要描述

- 设置消息回调地址

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SetCallbackUrl
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|url |是  |String |消息回调地址|

##### 请求示例(HTTP回调设置)

``` 
//消息回调地址设置规则 
//1. 搭建一个web服务 
//2. 设定post请求接口，ContentType:"application/json" json请求方式
//3. 接收参数设置成对象，对象属性为以下属性
		{
			"uuid":"", //uuid 运行实例id
			"json":"", //回调消息内容为json格式
			"type": //消息类型
		}
//具体回调内容及类型请查看下发文档
{
    "uuid":"f5a22e9b-9664-4250-b40a-08741dba549c",
    "url":"http://127.0.0.1:8084/wxwork/callback"
}
```
回调示例java版本
![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=9f20ec3f3a49733f458577c273d8bca0)
##### 请求示例(RABBITMQ)

``` 
{
    "uuid": "5a5e6ebc176da6ea0d8cb4ffd09f1956",
    "callbackType":"RABBITMQ",
    "mqExchange":"message.exchange.ultraMsg", //交换机 direct类型
    "mqRoutingKey":"ssss",//路由键
	"extraContent":"66666" //自定义回调额外内容可以为null 不传就行
}
```
##### 返回示例

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "设置成功"
}
```


### 1.3 获取运行中的实例 ✅ `GetRunClient`

##### 简要描述

- 获取运行中的实例：直接post请求地址就可以了

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetRunClient
  
##### 请求方式
- POST 

##### 参数
无
##### 请求示例 

``` 
无参数
```

##### 返回示例 

``` 
{
    "data": [
        {
            "1688857362464764": {
                "isLogin": true,
                "clientId": "8314b481e596fc85d850487e35bbe88b",
                "updateTime": 1715042291,
                "userid": 1688xxxx464764,
                "uuid": "8314b481e596fc85d850487e35bbe88b",
                "requrl": "http://47.94.7.218:",
                "object": {
                    "unionid": "ozynqstMEN_38dTCLJIlwe73ZhB0",
                    "create_time": 0,
                    "sex": 1,
                    "mobile": "135xxxxx01",
                    "Corp_full_name": "",
                    "acctid": "Baxxxang",
                    "scorp_id": "ww75xxxxccc9e",
                    "avatar": "https://wework.qpic.cn/wwpic/533154_tevrSjnsQVKBlHH_1678777868/0",
                    "corp_name": "明月科技",
                    "english_name": "",
                    "realname": "刘xxx",
                    "user_id": 16888xxxxx,
                    "nickname": "班xxx",
                    "position": "",
                    "Corp_desc": "",
                    "corp_id": 1970325xxxx87
                }
            }
        }
    ],
    "errcode": 0,
    "errmsg": "获取成功"
}
```



### 1.4 根据uuid查看实例详情 ✅ `GetRunClientByUuid`

##### 简要描述

- 根据uuid查看实例详情

##### 请求URL
- ` http://127.0.0.1:8083/wxwork/GetRunClientByUuid
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |要查询的账号状态|
##### 请求示例 

``` 
{
    "uuid":"xxxxxx" //uuid
}
```

##### 返回示例 

``` 
{
    "data": {
        "logintime": 1744427201,
        "loginType": 2, //0 未初始化  1初始化了未登录  2已登录
        "user_info": {
            "object": {
                "unionid": "ozynqsgN5h3Vuulx4E07Oy17j-7s",
                "create_time": 0,
                "sex": 1,
                "mobile": "xxxx",
                "acctid": "xxxx",
                "scorp_id": "xxxxx",
                "avatar": "https://wework.qpic.cn/wwpic3az/319926_srytF-ZZSX2lOL7_1743776394/0",
                "corp_name": "xxx",
                "english_name": "xx",
                "realname": "xxx",
                "user_id": xxxx,
                "nickname": "xxx",
                "position": "",
                "corp_id": xxxx,
                "corp_full_name": "xxxxxx",
                "corp_desc": ""
            },
            "userid": xxxxxx,
            "clientId": "e1d672c613717ca41fcbd6823da9015a",
            "updateTime": 1745125943,
            "uuid": "e1d672c613717ca41fcbd6823da9015a",
            "requrl": "http://47.94.7.218:8083",
            "isLogin": true
        },
        "uuid": "e1d672c613717ca41fcbd6823da9015a"
    },
    "errcode": 0,
    "errmsg": "获取成功"
}
```


### 1.5 关闭链接 ✅ `CloseConnect`


##### 简要描述

- 关闭链接

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/CloseConnent
  
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
    "uuid":"f5a22e9b-9664-4250-b40a-08741dba549c"
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "关闭成功"
}
```

---

## 二、登录 (login.go)

### 2.1 获取登录二维码 ✅ `GetQrCode`


##### 简要描述

- 获取登录二维码

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/getQrCode
  
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
    "uuid":"f5a22e9b-9664-4250-b40a-08741dba549c"
}
```

##### 返回示例 

``` 
{
    "data": {
        "qrcode": "http://47.94.7.218:8083/980343ff-bf71-4789-8a0a-af3f7fd40bc0.png", //登陆二维码的直接访问连接
		"qrcode_data":"/9j/4AAQSkZJRgABAQEASABIAAD/2wBDABQODxIPDRQSEBIXFRQYHjIhHhwcHj0s",//登录二维码的base64位文件流
        "Ttl": 600,
        "Key": "5DBC31948BB057101F9C6B93569FB39B"
    },
    "errcode": 0,
    "errmsg": "获取二维码成功"
}
```


### 2.2 输入验证码 ✅ `CheckCode`

##### 简要描述

- 第一次登录的时候需要输入验证码，如果登陆过（然后初始化接口vid传值了，则在获取二维码就不需要输入验证码了）

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/CheckCode
  
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
    "uuid":"2b0863724106a1160212bd1ccf025295",
    "qrcodeKey":"B547FC8D049DFA333D89075C04BA9178", //qr_codekey
    "code":"369130"//验证码
}
```

##### 返回示例 

``` 
{
    "data": null,
    "errcode": 0,
    "errorcode": 0,
    "errmsg": "ok"
}
```
##### 注意事项
如果调用这个接口出现了这个错误返回：qrcode_not need verify，这个就是输入验证码前把手机上的验证码界面x掉了，不要提前x掉，输入完验证码自己就会关掉。



### 2.3 自动登录 ✅ `AutomaticLogin`

##### 简要描述

- 自动登录：自动登录条件，根据初始化企业微信接口，传递的vid，获取的uuid进行自动登录，如果初始化没有传递vid则自动登录失败，只要登陆过一次后期可以一直自动登录

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/automaticLogin
  
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
    "uuid": "8793dc28-a863-463b-a5d6-ccd369adb859"
}
```

##### 返回示例

``` 
{
    "data": null,
    "errcode": 0,
    "errmsg": "登陆成功"
}
```


### 2.4 退出登录 ✅ `LoginOut`


##### 简要描述

- 退出登录

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/LoginOut
  
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
    "uuid":"f5a22e9b-9664-4250-b40a-08741dba549c"
}
```

##### 返回示例 

``` 
{
    "data": {
        "code": 0,
        "errMsg": "ok",
        "obj": {},
        "pbObj": null
    },
    "errcode": 0,
    "errmsg": "退出成功"
}
```

---

## 三、发送消息 (message.go)

> SDK 已封装 8 个消息接口，文档中的完整消息接口列表未能全部抓取

### 3.1 撤回消息 [已封装: RevokeMsg]

##### 简要描述

- 撤回消息

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/RevokeMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|msgid |是  |int |要撤回的消息id|
|roomid |是  |long |如果撤回的消息是群消息，则填写群id，联系人消息则填0|
##### 请求示例 

``` 
{
    "uuid":"a4ea6a39-4b3a-4098-a250-2a07bef57355",
    "msgid":1063645,
    "roomid":0
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


### 3.2 发送文本消息 [已封装: SendTextMsg]

##### 简要描述

- 发送文本消息

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendTextMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人id|
|isRoom |是  |bool |是否是群消息|
|content |是  |String |发送文本内容|
##### 请求示例 

``` 
{
    "uuid":"1753cdff-0501-42fe-bb5a-2a4b9629f7fb",
	"kf_id":0,   //用来客服消息回复，正常发送忽略这个字段。
    "send_userid":7881302555913738,
    "isRoom":false,
    "content":"ddddddd"
}
```

##### 返回示例 

``` 
{
    "data": {
        "receiver": 7881302555913738,
        "sender": 1688853790599424,
        "at_list": [],
        "sender_name": "",
		"app_info":"from_msgid_697590829808721944",
        "is_room": 0,
        "sendtime": 1652169352,
        "msg_id": 1066230,
        "server_id": 12381203,
        "msgtype": 2,
        "content": "ddddddd"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 3.3 发送CDN图片消息 [已封装: SendCDNImage]

##### 简要描述

- 发送CDN图片消息

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendCDNImgMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|
|cdnkey |是  |String |cdnkey|
|aeskey |是  |String |aeskey|
|md5 |是  |String |md5|
|fileSize |是  |int |文件大小|
##### 请求示例 

``` 
{
    "uuid": "47b24a10e92f66b2fc440608a19c86c8",
    "send_userid": "7881302555913738",
    "kf_id": 0, //客服id 如果需要回复客服消息则填客服id 没有客服消息忽略他
    "isRoom": false, //是否发送群消息
    "cdnkey": "306b020102046430620201000204060b19e102034c4cd20204986b259902046807456b042466303163353336302d636437642d346335302d626238632d38356638333361613238376402031038000203317d8004107d8a8e79499210c6242c36afd10647c70201010201000400",
    "aeskey": "B26DB4824341608DB848A2DA563B7147",
    "md5": "7d8a8e79499210c6242c36afd10647c7",
    "fileSize": 3243381,
    "width": 1279, //宽度
    "height": 1706,//高度
    "thumb_image_height": 512, //缩略图高度
    "thumb_image_width": 384,//缩略图宽度
    "thumb_file_size": 15195,//缩略图大小
    "thumb_file_md5": "4331947dc5f67c0d8a2893653a3f0cee",//缩略图md5
    "is_hd": 1 //用来转发接收到的图片消息的， 如果接收的图片消息中的is_hd字段是1就填1 是2就填2，  正常上传发送忽略他
}
```

##### 返回示例 

``` 
{
    "data": {
        "receiver": 7881302555913738,
        "cdn_key": "308189020102048181307f0201000204ea44290002030f55cb0204a48e1e6f0204626c17c604444e45574944315f6561343432393030613438653165366636323663313833385f31313841414234302d453237422d346539312d383843442d33343735453032423531313202010002023c600410d9c8750bed0b3c7d089fa7d55720d6cf0201010201000400",
        "file_name": "",
        "sender_name": "",
        "is_room": 0,
        "server_id": 12381205,
        "size": 15444,
		"app_info":"from_msgid_697590829808721944",
        "sender": 1688853790599424,
        "aes_key": "66303662386666346430373037326666",
        "sendtime": 1652169457,
        "msg_id": 1066235,
        "msgtype": 14,
        "md5": "d9c8750bed0b3c7d089fa7d55720d6cf"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 3.4 发送CDN文件消息 [已封装: SendCDNFile]

##### 简要描述

- 发送CDN文件消息

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendCDNFileMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|

##### 请求示例 

``` 
{
    "uuid":"1753cdff-0501-42fe-bb5a-2a4b9629f7fb",
    "send_userid":7881302555913738,
    "cdnkey":"30680201020461305f0201xxxx0501d03e3d1c0480fd7d1193205a080160201050201000400",
    "aeskey":"38633161326636666630396135353630",
    "md5":"501d03e3d1c0480fd7d1193205a08016",
    "file_name":"111.txt",
    "fileSize":2813,
    "isRoom":false
}
```

##### 返回示例 

``` 
{
    "data": {
        "receiver": 7881302555913738,
        "cdn_key": "30680201020461305f0201000204ea44290002030f55cb0204a48e1e6f0204626d8bb0042430453235394136372d443435322d343032302d393341372d38344445364532363030303902010002020b000410501d03e3d1c0480fd7d1193205a080160201050201000400",
        "file_name": "协议.txt",
        "sender_name": "",
        "is_room": 0,
        "server_id": 12381206,
        "size": 2813,
		"app_info":"from_msgid_697590829808721944",
        "sender": 1688853790599424,
        "aes_key": "38633161326636666630396135353630",
        "sendtime": 1652169489,
        "msg_id": 1066237,
        "msgtype": 15,
        "md5": "501d03e3d1c0480fd7d1193205a08016"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 3.5 发送CDN语音消息 [已封装: SendCDNVoiceMsg]

##### 简要描述

- 发送CDN语音消息

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendCDNVoiceMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
    "uuid":"1753cdff-0501-42fe-bb5a-2a4b9629f7fb",
    "send_userid":7881302555913738,
    "cdnkey":"308180020102047930770201000204bf366e0a02031e90380204afdf892b0204626d903d044c333230383031353337305f36353646373236463737363537303633363136313643363336373733373136325f34643061383739323537383238643361376636323565336234373835383961330201000202145004000201050201000400",
    "aeskey":"656F726F7765706361616C6367737162",
    "md5":"4d0a879257828d3a7f625e3b478589a3",
    "voice_time":4,
    "fileSize":5184,
    "isRoom":false
}
```

##### 返回示例 

``` 
{
    "data": {
        "receiver": 7881302555913738,
        "cdn_key": "308180020102047930770201000204bf366e0a02031e90380204afdf892b0204626d903d044c333230383031353337305f36353646373236463737363537303633363136313643363336373733373136325f34643061383739323537383238643361376636323565336234373835383961330201000202145004000201050201000400",
        "file_name": "",
        "sender_name": "",
        "is_room": 0,
		"app_info":"from_msgid_697590829808721944",
        "server_id": 12381207,
        "size": 5184,
        "sender": 1688853790599424,
        "aes_key": "656F726F7765706361616C6367737162",
        "sendtime": 1652169528,
        "msg_id": 1066239,
        "msgtype": 16,
        "md5": "4d0a879257828d3a7f625e3b478589a3"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```


### 3.6 发送CDN视频 [已封装: SendCDNVideoMsg]

##### 简要描述

- 发送CDN视频消息

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendCDNVideoMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
    "uuid":"505f61baa654b5b684a8c1917ad8bb00",
    "send_userid":7881302555913738,
    "cdnkey":"306b02010204643062020100020x04109c6131377dae038824b43f8e388efa1c0201040201000400",
    "aeskey":"BE02BA568E3B545F04009F5E81972C0E",
    "md5":"9c6131377dae038824b43f8e388efa1c",
    "video_duration":38,
    "fileSize":3561782,
    "video_img_size":65472, 
    "video_width":540, //视频宽度
    "video_height":1170,//视频高度
    "isRoom":false
}
```

##### 返回示例 

``` 
{
    "data": {
        "video_duration": 38,//视频时常
        "video_height": 1170,//视频高度
        "openim_cdn_ld_aeskey": "",
        "receiver": 78813025559xxx,  //接收人
        "openim_cdn_ld_size": 65472, //视频缩略图大小
        "file_name": "",
        "video_width": 540, //视频宽度
        "sender_name": "",
        "cdn_Key": "",
        "is_room": 0,
        "server_id": 7651694, //消息seq
        "file_size": 3561782,//视频大小
        "size": 3561782,//视频大小
        "VideoDuration": 38, //视频时常
        "sender": 1688854256622049,
        "file_id": "306b020102046430620201000204060b19ex131377dae038824b43f8e388efa1c0201040201000400", //视频id
        "openim_cdn_authkey": "",
        "aes_key": "BE02BA568E3B545F04009F5E81972C0E", //aeskey
        "app_info": "CIOACBDIjp7Q6jIY4bOssJCAgAMgEA==",
        "sendtime": 1746609341, //发送时间
        "msg_id": 1014578,
        "msgtype": 23,//消息类型
        "md5": "9c6131377dae038824b43f8e388efa1c", 
        "preview_img_url": ""
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 3.7 发送大视频消息 [已封装: SendCDNBigVideoMsg]


##### 简要描述

- 发送大视频文件 需要走大文件上传别走cdn上传！！！！！！！！！！！！！！！！！！！！

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendCDNBigVideoMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
    "uuid":"1688853790xxx", //uuid 默认随机生成如果初始化传了id则用初始化的id作为uuid
    "send_userid":78813025xxx8,
    "cdnkey":"*1*u3W2k/2+7TigF3ieM2JYWY12fJ0m+w80zO84cIONH6YTO1P6L92K6u8yQCjGfAlJ8SxDBE/+9OD0s6neeXIErix3BGU8/AzaPPEEOPUkSQwO77/Yo1Gff2kkbclynhl7pnjN9Ureh7j5PptTxeuOjRmr7s=",
    "file_name":"d8f0a590e1x7e5d0265b5ae.MP4",
    "md5":"976A026AB9Bx75514A39E5C0",
    "video_duration":178,
    "fileSize":63647298,
    "imgurl":"https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",//视频封面图
    "isRoom":false
}
```

##### 返回示例 

``` 
{
    "data": {
        "receiver": 7881302333338,
        "cdn_Key": "*1*u3W2k/2+7TigF3icccYV28gBh3eTguzdQNDSvTsX6weXFApoG+5jHWihl4Tq7kaiosiwiKRcWqOipf/77GX1EC1yqSKYkW1jQp2mVi3u8NfqEAyHheXZYwmvCoV85QMvmxIKd0XXTxIjAdQhZJ4lVVo5euOjRmr7s=",
        "sender_name": "",
        "file_name": "d8f0a59cccb0a7e5d0265b5ae.MP4",
        "is_room": 0,
        "server_id": 12781142,
        "size": 178,
		"app_info":"from_msgid_697590829808721944",
        "sender": 16888xxx,
        "aes_key": "",
        "sendtime": 1676876091,
        "msg_id": 1112326,
        "msgtype": 22,
        "md5": ""
    },
    "errcode": 0,
    "errmsg": "ok"
}
```


### 3.8 发送大文件消息 [已封装: SendCDNBigFileMsg]


##### 简要描述

- 发送大文件消息 需要走大文件上传别走cdn上传！！！！！！！！！！！！！！！！！！！！

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendCDNBigFileMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
    "uuid":"1688d599424",
    "send_userid":78d5913738,
    "cdnkey":"*1*F34hSNMREYF+BNgTfNtYo7Lg5/va+EoTPkyESrZSpwtNs1IFEpFfXWbWXF+kvx+rrm7SUtsNBbdEUxBcMc5zyqXZ3Z1pT7CShVKq1LmjzbhtG3ZAK/NFdftW9WvvWpUoq2ynfnC765zusYDzYEiH5bbc4g=",
    "md5":"9c2e8389567e170012d3511166000151",
    "file_name":"2f0ecd499b93bd9cb0da022.zip",
    "fileSize":314065,
    "isRoom":false
}
```

##### 返回示例 

``` 
{
    "data": {
        "receiver": 788130x913738,
        "cdn_key": "*1*F34hSNMREYF+BNgTfNtYo7Lg5/va+EoTPkyESrZSpwtNs1IFEpFfXWbWXF+kvx+rrm7SUtsNBbdEUxBcMc5zyqXZ3Z1pT7CShVKq1LmjzbhtG3ZAK/NFx9WvvWpUoq2ynfnC765zusYDzYEiH5bbc4g=",
        "file_name": "2f0ec2d93x99b93bd9cb0da022.zip",
        "sender_name": "",
        "is_room": 0,
        "server_id": 12781146,
        "size": 314065,
        "sender": 16888x99424,
        "aes_key": "",
        "sendtime": 1676878382,
        "msg_id": 1112335,
        "msgtype": 20,
        "md5": "9c2e8389567e170012d3511166000151"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 3.9 发送链接卡片消息 [已封装: SendLinkMsg]


##### 简要描述

- 发送连接卡片

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendLinkMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
    "uuid":"543ed7f3-6ec1-4b1a-9e47-db8339a140f7",
    "send_userid":7881302555913738,
    "url":"https://www.baidu.com/",
    "title":"百度",
    "content":"百度内容",
    "imgurl":"https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",
    "isRoom":false
}
```

##### 返回示例 

``` 
{
    "data": {
        "des": "百度内容",
        "receiver": 7881302555913738,
        "headImg": "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",
        "sender": 1688853790599424,
        "sender_name": "",
		"app_info":"from_msgid_697590829808721944",
        "is_room": 0,
        "sendtime": 1652169666,
        "msg_id": 1066243,
        "server_id": 12381209,
        "title": "百度",
        "msgtype": 13,
        "url": "https://www.baidu.com/"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 3.10 发送小程序 [已封装: SendAppMsg]

##### 简要描述

- 发送小程序

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendAppMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
    "uuid":"543ed7f3-6ec1-4b1a-9e47-db8339a140f7",
    "send_userid":7881302555913738,
    "desc":"内容",
	"appName":"小程序名称头像右边的",
    "title":"小程序标题头像下边的",
    "weappIconUrl":"https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",//小程序头像
    "pagepath":"pages/train/index/index.html",//小程序跳转地址
    "username":"gh_c4a2a98a7366@app", //账号原始id
    "appid":"wx45dff5234240ad90",//小程序appid
    "cdnkey":"30680201020461305f750bed0b3c7d089fa7d55720d6cf0201010201000400", //小程序封面图 cdn上传获取
    "md5":"d9c8750bed0b3c7d089fa7d55720d6cf",//封面图md5
    "aeskey":"34323235376238666264636630663166",//封面图 返回的aeskey
    "fileSize":15444,//封面图大小
    "isRoom":false
}
```

##### 返回示例 

``` 
{
    "uuid":"3240fde0-45e2-48c0-90e8-cb098d0ebe43",
    "send_userid":7881302555913738,
    "desc":"测试内容",
    "title":"小程序标题头像下边的",
	"appName":"小程序名称头像右边的",
    "weappIconUrl":"https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",
    "pagepath":"pages/train/index/index.html",
    "username":"gh_c4a2a98a7366@app",
	"app_info":"from_msgid_697590829808721944",
    "appid":"wx45dff5234240ad90",
    "cdnkey":"30680201020461305f0201000204ea44290002030f55cb0204649f216f020462713262042444313436453338312d453244332d346438322d413443352d35453035453930384342463702010002023c600410d9c8750bed0b3c7d089fa7d55720d6cf0201010201000400",
    "md5":"d9c8750bed0b3c7d089fa7d55720d6cf",
    "aeskey":"34323235376238666264636630663166",
    "fileSize":15444,
    "isRoom":false
}
```

### 3.10b 发送视频号消息 [已封装: SendVideoNumber]

##### 简要描述

- 发送视频号消息

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendVideoNumber
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
    "uuid":"6d92d9e5-88fd-4d9f-8007-3e81f57a35a2",
    "send_userid":7881302555913738,
    "isRoom":false,
    "coverUrl":"http://wxapp.tc.qq.com/251/20304/stodownload?encfilekey=rjD5jyTuFrIpZ2ibE8T7YmwgiahniaXswqzREChAjGtdKh8GdzAL9lfgmmD9WIYiaF5OpAwBEjfBR8BecoibHmf3A2lfRJbj1zkUzCyPB2bAD5whFARU0fo1icUw&adaptivelytrans=0&bizid=1023&dotrans=0&hy=SH&idx=1&m=&scene=0&token=AxricY7RBHdVxH2gvQP2wl93KARNB9ELWFOYmcKON2EwsCqIofeHvdnibvw2icUJ30ODGQicZ03U3z8&finder_expire_time=1652448294&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAOf827OoyYAAAAAstQy6ubaLX4KHWvLEZgBPElIMQQFweEf6DzNPgMIDXIE6IPCSmgIHh35qH2EMG",
    "thumbUrl":"http://wxapp.tc.qq.com/251/20304/stodownload?encfilekey=rjD5jyTuFrIpZ2ibE8T7YmwgiahniaXswqzREChAjGtdKh8GdzAL9lfgmmD9WIYiaF5OpAwBEjfBR8BecoibHmf3A2lfRJbj1zkUzCyPB2bAD5whFARU0fo1icUw&adaptivelytrans=0&bizid=1023&dotrans=0&hy=SH&idx=1&m=&scene=0&token=AxricY7RBHdVxH2gvQP2wl93KARNB9ELWFOYmcKON2EwsCqIofeHvdnibvw2icUJ30ODGQicZ03U3z8&finder_expire_time=1652448294&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAOf827OoyYAAAAAstQy6ubaLX4KHWvLEZgBPElIMQQFweEf6DzNPgMIDXIE6IPCSmgIHh35qH2EMG",
    "avatar":"http://wx.qlogo.cn/finderhead/Y3WgNLFjO0cpO1hficKwicUcWiadyBmFf3Fk89eSBRI1xlQW1VebNKbeg/0",
    "nickname":"天下泉城",
    "desc":"#天下泉城 济南加油！",
    "url":"https://channels.weixin.qq.com/web/pages/feed?eid=export%2FUzFfAgtgekIEAQAAAAAAOf827OoyYAAAAAstQy6ubaLX4KHWvLEZgBPElIMQQFweEf6DzNPgMIDXIE6IPCSmgIHh35qH2EMG",
    "extras":"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",
	"objectId":"14479811159081949639",
	"objectNonceId":"1113158575529036062_0_0_0_0_0"
}
```

##### 返回示例 

``` 
{
    "data": {
        "thumb_url": "http://wxapp.tc.qq.com/251/20304/stodownload?encfilekey=rjD5jyTuFrIpZ2ibE8T7YmwgiahniaXswqzREChAjGtdKh8GdzAL9lfgmmD9WIYiaF5OpAwBEjfBR8BecoibHmf3A2lfRJbj1zkUzCyPB2bAD5whFARU0fo1icUw&adaptivelytrans=0&bizid=1023&dotrans=0&hy=SH&idx=1&m=&scene=0&token=AxricY7RBHdVxH2gvQP2wl93KARNB9ELWFOYmcKON2EwsCqIofeHvdnibvw2icUJ30ODGQicZ03U3z8&finder_expire_time=1652448294&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAOf827OoyYAAAAAstQy6ubaLX4KHWvLEZgBPElIMQQFweEf6DzNPgMIDXIE6IPCSmgIHh35qH2EMG",
        "cover_url": "http://wxapp.tc.qq.com/251/20304/stodownload?encfilekey=rjD5jyTuFrIpZ2ibE8T7YmwgiahniaXswqzREChAjGtdKh8GdzAL9lfgmmD9WIYiaF5OpAwBEjfBR8BecoibHmf3A2lfRJbj1zkUzCyPB2bAD5whFARU0fo1icUw&adaptivelytrans=0&bizid=1023&dotrans=0&hy=SH&idx=1&m=&scene=0&token=AxricY7RBHdVxH2gvQP2wl93KARNB9ELWFOYmcKON2EwsCqIofeHvdnibvw2icUJ30ODGQicZ03U3z8&finder_expire_time=1652448294&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAOf827OoyYAAAAAstQy6ubaLX4KHWvLEZgBPElIMQQFweEf6DzNPgMIDXIE6IPCSmgIHh35qH2EMG",
        "receiver": 7881302555913738,
        "extras": "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",
        "sender_name": "",
		"app_info":"from_msgid_697590829808721944",
        "is_room": 0,
        "avatar": "http://wx.qlogo.cn/finderhead/Y3WgNLFjO0cpO1hficKwicUcWiadyBmFf3Fk89eSBRI1xlQW1VebNKbeg/0",
        "server_id": 12381210,
        "url": "https://channels.weixin.qq.com/web/pages/feed?eid=export%2FUzFfAgtgekIEAQAAAAAAOf827OoyYAAAAAstQy6ubaLX4KHWvLEZgBPElIMQQFweEf6DzNPgMIDXIE6IPCSmgIHh35qH2EMG",
        "sender": 1688853790599424,
        "nickname": "天下泉城",
        "sendtime": 1652169903,
        "msg_id": 1066245,
        "msgtype": 141,
        "desc": "#天下泉城 济南加油！"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```


### 3.11 发送视频号直播消息 [已封装: SendVideoNumberZhiBo]


##### 简要描述

- 发送视频号直播

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendVideoNumberZhiBo
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
        "uuid": "57224e08ace6632e6eff73b77320ceb9",
        "send_userid": 7881302555913738,
        "isRoom": false,
        "coverUrl": "https://finder.video.qq.com/261/20304/stodownload?encfilekey=oibeqyX228riaCwo9STVsGLIia39tYEhYEEzrbjBy8M5IziaWiabDVOeztVKTTLE6QQSL0IjE17CicJb5gibat8LWkgomsEdxSyZpOEjx18bqyPeMhmE5yn7AGUx1g3jxBXUZXMoqqf1yiblic2s&token=cztXnd9GyrHd4iajaaPHCJW4ZtfSic8HlHVaWZ1JLMFSbRsbPFnQPdNlhk11ASpWdg0Oj2CL5t1lR7Vd2yjl4Ut3lAjx5oORV6ymyfMdcpqeePq29EdWeKBSDO5MlLRZb1z48EmHEeJUJp3TiaHUq8jkAFYqvvfFuvvmia8x8PLqFR70Hjg0q21sSytBmFoT1HyW&bizid=1023&hy=SZ&idx=1&m=a5d642b207fa007e27cc99051888c93f&finder_expire_time=1767928339&wxampicformat=503&picformat=200&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAKE8Gm6qxKAAAAAstQy6ubaLX4KHWvLEZgBPEpKEQGE5SObGOzNPgMIKM6o_uZoIKXQZfrl1JnToi",
        "thumbUrl": "https://finder.video.qq.com/261/20304/stodownload?encfilekey=oibeqyX228riaCwo9STVsGLIia39tYEhYEEzrbjBy8M5IziaWiabDVOeztVKTTLE6QQSL0IjE17CicJb5gibat8LWkgomsEdxSyZpOEjx18bqyPeMhmE5yn7AGUx1g3jxBXUZXMoqqf1yiblic2s&token=cztXnd9GyrHd4iajaaPHCJW4ZtfSic8HlHVaWZ1JLMFSbRsbPFnQPdNlhk11ASpWdg0Oj2CL5t1lR7Vd2yjl4Ut3lAjx5oORV6ymyfMdcpqeePq29EdWeKBSDO5MlLRZb1z48EmHEeJUJp3TiaHUq8jkAFYqvvfFuvvmia8x8PLqFR70Hjg0q21sSytBmFoT1HyW&bizid=1023&hy=SZ&idx=1&m=a5d642b207fa007e27cc99051888c93f&finder_expire_time=1767928339&wxampicformat=503&picformat=200&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAKE8Gm6qxKAAAAAstQy6ubaLX4KHWvLEZgBPEpKEQGE5SObGOzNPgMIKM6o_uZoIKXQZfrl1JnToi",
        "avatar": "http://wx.qlogo.cn/finderhead/BKDgOfSUuOXibnVOZHgcSxVATOX91VYH2QqYzYuPRTnCKln6kLlre0HoUHyianicwYAsichHibtYD4F8/0",
        "nickname": "三角洲行动",
        "desc": "还不来我直播间带你百万撤离！",
        "url": "https://channels.weixin.qq.com/web/pages/live?eid=export%2FUzFfAgtgekIEAQAAAAAAKE8Gm6qxKAAAAAstQy6ubaLX4KHWvLEZgBPEpKEQGE5SObGOzNPgMIKM6o_uZoIKXQZfrl1JnToi",
        "extras": "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",
        "objectId":"14819532136085195120",
        "objectNonceId":"15766145125716611028_0_0_0_0_0_85eb5562-e13f-11f0-b082-e9c6925e17aa"
    }
```

##### 返回示例 

``` 
{
    "data": {
        "thumb_url": "https://finder.video.qq.com/261/20304/stodownload?encfilekey=oibeqyX228riaCwo9STVsGLIia39tYEhYEEzrbjBy8M5IziaWiabDVOeztVKTTLE6QQSL0IjE17CicJb5gibat8LWkgomsEdxSyZpOEjx18bqyPeMhmE5yn7AGUx1g3jxBXUZXMoqqf1yiblic2s&token=cztXnd9GyrHd4iajaaPHCJW4ZtfSic8HlHVaWZ1JLMFSbRsbPFnQPdNlhk11ASpWdg0Oj2CL5t1lR7Vd2yjl4Ut3lAjx5oORV6ymyfMdcpqeePq29EdWeKBSDO5MlLRZb1z48EmHEeJUJp3TiaHUq8jkAFYqvvfFuvvmia8x8PLqFR70Hjg0q21sSytBmFoT1HyW&bizid=1023&hy=SZ&idx=1&m=a5d642b207fa007e27cc99051888c93f&finder_expire_time=1767928339&wxampicformat=503&picformat=200&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAKE8Gm6qxKAAAAAstQy6ubaLX4KHWvLEZgBPEpKEQGE5SObGOzNPgMIKM6o_uZoIKXQZfrl1JnToi",
        "cover_url": "https://finder.video.qq.com/261/20304/stodownload?encfilekey=oibeqyX228riaCwo9STVsGLIia39tYEhYEEzrbjBy8M5IziaWiabDVOeztVKTTLE6QQSL0IjE17CicJb5gibat8LWkgomsEdxSyZpOEjx18bqyPeMhmE5yn7AGUx1g3jxBXUZXMoqqf1yiblic2s&token=cztXnd9GyrHd4iajaaPHCJW4ZtfSic8HlHVaWZ1JLMFSbRsbPFnQPdNlhk11ASpWdg0Oj2CL5t1lR7Vd2yjl4Ut3lAjx5oORV6ymyfMdcpqeePq29EdWeKBSDO5MlLRZb1z48EmHEeJUJp3TiaHUq8jkAFYqvvfFuvvmia8x8PLqFR70Hjg0q21sSytBmFoT1HyW&bizid=1023&hy=SZ&idx=1&m=a5d642b207fa007e27cc99051888c93f&finder_expire_time=1767928339&wxampicformat=503&picformat=200&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAKE8Gm6qxKAAAAAstQy6ubaLX4KHWvLEZgBPEpKEQGE5SObGOzNPgMIKM6o_uZoIKXQZfrl1JnToi",
        "objectNonceId": "15766145125716611028_0_0_0_0_0_85eb5562-e13f-11f0-b082-e9c6925e17aa",
        "receiver": 7881302555913738,
        "extras": "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",
        "sender_name": "6666",
        "is_room": 0,
        "avatar": "http://wx.qlogo.cn/finderhead/BKDgOfSUuOXibnVOZHgcSxVATOX91VYH2QqYzYuPRTnCKln6kLlre0HoUHyianicwYAsichHibtYD4F8/0",
        "server_id": 13274471,
        "url": "https://channels.weixin.qq.com/web/pages/live?eid=export%2FUzFfAgtgekIEAQAAAAAAKE8Gm6qxKAAAAAstQy6ubaLX4KHWvLEZgBPEpKEQGE5SObGOzNPgMIKM6o_uZoIKXQZfrl1JnToi",
        "coverUrl": "https://finder.video.qq.com/261/20304/stodownload?encfilekey=oibeqyX228riaCwo9STVsGLIia39tYEhYEEzrbjBy8M5IziaWiabDVOeztVKTTLE6QQSL0IjE17CicJb5gibat8LWkgomsEdxSyZpOEjx18bqyPeMhmE5yn7AGUx1g3jxBXUZXMoqqf1yiblic2s&token=cztXnd9GyrHd4iajaaPHCJW4ZtfSic8HlHVaWZ1JLMFSbRsbPFnQPdNlhk11ASpWdg0Oj2CL5t1lR7Vd2yjl4Ut3lAjx5oORV6ymyfMdcpqeePq29EdWeKBSDO5MlLRZb1z48EmHEeJUJp3TiaHUq8jkAFYqvvfFuvvmia8x8PLqFR70Hjg0q21sSytBmFoT1HyW&bizid=1023&hy=SZ&idx=1&m=a5d642b207fa007e27cc99051888c93f&finder_expire_time=1767928339&wxampicformat=503&picformat=200&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAKE8Gm6qxKAAAAAstQy6ubaLX4KHWvLEZgBPEpKEQGE5SObGOzNPgMIKM6o_uZoIKXQZfrl1JnToi",
        "sender": 1688857368772759,
        "nickname": "三角洲行动",
        "app_info": "CIOACBCbup2ctTMYl4Gr/JuAgAMgKQ==",
        "room_conversation_id": "0",
        "sendtime": 1766632938,
        "msg_id": 1008780,
        "thumbUrl": "https://finder.video.qq.com/261/20304/stodownload?encfilekey=oibeqyX228riaCwo9STVsGLIia39tYEhYEEzrbjBy8M5IziaWiabDVOeztVKTTLE6QQSL0IjE17CicJb5gibat8LWkgomsEdxSyZpOEjx18bqyPeMhmE5yn7AGUx1g3jxBXUZXMoqqf1yiblic2s&token=cztXnd9GyrHd4iajaaPHCJW4ZtfSic8HlHVaWZ1JLMFSbRsbPFnQPdNlhk11ASpWdg0Oj2CL5t1lR7Vd2yjl4Ut3lAjx5oORV6ymyfMdcpqeePq29EdWeKBSDO5MlLRZb1z48EmHEeJUJp3TiaHUq8jkAFYqvvfFuvvmia8x8PLqFR70Hjg0q21sSytBmFoT1HyW&bizid=1023&hy=SZ&idx=1&m=a5d642b207fa007e27cc99051888c93f&finder_expire_time=1767928339&wxampicformat=503&picformat=200&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAKE8Gm6qxKAAAAAstQy6ubaLX4KHWvLEZgBPEpKEQGE5SObGOzNPgMIKM6o_uZoIKXQZfrl1JnToi",
        "msgtype": 146,
        "feed_type": 9,
        "objectId": "14819532136085195120",
        "desc": "还不来我直播间带你百万撤离！"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 3.12 发送gif表情 [已封装: SendEmotionMessage]

##### 简要描述

- 发送gif表情

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendEmotionMessage
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数


|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
    "uuid":"41cba32733a3643f7bdc4cbdd360d797",
    "send_userid":7881302555913738,
    "isRoom":false,
    "imgurl":"https://wework.qpic.cn/wwpic3az/wwwx_954b864d1be1d3afbf6fd6581c1903c1/0",
    "width":240, //宽度
    "height":240,//高度
    "emotionType":0
}
```

##### 返回示例 

``` 
{
    "data": {
        "receiver": 7881302555913738,
        "sender": 1688853790599424,
        "EmotionType": "EMOTION_STATIC",
        "sender_name": "",
        "is_room": 0,
        "sendtime": 1652169938,
        "msg_id": 1066247,
		"app_info":"from_msgid_697590829808721944",
        "server_id": 12381211,
        "msgtype": 29,
        "url": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimglf3.nosdn0.126.net%2Fimg%2FWEFHU2lQKzFVM1lGbWFTcG84YXdudFl2dnhPN3ZSY0g4UFRzSlpzMjZ1TGs5MEF4YTRYbjN3PT0.gif&refer=http%3A%2F%2Fimglf3.nosdn0.126.net&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1654435655&t=4ef360eed30a307a951785967a097030"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 3.13 发送名片消息 [已封装: SendBusinessCardMsg]


##### 简要描述

- 发送名片消息

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendBusinessCardMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
    "uuid":"6d92d9e5-88fd-4d9f-8007-3e81f57a35a2",
    "send_userid":7881302555913738, //要给谁发名片
	"headImg":"http://www.baidu.com/dxxxx.jpg",
    "isRoom":false,
    "username":7881302555913738, //被发的名片人id
    "nickname":"被发送人的昵称",
    "enterpriseName":"企业名称"
}
```

##### 返回示例 

``` 
{
    "data": {
        "receiver": 7881302555913738,
        "headImg": "",
        "sender_name": "",
        "is_room": 0,
        "server_id": 12381212,
        "sender": 1688853790599424,
        "nickname": "",
        "sendtime": 1652169970,
		"app_info":"from_msgid_697590829808721944",
        "msg_id": 1066249,
        "msgtype": 41,
        "enterpriseName": "",
        "corp_id": 0,
        "username": 7881302555913738
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 3.14 发送位置消息 [已封装: SendLocationMsg]


##### 简要描述

- 发送位置消息

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendLocationMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人或群id|
|isRoom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
    "uuid":"6d92d9e5-88fd-4d9f-8007-3e81f57a35a2",
    "send_userid":7881302555913738,
    "isRoom":false,
    "longitude":115.54214,
    "latitude":34.82553,
    "address":"山东菏泽曹县",
    "detailed_address":"山东曹县牛B666我的宝贝儿"
}
```

##### 返回示例 

``` 
{
    "data": {
        "detailed_address": "山东曹县牛B666我的宝贝儿",
        "address": "山东菏泽曹县",
        "receiver": 7881302555913738,
        "sender": 1688853790599424,
        "latitude": 34.82553,
        "sender_name": "",
		"app_info":"from_msgid_697590829808721944",
        "is_room": 0,
        "sendtime": 1652169999,
        "msg_id": 1066251,
        "server_id": 12381213,
        "msgtype": 6,
        "longitude": 115.54214
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 3.15 发送群@消息 [已封装: SendTextAtMsg]


##### 简要描述

- 发送群聊@消息

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendTextAtMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的群id|
##### 请求示例 

``` 
{
    "uuid":"603fc22b-2e0e-4faa-a03c-8880733d1228",
    "send_userid":10696052955013024,//要发送的id
    "atids":[
        7881302555913738  //@群成员id
    ],
    "content":"ddddddd", //内容 //发送内容
	"isRoom":true //群true  
}
```

##### 返回示例 

``` 
{
    "data": {
        "receiver": 0,
        "sender": 1688853790599424,
        "at_list": [
            {
                "user_id": 7881302555913738,
                "nikeName": ""
            }
        ],
        "sender_name": "",
		"app_info":"from_msgid_697590829808721944",
        "room_conversation_id": 10696052955013024,
        "is_room": 1,
        "sendtime": 1654572835,
        "msg_id": 1068996,
        "server_id": 12480017,
        "msgtype": 2,
        "content": "ddddddd"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```
### 3.16 同步消息记录 [已封装: SyncAllData]

##### 简要描述

- 同步消息记录，用来同步消息记录，如果电脑微信在离线期间，手机端进行了消息操作，调用这个接口可以把离线期间的消息同步过来。
使用场景，登录成功的时候使用，然后就不需要用到它了，如果对消息没那么重要则不需要用它
##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SyncAllData
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|limit |是  |int |每次返回大小|
|seq |是  |int |查询下标|
``` 
建议seq传消息回调中的最后一次的server_id,下发中的消息回调，每个消息都会带
着server_id，这个id就是这里面的seq，建议每次接收那边的消息，就更新一下server_id，然后如果离线，过了几天，再次用
协议登录，就直接传递上次的server_id就会直接同步从上次最后一次保存server_id为开始，进行同步。
```
如下图的server_id，就是这里的seq
![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=bf1f7393427b7313088dc6822af35733)
##### 请求示例 

``` 
{
    "uuid":"ecb033af-6fcd-4ec2-880e-41f070b65eaf",
    "limit":1000, 
    "seq":12481627
}
```

##### 返回示例 

``` 
{
    "data": {
        "is_select": 1, //这个参数是判断是否还需要进行消息遍历 1代表没有遍历完还需要进行遍历 0则是遍历完成不需要进行下一步的遍历
        "seq": 12481628,//下次调用传这个seq，和其他功能接口的seq同一个意思。
        "content": "消息记录查询完成，已经分发到消息回调中，请在回调中处理离线消息"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 3.17 群发消息 [已封装: SendGroupsMsg]


##### 简要描述

- 控制台群发消息

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SendGroupsMsg
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|vids |是  |long |要发送的人或者群idid|
|isroom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
	"uuid":"xxxxxxxx",
    "vids": [
        "1688851883777952"// 群id或者好友id 不能混填
    ], 
    "isroom": false, //是否发送给群 true是 false否
    "msg_list": [
        {
            "type": 0, //文字
            "content": "群发消息[微笑]"
        }, 
        {
            "type": 14, //图片
            "aeskey": "63636533323765343938636532633264", 
            "cdnkey": "30819202010204818a3081870201000204ea44290002030f55cb0204e2a91e6f0204643d4475044c4e45574944315f65613434323930306532613931653666363433xx981d4320d0c4af27774c2d6de31781e90201010201000400", 
            "fileSize": 24389, 
            "md5": "981d4320d0c4af27774c2d6de31781e9"
        }, 
        {
            "type": 13, //连接
            "url": "https://www.baidu.com/?tn=02003390_10_hao_pg", 
            "title": "百度一下，你就知道", 
            "content": "全球领先的中文搜索引擎、致力于让网民更便捷地获取信息，找到所求。百度超过千亿的中文网页数据库，可以瞬间找到相关的搜索结果。", 
            "headImg": "https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/topnav/newfanyi-da0cea8f7e.png"
        }, 
        {
            "type": 78, //小程序
            "title": "测试标题", 
            "desc": "测试简介", 
            "weappIconUrl": "http://mmbiz.qpic.cn/mmbiz_png/BMAjdIR5qwChFiaRnxUTiadJicGwMpqchXVdfQQ67tHTzfgd0gkibupgeOKtHtyE9m9SZopHXxcpsSicWxzLKXwk0qA/640?wx_fmt=png&wxfrom=200", 
            "appid": "wx45dff5234240ad90", 
            "username": "gh_c4a2a98a7366@app", 
            "pagepath": "pages/train/index/index.html", 
            "cdnkey": "30819202010204818a30818702xx981d4320d0c4af27774c2d6de31781e90201010201000400", 
            "md5": "981d4320d0c4af27774c2d6de31781e9", 
            "aeskey": "63636533323765343938636532633264", 
            "fileSize": 24389
        }, 
        {
            "type": 15, //文件
            "fileName": "7014419642333305662.MP4", 
            "aeskey": "31343031353436666630396132346638", 
            "cdnkey": "30819302010204818b30818802010xx3b9399f4489989e2bb886fe35fce4df40201050201000400", 
            "fileSize": 753510, 
            "md5": "3b9399f4489989e2bb886fe35fce4df4"
        }, 
        {
            "type": 23, //视频
            "aeskey": "39616335386438356162623034623038", 
			"md5": "xxxxxxxxxxx", 
            "VideoDuration": 5, 
			"video_img_size"6000,// 视频封面图大小
            "cdnkey": "30690201020462306002010xx393746383402010002030b7f7004103b9399f4489989e2bb886fe35fce4df40201040201000400", 
            "fileSize": 753510
        }
    ]
}
```

##### 返回示例 

``` 
{
    "data": {
        "msg_id": 1066230
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 3.18 语音转文字 [已封装: SpeechToTextEntity]


##### 简要描述

- 语音转文字

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/SpeechToTextEntity
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|msgid |是  |int |要转文字的语音消息id|
##### 请求示例 

``` 
{
    "uuid":"a4ea6a39-4b3a-4098-a250-2a07bef57355",
    "msgid":1063645
}
```

##### 返回示例 

``` 
{
    "data": {
        "text": "牛牛牛牛。"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```


### 3.19 消息已读 [已封装: MarkAsRead]

##### 简要描述

- 已读消息

##### 请求URL
- ` http://127.0.0.1:8083/wxwork/MarkAsRead
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |每个实例的唯一标识，根据uuid操作具体企业微信|
|send_userid |是  |long |要发送的人id|
|isRoom |是  |bool |是否是群消息|
##### 请求示例 

``` 
{
    "uuid":"1753cdff-0501-42fe-bb5a-2a4b9629f7fb",
    "send_userid":7881302555913738, //好友或者群id
    "isRoom":false
}
```

##### 返回示例 

``` 
{
    "data": {
        "receiver": 0,
        "sender": 1688857826547606,
        "sender_name": "",
        "app_info": "CIOACBDHsODn4jEYlq/P1p2AgAMgAw==",
        "room_conversation_id": "10973555009565180",
        "is_room": 1,
        "sendtime": 1710151440,
        "msg_id": 1013548,
        "server_id": 6306123,
        "msgtype": 2001
    },
    "errcode": 0,
    "errmsg": "ok"
}
```


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

## 七、CDN上传和大文件上传下载 (cdn.go)

### 大文件

#### 7.1 大文件网络上传

##### 简要描述

- 大文件网络上传

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/BigFileUploadLink
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|url |是  |String |网络图片地址|

##### 请求示例 

``` 
{
    "uuid":"2b0863724106a1160212bd1ccf025295",
    "authkey":"0AAxxx031", 
    "filekey":"346b7bff-08d5-4ac2-bc67-fd10e3eb2388", 
    "fileurl":"https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png", 
    "filename":"PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png"
}
```

##### 返回示例 

``` 
{
    "data": {
        "file_key": "346b7bff-08d5-4ac2-bc67-fd10e3eb2388",
        "filename": "PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",
        "file_id": "*1*u3xjYZlKKSNgRNiv8t7wd0BIOlYGMxmq7n4lwdvaqejMvIiCf2jBDQDssuXZW5zJBwGAfzFMX5C2xfdkh6r47BUyl0VoX1XVH9eAOgAj2jYLna64xaYd9ZxU0GlvhbRc7Df9zHzG24m4G5y86R/cNUgfg+XkJbS+MHiqJUuqHJiKGOXH7R2r3LbxVgy2LUyJS3U5qDeM=",
        "file_size": 15444,
        "md5": "d9c8750bed0b3c7d089fa7d55720d6cf"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

#### 7.2 大文件上传


##### 简要描述

- 大文件上传

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/BigUploadFile
  
##### 请求方式
- POST 

- ContentType:"multipart/form-data"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |uuid|
|authkey |是  |String |authkey|
|filekey |是  |String |filekey|
|file |是  |File |aeskey|
##### 请求示例 

![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=a28d49649c0b63249c926bd6861a1542&file=file.png)

##### 返回示例 

``` 
{
    "data": {
        "file_key": "1568ef8b-7243-420d-832b-1xxxxxf",
        "filename": "2f0ec2d938xxxx9cb0da022.zip",
        "file_id": "*1*F34hSNMREYF+BNgTfNtYo7Lg5/va+EoTPkyESrZSpwtNs1IFEpFfXWbWXF+kvx+rrm7SUtsNBbdEUxBcMc5zyqXZ3Z1pT7CShVKq1LmjzbhtG3ZAK/xvvWpUoq2ynfnC765zusYDzYEiH5bbc4g=",
        "file_size": 314065,
        "md5": "9c2e838x1"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```


#### 7.3 大文件下载


##### 简要描述

- 大文件下载

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/bigFileDownUrl
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |账号uuid|

##### 请求示例 

``` 
{
    "uuid": "b31d6790252c7d5df6d3273bd2636f95",
    "fileid": "*1*JfLw/9mM4Z6IeL8sTcOta79sgemyuuRLWOLWjNvB8+LWRGWCYESpv7IVAvv4Dz3ps+m+jrhjUMhJ5dyD8Q0MOQ7Y8Lf1WhFzKd9Pm/1zHbsachObBrPjRmwjtbBbYJEq4499gHSuWf0p3GRAwq1IPDHslZQxjJQFxXtM+QAk/lX6eaOaETG4KmQfrrUHRcUOyzBsG6+PpwCjrWfYJKizYuUFqz65kO0N63QwRHbq0VF+1fMiJgGT/jg145aHffGqBTiC5Rzzq0CDH0Fa/N8jav6yGW73UbpTWanEYduYiZT8Jh1BjmC0tvCGv3ps1n3xA0vhClxO1CGaJLS+tEsaOa0RSiP1R7C833EfWfX36OwSa3TsOEA/gjNteKESe9qEYXuJkAMGh7Cbhey+gti1LnNSUWahbzV7uGJzYk8dOBzqw2Hiz3HQ6MiuqgvAUVBmoMsxy8tHpxrscbleLVqunsmdywBZXEBe42ZWbbQRjPE=",
    "file_name": "666.mp4",
    "size": 371275
}
```

##### 返回示例 

``` 
{
    "data": {
        "url": "http://121.40.87.201:8060/download/File/20250421/1451ebe03322ea9b665e37b5696a96b7.mp4",
        "md5": "CF4648FBBA449A1F1477C25944515619"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```


#### 7.4 大文件下载参数获取


##### 简要描述

- 如果嫌弃自带的大文件下载速度过慢，可以直接用这个接口拿到下载所需参数 自己去post请求下载

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/getBigFileDownData
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String |账号uuid|

##### 请求示例 

``` 
{
    "uuid": "b31d6790252c7d5df6d3273bd2636f95",
    "fileid": "*1*JfLw/9mM4Z6IeL8sTcOta79sgemyuuRLWOLWjNvB8+LWRGWCYESpv7IVAvv4Dz3ps+m+jrhjUMhJ5dyD8Q0MOQ7Y8Lf1WhFzKd9Pm/1zHbsachObBrPjRmwjtbBbYJEq4499gHSuWf0p3GRAwq1IPDHslZQxjJQFxXtM+QAk/lX6eaOaETG4KmQfrrUHRcUOyzBsG6+PpwCjrWfYJKizYuUFqz65kO0N63QwRHbq0VF+1fMiJgGT/jg145aHffGqBTiC5Rzzq0CDH0Fa/N8jav6yGW73UbpTWanEYduYiZT8Jh1BjmC0tvCGv3ps1n3xA0vhClxO1CGaJLS+tEsaOa0RSiP1R7C833EfWfX36OwSa3TsOEA/gjNteKESe9qEYXuJkAMGh7Cbhey+gti1LnNSUWahbzV7uGJzYk8dOBzqw2Hiz3HQ6MiuqgvAUVBmoMsxy8tHpxrscbleLVqunsmdywBZXEBe42ZWbbQRjPE=",
	"file_name": "666.mp4",
    "size": 371275
}
```

##### 返回示例 

``` 
{
    "data": {
        "isBigC2c": true,
        "authCookie": "weixinnum=1688857368772759&authkey=700800100218702270f3ed0fd7ab7833e82665b913eaddd6c299b99ed6003c9b7419bb7d9e879c5b71b853dca1cd4887ba832235dae526f684ff8dd74d2a89d4fff73c7296bea52bad0d04763636ea982cf43b0afef8295c5f2c6c8194eb9d867cd36d0c930fad626c958686749472a1438e9201ced3cbe50b", //下载大文件需要的cookie
        "errorMessage": "SUCCESS",
        "errorCode": 0,
        "url": "https://shfront.wxwork.qq.com:443/downloadobject?fileid=08011204313330312210313638383835343739383731313939382a023335322438383132383837382d633032302d346635372d626335652d65353738653464313137663338fdacc001421413a385153c6c02ef3908799da1da940a299846a748015802600768e807720731303030303030900192a6e2c3069a0100a001eaa102&weixinnum=1688857368772759&authkey=700800100218702270f3ed0fd7ab7833e82665b913eaddd6c299b99ed6003c9b7419bb7d9e879c5b71b853dca1cd4887ba832235dae526f684ff8dd74d2a89d4fff73c7296bea52bad0d04763636ea982cf43b0afef8295c5f2c6c8194eb9d867cd36d0c930fad626c958686749472a1438e9201ced3cbe50b&filename=3AmhwLPerR.jpg",//下载大文件用到的url地址 post请求
        "md5": "34633932306539336339613966333666",
		"size": 371275//文件大小 就是你传的参数
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

##### 返回参数请求使用方式

![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=baff00daea59b13487b1b399fb3433c5&file=file.png)



### 图片

#### 7.5 cdn上传网络图片


##### 简要描述

- CDN上传网络图片重置版

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/CdnUploadImgLink
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|url |是  |String |网络图片地址|

##### 请求示例 

``` 
{
    "uuid":"094d0f95d60e8f1404c83114608be5f3",
    "url":"http://xxxx:8060/download/Img/20250331/66666.png"
}
```

##### 返回示例 

``` 
{
    "data": {
        "cdn_key": "306b020102046430620201000204060b19e102034c4cd20204986b259902046807456b042466303163353336302d636437642d346335302d626238632d38356638333361613238376402031038000203317d8004107d8a8e79499210c6242c36afd10647c70201010201000400",
        "fileid": "306b020102046430620201000204060b19e102034c4cd20204986b259902046807456b042466303163353336302d636437642d346335302d626238632d38356638333361613238376402031038000203317d8004107d8a8e79499210c6242c36afd10647c70201010201000400",
        "aes_key": "B26DB4824341608DB848A2DA563B7147", //加密key
        "md5": "7d8a8e79499210c6242c36afd10647c7", //md5
        "width": 1279, //宽度
        "height": 1706,//高度
        "size": 3243381,//大小
        "mid_image_size": 164550, //中图大小
        "thumb_image_width": 384,//缩略图宽度
        "thumb_image_height": 512,//缩略图高度
        "thumb_file_size": 15195,//缩略图大小
        "thumb_file_md5": "4331947dc5f67c0d8a2893653a3f0cee",//缩略图md5
        "retcode": "0"
    },
    "errcode": 0,
    "errmsg": "成功"
}
```


#### 7.6 cdn上传本地图片


### 视频

#### 7.7 cdn上传网络视频

##### 简要描述

- CDN上传网络视频

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/UploadCdnVideoLink
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|url |是  |String |网络图片地址|

##### 请求示例 

``` 
{
	"uuid":"094d0f95d60e8f1404c83114608be5f3",
    "url":"http://47.94.7.218:8060/Video/1.mp4", //小于25m 超过走大文件上传和发送
    "filename":"111111.mp4",
    "video_img_url":"https://www.baidu.com/img/flexible/logo/pc/result.png"
}
```

##### 返回示例 

``` 
{
    "data": {
        "video_height": 1170, //视频高度
        "size": 3561782,
        "VideoDuration": 38,
        "cdn_key": "306b020102046430620201000204060b19e102034c62b20204170a3cda0204681b245f042464623337376435332d303937652d346639322d616435392d6437383163303835396432360203100004020336594004109c6131377dae038824b43f8e388efa1c0201040201000400",
        "aes_key": "D3F9FE26E64F1806ED5F21CC3B51D31D",
        "video_width": 540,//视频宽度
        "video_img_size": 65472,
        "fileid": "306b020102046430620201000204060b19e102034c62b20204170a3cda0204681b245f042464623337376435332d303937652d346639322d616435392d6437383163303835396432360203100004020336594004109c6131377dae038824b43f8e388efa1c0201040201000400",
        "md5": "9c6131377dae038824b43f8e388efa1c",
        "retcode": "0"
    },
    "errcode": 0,
    "errmsg": "成功"
}
```

#### 7.8 cdn上传本地视频 todo


### 文件和语音skilk格式

#### 7.9 cdn上传网络文件

##### 简要描述

- CDN上传网络文件

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/UploadCdnLinkFile
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|url |是  |String |网络图片地址|

##### 请求示例 

``` 
{
	"uuid":"094d0f95d60e8f1404c83114608be5f3",
    "url":"http://47.94.7.218:8060/File/222.txt",
    "filename":"222.txt"
}
```

##### 返回示例 

``` 
{
    "data": {
        "size": 16,
        "cdn_key": "30670201020460305e0201000204dad3d79602030f55cb0204e2a91e6f020465251e0a042462313665323862372d366431312d346638612d383830652d3132366665633535656363390201000201100410b59c67bf196a4758191e42f76670ceba0201050201000400",
        "aes_key": "3BF8F1F2D903522FCE797EDDEC7D892B",
        "fileid": "30670201020460305e0201000204dad3d79602030f55cb0204e2a91e6f020465251e0a042462313665323862372d366431312d346638612d383830652d3132366665633535656363390201000201100410b59c67bf196a4758191e42f76670ceba0201050201000400",
        "md5": "b59c67bf196a4758191e42f76670ceba"
    },
    "errcode": 0,
    "errmsg": "成功"
}
```

#### 7.10 cdn上传本地文件  todo


###  cdn下载和个v文件下载

#### 7.11 cdn下载文件

##### 简要描述

- CDN下载文件重置版本

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/DownloadFile
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String | 对应账号uuid|
|fileid |是  |String |文件id|
|aes_key |是  |long |aes_key|
|filetype |是  |bool |下载文件类型如果内部图片消息中is_hd字段是1就传1是2就传2 3缩略图 4视频 5文件语音|
|file_name |是  |bool |文件名|
|size |是  |bool |文件大小|
##### 请求示例 

``` 
{
	"uuid":"094d0f95d60e8f1404c83114608be5f3",
    "fileid":"306902010204623x734363561320201000203184c7004107dc7b1a27f9aa62c26b0790cc9cd859b0201040201000400",
    "aes_key":"376A54732B7F4D3EE313FB09AC1A6A1E",
    "filetype":4, //4 视频 5文件语音也走这个 1图片 3缩略图 传入视频消息的cdnkey 
    "file_name":"测试.mp4", //文件名必填
    "size":1592432  //文件大小如果是3类型的则在视频消息中找视频缩略图大小 有个open_xxx_size的字段。 别传视频大小会失败。
}
```

##### 返回示例 

``` 
{
    "data": "http://127.0.0.1:8060/Video/测试.mp4",
    "errcode": 0,
    "errmsg": "成功"
}
```

#### 7.12 外部联系人图片视频文件下载


##### 简要描述

- 外部联系人文件下载

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/DownloadWeChatFile
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uuid |是  |String | 对应账号uuid|
|url |是  |String |CDNkey|
|auth_key |是  |String |认证key|
|aes_key |是  |String |aes_key|
|file_name |是  |String |文件名称|
|size |是  |int |文件大小|
##### 请求示例 

``` 
{
	"uuid":"094d0f95d60e8f1404c83114608be5f3",
    "url": "https://imunion.weixin.qq.com/cgi-bin/mmae-bin/tpdownloadmedia?param=v1_e80c6c6c0cxxxx3544d9c9084a8fe4dd608a9fe46e0091d513f47c49eb0536adecd48545f28732f2c5874d70ea579686ebaf26904df5b2d51dba74960afb7c218789c47cdcb1cdec938955de8d4851b2a8cb3fdd7a0ab9310d18a94e78c297719cdd65dd36ea5daead3b", 
    "aes_key": "afb5844dx983bfb8d4", 
    "auth_key": "v1_e80c6c6cxxxxxd29bac49a77ef1e4dcc51b6f8143eb67f23b4", 
    "file_name": "test.jpg", 
    "size": 79992
}
```

##### 返回示例 

``` 
{
    "data": "http://127.0.0.1:8084/Img/39098ac9b48106669cbf01ccd998b8a3.MP4",
    "errmsg": "ok",
    "errorcode": 0
}
```

---

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

## 九、账号基础操作 (account.go)

### 9.1 获取当前账号详情 [已封装: GetProfile]
##### 简要描述

- 获取当前账号详情

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetProfile
  
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
    "uuid":"f5a22e9b-9664-4250-b40a-08741dba549c"
}
```

##### 返回示例 

``` 
{
    "data": {
        "unionid": "ozynqskEsd4gEtIJpxXB_YPu1T4o",
        "create_time": 0,
        "sex": 1,
        "mobile": "13573753761",
        "corp_full_name": "",
        "acctid": "LiuXianKun",
		"scorp_id": "ww6509c50a6839c841",
        "avatar": "https://wework.qpic.cn/wwpic/839179_8l0dm09iQ0qV76v_1652237791/0",
        "corp_name": "明月科技",
        "english_name": "",
        "realname": "刘xxx",
        "user_id": 1688853132666001,
        "nickname": "刘xxx",
        "position": "",
        "corp_desc": "",
        "corp_id": 1970325105140387
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 9.2 获取当前账号二维码 [已封装: GetThisQRCode]
##### 简要描述

- 获取当前账号二维码

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetThisQrCode
  
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
    "uuid":"f5a22e9b-9664-4250-b40a-08741dba549c"
}
```

##### 返回示例 

``` 
{
    "data": {
        "QrCodeUrl": "https://wework.qpic.cn/wwpic/614600_ky2A2t1FRfKmue5_1652238985/0"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 9.3 获取当前账号对外名片二维码 [已封装: ExternalBusinessCard]

##### 简要描述

- 获取当前账号对外名片二维码

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/ExternalBusinessCard
  
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
    "uuid":"f5a22e9b-9664-4250-b40a-08741dba549c"
}
```

##### 返回示例 

``` 
{
    "data": {
	//base64二维码图片
        "qr_image": "iVBORw0KGgoAAAANSUhEUgAAAxgAAAMYCAYAAABIfUFkAAAgAElEQVR4Xu3d53U1VZYwYF0t/aeZBOj+EsAkgAmgMQlgAsAF0JgAMBMApgPABQD9JoBJAIYJYGgSuPfTFi0Qkq5UZlfVMc9di0XPULVrn2cfSbXvKbM7nH9OfAgQIECAAAECBAgQIDBf4HA6P4YIBAgQIECAAAECBAgQ+E1Ag2EmECBAgAABAgQIECCQJqDBSKMUiAABAgQIECBAgAABDYY5QIAAAQIECBAgQIBAmoAGI41SIAIECBAgQIAAAQIENBjmAAECBAgQIECAAAECaQIajDRKgQgQIECAAAECBAgQ0GCYAwQIECBAgAABAgQIpAloMNIoBSJAgAABAgQIECWlkzgBAgQIECBAgACB8gQ0GOXVREYECBAgQIAAAQIEqhXQYFRbOokTIECAAAECBAgQKE9Ag1FeTWREgAABAgQIECBAoFoBDUa1pZM4AQIECBAgQIAAgfIENBjl1URGBAgQIECAAAECBKoV0GBUWzqJEyBAgAABAgQIEChPQINRXk1kRIAAAQIECBAgQKBagd3h/FNt9hInQIAAAQIECBAgQGCIQJzz74ZsOHObgxWMmYJ2J0CAAAECBAgQIEDgDwENhtlAgAABAgQIECBAgECagAYjjVIgAgQIECBAgAABAgQ0GOYAAQIECBAgQIAAAQJpAhqMNEqBCBAgQIAAAQIECBA4a4Fgt5t2Q/yxB2hlxyvduPTxZueXHS+7vtn5lR4v26/0eNn1yB5vb/llj7e3eNnzr/R4pdc322/KeGOf/X6fnUr38eKc9fR0/LrAVvUYn2n3JQZAgAABAgQIECBAgMAxAQ2GuUGAAAECBAgQIECAQJqABiONUiACBAgQIECAAAECBDQY5gABAgQIECBAgAABAmkCGow0SoEIECBAgAABAgQIENBgmAMECBAgQIAAAQIECKQJaDDSKAUiQIAAAQIECBAgQECDYQ4QIECAAAECBAgQIJAmoMFIoxSIAAECBAgQIECAAAENhjlAgAABAgQIECBAgECagAYjjVIgAgQIECBAgAABAgQ0GOYAAQIECBAgQIAAAQJpAhqMNEqBCBAgQIAAAQIECBDYHc4/GAgQIECAAAECBAgQaFogzvl3K4zwYAVjBWWHIECAAAECBAgQINCLgAajl0obJwECBAgQIECAAIEVBDQYKyA7BAECBAgQIECAAIFeBDQYvVTaOAkQIECAAAECBAisIKDBWAHZIQgQIECAAAECBAj0InC2xUB3u2k3sB974FXp8bYwHnPM3vyyxzvGuoVts/1Kj1d6zbL9jLd0gXn5Zc+X0uPN01p+79L9puQX++z3+xt4cQ53ejr+e23xfqPM9lt6do+v9NIZiU+AAAECBAgQIECAQLUCGoxqSydxAgQIECBAgAABAuUJaDDKq4mMCBAgQIAAAQIECFQroMGotnQSJ0CAAAECBAgQIFCegAajvJrIiAABAgQIECBAgEQIajPJqIiMCBAgQIECAAAEC1QpoMKotncQJECBAgAABAgQIlCegwSivJjIiQIAAAQIECBAgUK2ABqPa0kmcAAECBAgQIECAQHkCGozyaiIjAgQIECBAgAABAtUKaDCqLZ3ECRAgQIAAAQIECJQnoMEoryYyIkCAAAECBAgQIFCtgAaj2tJJnAABAgQIECBAgEB5AhqM8moiIwIECBAgQIAAAQLVCuwO559qs5c4AQIECBAgQIAAAQJDBOKcfzdkw5nbHKxgzBS0OwECBAgQIECAAAECfwhoMMwGAgQIECBAgAABAgTSBDQYaZQCESBAgAABAgQIECCgwTAHCBAgQIAAAQIECBBIEzhLiyQQAQIECGwu8D//8z8n3"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 9.4 设置头像 [已封装: AddImage]
- **请求URL**: `/wxwork/AddImage`
- **参数**: uuid, url(图片链接)

### 9.5 设置用户信息 [已封装: UpdateUserInfo]


##### 简要描述

- 设置用户信息

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/UpdateUserInfo
  
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
    "uuid":"6d25981b-071c-4e1e-aa01-5770d7a6559d",
    "emile":"1693590929@qq.com", //邮箱
    "name":"name嗷嗷",//名字
    "english_name":"yingwenmingcheng",//英文名
    "phone":"222",//座机
    "position":"232",//职位
    "mobile":"12222222",//手机号
    "gender":1//性别 1男2女
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



### 9.6 设置是否自动同意 [已封装: AutomaticConsent]


##### 简要描述

- 设置是否自动同意

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/AutomaticConsent
  
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
    "uuid":"bc48e30d0c00492083fdec4c1a7e5c94",
    "state":1 //1 是需要验证同意（需要手动点击同意） 0关闭验证 自动同意好友 
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



### 9.7 根据企业id获取企业详情 [已封装: GetCorpInfo]


##### 简要描述

- 根据i企业d获取企业详情

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetCorpInfo
  
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
    "uuid":"e1d672c613717ca41fcbd6823da9015a",
    "corpid": 1970324993045436
}
```

##### 返回示例 

``` 
{
    "data": {
        "list": [
            {
                "Staffnum": 0,
                "Auth_time": 1743673059,
                "Card_url": "https://work.weixin.qq.com/apph5/user/h5/corp?",
                "Corpid": 197032499x45xxx,
                "Admin_vid": 16888522833xx8,
                "V_superadmin_vid": 16888522xx208,
                "Corp_stat": "AUTH_CORP",
                "Corp_full_name": "xxxx有限公司",
                "IsAccepted": false,
                "Create_time": 1534914566,
                "Vid": 1688857345791146,
                "Scorp_id": "ww84bf86xxxxx",
                "Mail": "",
                "Ownername": "佘xxx",
                "Corp_desc": "",
                "Corp_name": "xxxxx",
                "Corp_logo": "https://p.qlogo.cn/bizmail/vhpBKorAa29UibBkcysHzF1M93Yq1iaEuMJf9tHk1DNCoH7gjC2aLOUg/0"
            }
        ]
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 9.8 获取当前账号封禁原因 [已封装: GetBanDetail]


##### 简要描述

- 获取当前账号封禁原因以及解封时间
![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=3b19dbf5f8f84ffde264ce6c623ec0a0&file=file.png)
![](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=0662b87ce22afdd8f4eabbaa032946c6&file=file.png)

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/getBanDetile
  
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
    "uuid": "d3062b1658696330c06de618b52e8f95",
    "url":"https://work.weixin.qq.com/webapp/spam/info?key=5lnY61CGjPFUvYNwLanEklEzH2H2kv9ndXtnqnOlgfMNWq4tPE8qRYS5nGoew4WXio%2FAQMui9qhbDFDApfo61XnSbW9S8OwA0K2CZb290XCD%2Fj%2FI0jI4GkF5fomdypWnTXgiw4jzjqGZaeCvNK6KlwAGVzR1C9sP7ITa36Nt4O7zj3598FQBlbkRXFOMbDESnkpqi1%2F4NSjSxaO4NmuJxQ%3D%3D&type=3&need_login=true" //从系统消息回调中获取
}
```

##### 返回示例 

``` 
{
    "data": {
        "vid": "168885628xxxx2",
        "gid": "225180304xxxx116",
        "mainData": {
            "ban_start_time": "1745647588", //开始封禁时间
            "appeal_id": "R9Eb5Sui7oDmQvbI79j1xNlXe9DBMRfy1i5Oee6lKTBX1wM2zbsM7iPUrYHDTSXS3GJd1Ohw0AUGqFbpzkCL/+VX1uz2GeR8S3hhFqY3qEFBnEx7iUTPQ5SQ4m7uh0Th",
            "ban_source": 0,
            "is_show_title": true,
            "err_code": 0,
            "ban_title": "你的部分功能已被限制使用", //封禁标题
            "type": "3",
            "is_show_reason": true,
            "is_show_func": true,
            "ban_finish_time": 1746252388, //结束封禁时间(number)
            "ban_category": 2,
            "appeal_status": 0,
            "can_appeal": true,
            "ban_id_type": 1,
            "need_login": "true",
            "ban_reason": "涉嫌存在骚扰/恶意引流/欺诈等违规行为", //封禁原因
            "template_type": 0,
            "exposed_id": "v008b6bbbbd9",
            "ban_func": "新建客户群；邀请客户加入群聊；添加客户；客户聊天；客户朋友圈；群发助手", //封禁功能
            "is_show_time": true,
            "can_urge": false,
            "key": "5lnY61CGjPFUvYNwLanEklEzH2H2kv9ndXtnqnOlgfMNWq4tPE8qRYS5nGoew4WXio/AQMui9qhbDFDApfo61XnSbW9S8OwA0K2CZb290XCD/j/I0jI4GkF5fomdypWnTXgiw4jzjqGZaeCvNK6KlwAGVzR1C9sP7ITa36Nt4O7zj3598FQBlbkRXFOMbDESnkpqi1/4NSjSxaO4NmuJxQ=="
        },
        "isWechat": false,
        "isWxwork": true,
        "is_superadmin": 0,
        "order_id": "R9Eb5Sui7oDmQvbI79j1xNlXe9DBMRfy1i5Oee6lKTBX1wM2zbsM7iPUrYHDTSXS3GJd1Ohw0AUGqFbpzkCL/+VX1uz2GeR8S3hhFqY3qEFBnEx7iUTPQ5SQ4m7uh0Th",
        "corp_id": "197032xxxx045"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

---

## 十、朋友圈操作 (moments.go)

### 10.1 获取朋友圈列表 ✅ `GetSnsList`

##### 简要描述

- 获取朋友圈列表

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetSnsList
  
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
    "uuid":"168885379xxxx",
    "seq":-1,
    "limit":10
}
```

##### 返回示例 

``` 
{
    "data": [
        {
            "flag": 0,//0是未删除  1是被删除
            "task_sid": 0,
            "postid": "CIOACBDL4reugTIYhL33qpWAgAMgDg==",
            "media": [ //图片视频
                {
                    "size": 90092,
                    "width": 600,
                    "video_len": 0,
                    "store_type": 0,
                    "type": 1, //文件类型 1 图片 2视频 3视频封面
                    "fileid": "*1*viLte4iMW9/VC5OZ/ZZ7FuE9VD+AhmsNxWKoL+WZI6sIZZBgOJs90NPaP6WefVRdiM+JCAfcHuTXGiawQgAn3rGdHaXkWIJpytUgbBR4zT1vxhvnwdgxHfmhPtzN1zYe+9+aYNHW6c5EAS1WWHb5XWTO33m2zgrLlR17Abv1fGgg546umsOe1EpNNYUFYf4V/0F50lLI3KrfzRlzup9V7KDKN2JsX2Hfm/18caLvy03+EpcXCI+jTUIVoZCZ5jdJSzzW2QhsC/NPE/QVOkE0XkhsI0+5P/nsHgqwjAEUzY89hDyxQO8BCvlgiYmsXRWlJ0x/HNBbeOsUdwq8R3vsCBGffWiJDsMZ5WYfgvb0DL7oMuUvzVxYfUoxtq3J66lmBiUZTFenlJYAmzaHqJlHWnFEMSnK/DfWOr2jiZvFKV+6z9JXn7O1yrY725OF13asLOu9pJfKQDfNbLDICpL4M7TUOguTXfyGTqo7dOR2H60=",
                    "md5": "29a647af6dc7ec08653aa37db2c5c273",
                    "height": 411
                }
            ],
            "type": 0,
            "content": "9999999",//文字朋友圈内容
            "sid": 7380268805390937952, //朋友圈唯一id
            "update_time": 1718358258161,//更新时间
            "xid": [],//是否可见列表
            "author_vid": 1688855587446404,
            "delete_comment_list": [//删除的评论列表
                {
                    "from_id": 1688855587446404,
                    "flag": 3,
                    "commentid": 99,
                    "comment": "99999",
                    "time": 1718358136,
                    "postid": "CIOACBCDrYGxgTIYhL33qpWAgAMgGw==",
                    "ref_commentid": 33,
                    "sid": 7380268805390937952
                }
            ],
            "comment_list": [//评论列表
                {
                    "from_id": 1688855587446404,
                    "flag": 2,
                    "commentid": 3,
                    "comment": "11111111",
                    "time": 1718357924,
                    "postid": "CIOACBDBufSwgTIYhL33qpWAgAMgDg==",
                    "ref_commentid": 0,
                    "sid": 7380268805390937952
                },
                {
                    "from_id": 7881302555913738,
                    "flag": 2,
                    "commentid": 33,
                    "comment": "66666",
                    "time": 1718358066,
                    "postid": "",
                    "ref_commentid": 3,
                    "sid": 7380268805390937952
                },
                {
                    "from_id": 1688855587446404,
                    "flag": 2,
                    "commentid": 67,
                    "comment": "99999",
                    "time": 1718358082,
                    "postid": "CIOACBD/iv6wgTIYhL33qpWAgAMgGA==",
                    "ref_commentid": 33,
                    "sid": 7380268805390937952
                }
            ],
			"link_info": { //朋友圈连接消息内容
                "content_url": "",//连接地址
                "title": ""//链接标题
            },
            "visible_type": 1,
            "time": 1718352736,
            "seq": 1718352736332,
            "SnsLikeInfo": []//朋友圈点赞列表
        },
        {
            "flag": 0,
            "task_sid": 0,
            "postid": "CAQQ2YOmsAYYhL33qpWAgAMg47LoAw==",
            "media": [],
            "type": 0,
            "content": "6666666",
            "sid": 7352550640457469644,
            "update_time": 1711899098735,
            "xid": [],
            "author_vid": 1688855587446404,
            "delete_comment_list": [],
            "comment_list": [],
            "visible_type": 0,
            "time": 1711899098,
            "seq": 1711899098735,
            "SnsLikeInfo": []
        }
    ],
    "errcode": 0,
    "errmsg": "ok"
}
```


### 10.2 获取朋友圈详情 ✅ `GetSnsInfo`

##### 简要描述

- 获取朋友圈详情

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/GetSnsInfo
  
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
    "uuid":"xxxxxxxxxxxx",
    "sid":7380268805390937952 //朋友圈id
}
```

##### 返回示例 

``` 
{
    "data": {
        "flag": 0,
        "task_sid": 0,
        "postid": "CIOACBDL4reugTIYhL33qpWAgAMgDg==",
        "media": [ //图片视频信息
            {
                "size": 90092,
                "width": 600,
                "video_len": 0,
                "store_type": 0,
                "type": 1,//类型 图片 2视频 3视频缩略图
                "fileid": "*1*viLte4iMW9/VC5OZ/ZZ7FuE9VD+AhmsNxWKoL+WZI6sIZZBgOJs90NPaP6WefVRdiM+JCAfcHuTXGiawQgAn3rGdHaXkWIJpytUgbBR4zT1vxhvnwdgxHfmhPtzN1zYe+9+aYNHW6c5EAS1WxWOr2jiZvFKV+6z9JXn7O1yrY725OF13asLOu9pJfKQDfNbLDICpL4M7TUOguTXfyGTqo7dOR2H60=",
                "md5": "29a647af6dc7ec08653aa37db2c5c273",
                "height": 411
            }
        ],
        "type": 0,
        "content": "9999999",
        "sid": 7380268805390937952,//朋友圈id
        "update_time": 1718358258161,
        "xid": [ //可见人列表
            7881300143960291,
            7881302059975581,
            7881300621116017,
            7881299726922560,
            7881302555913738
        ],
        "author_vid": 1688855587446404,
        "delete_comment_list": [//删除的评论列表
            {
                "from_id": 1688855587446404,
                "flag": 3,
                "commentid": 99,
                "comment": "99999",
                "time": 1718358136,
                "postid": "CIOACBCDrYGxgTIYhL33qpWAgAMgGw==",
                "ref_commentid": 33,
                "sid": 7380268805390937952
            }
        ],
        "comment_list": [//当前评论列表
            {
                "from_id": 1688855587446404,
                "flag": 2,
                "commentid": 3,
                "comment": "11111111",
                "time": 1718357924,
                "postid": "CIOACBDBufSwgTIYhL33qpWAgAMgDg==",
                "ref_commentid": 0,
                "sid": 7380268805390937952
            },
            {
                "from_id": 7881302555913738,
                "flag": 2,
                "commentid": 33,
                "comment": "66666",
                "time": 1718358066,
                "postid": "",
                "ref_commentid": 3,
                "sid": 7380268805390937952
            },
            {
                "from_id": 1688855587446404,
                "flag": 2,
                "commentid": 67,
                "comment": "99999",
                "time": 1718358082,
                "postid": "CIOACBD/iv6wgTIYhL33qpWAgAMgGA==",
                "ref_commentid": 33,
                "sid": 7380268805390937952
            }
        ],
        "link_info": {
            "content_url": "",
            "title": ""
        },
        "visible_type": 1,
        "time": 1718352736,
        "seq": 1718352736332,
        "SnsLikeInfo": [ //点赞列表
              {
               "xid": 1688855587446404, //点赞人
               "time": 1718359990,//时间
               "postid": "",
               "sid": 7380268805390937952
           }
        ]
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 10.3 发送朋友圈 ✅ `PostPyq`

##### 简要描述

- 发送朋友圈

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/PostPyq
  
##### 请求方式
- POST 

- ContentType:"application/json"

##### 参数

| 参数名 | 必选 | 类型       | 说明                                         |
| :----- | :--- | :--------- | -------------------------------------------- |
| uuid   | 是   | String     | 每个实例的唯一标识，根据uuid操作具体企业微信 |
| vid    | 否   | 数组长整型 | 可见人数组空或者[] 则全部可见                |

- 所有的图片 视频 都走大文件上传 别走cdn上传 要走大文件上传！！！！！！！！！！！！！！！！！！！！！！！！！！！！
##### 请求示例 发送视频朋友圈

``` 
{
    "uuid":"2b0863724106a1160212bd1ccf025295",
    "content":"测试朋友圈发送",
    "file_content":
    [
        {
            "type":3, //视频图片类型 ，发送视频朋友圈这个type3是必带的 3 2 3 2这种不能只单发一个2 type1 是图片 视频只能发送一个一次请求里面只能一个3 2  多了无效
            "fileid":"*1*FdYBcjYlkXxMbjhLq4gaskUndl0tSn3NM6NuCii4NJEfibEUg0Lm7Jwk7L1JkWasl7MM/IeH5D0dbmPKXsLKO+dtUbrvw0Zjq/G6f2wXZf4LQQBRUkyH2ZnUAMzbSxEQ0QlQB58503KT0OETQadW/NxMm8=",
            "md5":"981d4320d0c4af27774c2d6de31781e9",
            "size":24389,
            "width":200,
            "height":200
        },
        {
            "type":2,//视频类型
            "fileid":"*1*FdYBcjYlkXiMuyxMWndBkjKgrHOcp641+X4L/APYAtqtu+tSub3EHKqQtivTVbf/qq/fr9Rotm9xfMID0c//EXclKcthkIwDJgSSlzmp1m9kfA8TyA14zfw+JZ+cFTobDmXwbhZGIHxS6H5zVCHz7703epfi0D2Nm1o119dfpa3gT5oC4xwfPbdDLdnwvqlDVIxSqDVCZfc+A+EQCr4YzavVPtzoKeS6ZDkFIJNN52G1Q3DNt/Q/ATaCNoOXiUyT8EIF2vHQ=",
            "md5":"7dc7b1a27f9aa62c26b0790cc9cd859b",
            "videoLen":4, //视频时常
            "size":1592424
        }
    ],
    "vids":[
        7881301488924502
    ]
}
```

##### 请求示例  发送图片朋友圈

``` 
{
    "uuid":"2b0863724106a1160212bd1ccf025295",
    "content":"测试朋友圈发送",
    "file_content":
    [
		{
            "type":1, //图片类型可以多个 但是不能和2 3 共存
            "fileid":"*1*FdYBcjYlkXxMbjhLq4gaskUndl0tSn3NM6NuCii4NJEfibEUg0Lm7Jwk7L1JkWasl7MM/IeH5D0dbmPKXsLKO+dtUbrvw0Zjq/G6f2wXZf4LQQBRUkyH2ZnUAMzbSxEQ0QlQB58503KT0OETQadW/NxMm8=",
            "md5":"981d4320d0c4af27774c2d6de31781e9",
            "size":24389,
            "width":200,
            "height":200
        }
    ],
    "vids":[
        7881301488924502
    ]
}
```
##### 请求示例  发送连接朋友圈

``` 
{
    "uuid":"2b0863724106a1160212bd1ccf025295",
    "content":"测试朋友圈发送",
    "file_content":
    [
		{
            "type":1, //连接封面图
            "fileid":"*1*FdYBcjYlkXxMbjhLq4gaskUndl0tSn3NM6NuCii4NJEfibEUg0Lm7Jwk7L1JkWasl7MM/IeH5D0dbmPKXsLKO+dtUbrvw0Zjq/G6f2wXZf4LQQBRUkyH2ZnUAMzbSxEQ0QlQB58503KT0OETQadW/NxMm8=",
            "md5":"981d4320d0c4af27774c2d6de31781e9",
            "size":24389,
            "width":200,
            "height":200
        }
    ],
    "vids":[
        7881301488924502
    ],
    "link_info":{ //链接朋友圈
        "title":"标题测试啊啊啊啊啊啊啊啊",
        "content_url":"https://www.baidu.com/?tn=02003390_10_hao_pg"
    }


}
```
##### 请求示例  发送文字朋友圈

``` 
{
    "uuid":"2b0863724106a1160212bd1ccf025295",
    "content":"测试朋友圈发送",
    "file_content":[],
    "vids":[]
}
```

##### 返回示例 

``` 
{
    "data": {
        "sid": 7553604806482115006,
        "author_vid": 1688857368772759,
        "seq": 1758710669314,
        "time": 1758710669,
        "content": "测试123123",
        "media": [
            {
                "fileid": "*1*F34hSNMREYF+BNgTfNtYo7Lg5/va+EoTPkyESrZSpwvPAoJpAras/Zrv9SF1o4pu4bOgJMG3D+2VaBOiBAtl38Ydm6+Z9kpKGVJm7e6IoGsZMUsu/SYMOsIowWjlehHbHXex+9V3yBMD+LgM9XD9j/FVhrrAOkGGa3d3hv10YohZqKtksQaJnE/EcATcv9yUKukgaOM//VtrPLdjS8uTf4NlblRUYLA/WWbgm6xhH0JpBbb2FaF33nQl0viUJFqaK2OJy5Lh3FhI4mOtecXgMFVfcJ+wL3ecfE1D9Duz0XwbPANmeAuaGugKM7VAE6RJ5Win18tN11ScE4zbrKEYwJcbXCw8qV1zk4wqlfTufEimE1cGNNlUu95FKjnfrEml702uNTaCtoY9gdDSMd/JAd0VZ8P3j8dRV1r6l1MRfdu60Yzf34LY2vz7S9THWY2cwczOelofeEM43yudTarHtJUIM7j1pM50R6ELInOogSw=",
                "md5": "626f56cc7d16f5a94017442c67975f93",
                "width": 0,
                "height": 0,
                "store_type": 0,
                "type": 2,
                "video_len": 3,
                "size": 611560
            },
            {
                "fileid": ":*1*F34hSNMREYF+BNgTfNtYo7Lg5/va+EoTPkyESrZSpwvPAoJpAras/Zrv9SF1o4pu4bOgJMG3D+2VaBOiBAtl38d4/yawYresryehdtMQYWmfeWvA3EiEE8E2bRGyR9hjiosKaQg+1Ib5Ql6X/3PrSz0ccjFMAgpyIOm2cPTYpmX8dD2+iD1ktWWRM2VSuxaLS1pJ1MIkedavlCGhtODA4QLKoSs8glwWM7F4jk+VDwdKdaLaQihkTGin285EzmS25M4haTo/Cm9Nur37WMQsV7TkrnZVmMfD9AvGmoRmXFpAs70oTypq4FU83UGFuKvExV4judYel3r9Sr4nAIONZ5QsqvV6cwplC8rmO2oz5Mqu8muPXnWYX7jEe9f7QP10DakupWUuWGAR2BJHzn/n266QL/mrrj3xt3ycrBWXoW0fsf34HGqsYMWPmKDnMjlfT9s+u83oCIJYS9SEU/fbSPTrwlR0lGzKCQLUqwFJw+8=",
                "md5": "9c333d9b82227489ad4de467d4f21307",
                "width": 0,
                "height": 0,
                "store_type": 0,
                "type": 3,
                "video_len": 0,
                "size": 24989
            }
        ],
        "comment_list": [],
        "postid": "CIOACBCzvszalzMYl4Gr/JuAgAMgEA==",
        "flag": 0,
        "update_time": 1758710669314,
        "delete_comment_list": [],
        "xid": [
            7881302555913738
        ],
        "type": 0,
        "task_sid": 0,
        "link_info": {
            "title": "",
            "content_url": ""
        },
        "poiInfo": {
            "city": "",
            "latitude": "",
            "longitude": "",
            "poiAddress": "",
            "poiClassifyId": "",
            "poiClassifyType": "",
            "poiClickableStatus": "",
            "poiName": "",
            "poiScale": ""
        },
        "visible_type": 0,
        "SnsLikeInfo": []
    },
    "errcode": 0,
    "errmsg": ""
}
```

### 10.4 删除朋友圈 ✅ `DeleteSns`

##### 简要描述

- 删除朋友圈

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/DeleteSns
  
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
    "uuid":"1688855587446404",
    "sid":7380268805390937952 //朋友圈id
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

### 10.5 发送朋友圈评论 ✅ `CommentSns`

##### 简要描述

- 发送朋友圈评论

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/CommentSnsReq
  
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
    "uuid":"1688857826547606",
    "ref_commentid":0, //评论id 默认传0  如果需要回复别人的评论则传别人评论的id 从获取朋友圈列表中获取
    "sid":7297137800987783140, //朋友圈消息对应的id
    "comment":"哈个头" //评论内容
}
```

##### 返回示例 

``` 
{
    "data": {
        "commentid": 0
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 10.6 删除朋友圈评论 ✅ `DeleteSnsComment`


##### 简要描述

- 删除朋友圈评论

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/DeleteSnsComment
  
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
    "uuid":"22222swwwwwwwwwwwwwww",
    "sid":7380268805390937952,//朋友圈id
    "commentid":99//评论id
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

### 10.7 朋友圈点赞或者取消点赞 ✅ `LikeOrUnLikeSns`


##### 简要描述

- 朋友圈点赞或者取消点赞

##### 请求URL
- ` http://47.94.7.218:8083/wxwork/LikeOrUnLikeSns
  
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
    "uuid":"1688855587446404",
    "operation":1,//0点赞  1取消点赞
    "sid":7380268805390937952 //朋友圈id
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

---

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

## 十三、客服管理 ❌ 整个模块未封装

> 文档页面未成功抓取，以下为导航栏中可见的接口名称，具体参数待确认

- 获取客服列表
- 其他客服相关接口待确认

---

## 十四、商户收款 ❌ 整个模块未封装

### 14.1 获取商户列表 ❌ 未封装

- **请求URL**: `/wxwork/MerchantInformation`
- **请求方式**: POST
- **参数**:

| 参数名 | 必选 | 类型 | 说明 |
|--------|------|------|------|
| uuid | 是 | String | 实例唯一标识 |

- **返回示例**:
```json
{
    "data": [
        {
            "mch_type": 1002,
            "mch_name": "xxx照xxx",
            "company_name": "古田xxxx照明店",
            "mch_id": "1650xxxx"
        }
    ],
    "errcode": 0,
    "errmsg": "ok"
}
```

### 14.2 获取收款二维码 ❌ 未封装

- **请求URL**: `/wxwork/GetUserMerchantQrCode`
- **请求方式**: POST
- **参数**:

| 参数名 | 必选 | 类型 | 说明 |
|--------|------|------|------|
| uuid | 是 | String | 实例唯一标识 |
| name | 是 | String | 商户名称 |
| amount | 是 | float | 金额 |
| mchid | 是 | String | 商户ID |

- **返回示例**:
```json
{
    "data": {
        "ret": 0,
        "qrcode_url": "http://127.0.0.1:8083/vdet6JKr2d00005c07ec97.png",
        "project_id": "vdet6JKr2d00005c07ec97"
    },
    "errcode": 0,
    "errmsg": "ok"
}
```

### 14.3 发送收款小程序 ❌ 未封装

- **请求URL**: `/wxwork/sendPayment`
- **请求方式**: POST
- **参数**:

| 参数名 | 必选 | 类型 | 说明 |
|--------|------|------|------|
| uuid | 是 | String | 实例唯一标识 |
| name | 是 | String | 商户名称 |
| amount | 是 | float | 金额 |
| mchid | 是 | String | 商户ID |
| buyer_id | 是 | long | 买家ID |
| words | 否 | String | 收款留言 |

---

## 附：SDK 已有但文档中未明确分类的方法

| SDK方法 | API路径 | 说明 |
|---------|---------|------|
| `SearchContactByPhone` | `/wxwork/SearchContact` | 手机号搜索联系人 |
| `AddSearch` | `/wxwork/AddSearch` | 添加搜索结果好友 |
| `AddCard` | `/wxwork/AddBusinessCard` | 添加名片好友 |
| `IsRunningByVid` | (本地辅助) | 判断某个号是否运行 |



# 企业微信协议服务回调下发消息

## 一. 登陆

### 1.1 扫码返回数据

**简要描述：** 

- 扫码返回数据

**返回类型：** 
- ` 100001 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` json
{
    "Vid":1688853790599424,
    "Qrcode_key":"D368DC88254A3E79FF72942B588D7FBE",
    "Corpid":1970325109135931,
    "Nick_name":"姓名",
    "Icon_url":"http://wework.qpic.cn/bizmail/yfG7qDz7ianwM1s5JXtgBLsaZKQfQ5YxZIiaQpUxAkEu9icIMaKxF3C4Q/0",
    "Logo":"https://p.qlogo.cn/bizmail/haQ4Dk0CFb61TIkB1JVibTUPkRynsZfK1hG9X0gDCv4cicaCObHtUCsA/0"
}

```

### 1.2 取消扫码返回

**简要描述：** 

- 取消扫码返回

**返回类型：** 
- ` 100003 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` json
{
    "Vid":1688853790599424,
    "Qrcode_key":"D368DC88254A3E79FF72942B588D7FBE",
    "Corpid":1970325109135931,
    "Nick_name":"姓名",
    "Icon_url":"http://wework.qpic.cn/bizmail/yfG7qDz7ianwM1s5JXtgBLsaZKQfQ5YxZIiaQpUxAkEu9icIMaKxF3C4Q/0",
    "Logo":"https://p.qlogo.cn/bizmail/haQ4Dk0CFb61TIkB1JVibTUPkRynsZfK1hG9X0gDCv4cicaCObHtUCsA/0"
}

```

### 1.3 确认扫码返回数据

**简要描述：** 

- 确认扫码返回数据

**返回类型：** 
- ` 100002 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` json
{
    "Vid":1688853790599424,
    "Qrcode_key":"D368DC88254A3E79FF72942B588D7FBE",
    "Corpid":1970325109135931,
    "Nick_name":"姓名",
    "Icon_url":"http://wework.qpic.cn/bizmail/yfG7qDz7ianwM1s5JXtgBLsaZKQfQ5YxZIiaQpUxAkEu9icIMaKxF3C4Q/0",
    "Logo":"https://p.qlogo.cn/bizmail/haQ4Dk0CFb61TIkB1JVibTUPkRynsZfK1hG9X0gDCv4cicaCObHtUCsA/0"
}

```

### 1.4 登陆成功返回数据

**简要描述：** 

- 登录成功返回数据

**返回类型：** 
- ` 104001 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` json
{
    "Corpid":0,
    "unionid":"ozynqskXWQK3wAWWKT_izAOfcIlY",
    "create_time":0,
    "Admin_vid":0,
    "sex":1,
    "mobile":"135xxxx",
    "Corp_full_name":"十年一刻",
    "acctid":"xxxxpeng",
    "avatar":"http://wework.qpic.cn/bizmail/yfG7qDz7ianwM1s5JXtgBLsaZKQfQ5YxZIiaQpUxAkEu9icIMaKxF3C4Q/0",
    "corp_name":"十年一刻",
    "Create_time":0,
    "english_name":"别名",
    "realname":"xxxxxx",
    "Vid":0,
    "Mail":"",
    "Ownername":"",
    "user_id":1688853790599424,
    "nickname":"姓名",
    "position":"111",
    "Corp_desc":"",
    "corp_id":1970325109135931,
    "Corp_name":"",
    "Corp_logo":""
}

```

### 1.5 企业切换回调


##### 简要描述

- 企业切换

##### 操作码

- 100006


##### 返回示例 

``` 
{
	"this_vid":1688854256622049//切换前当前登陆的账号vid
    "vid": 1688857338601153, //切换企业后的账号vid
    "corp_full_name": "",
    "corpid": 1970326674001513, //切换后的企业id
    "create_time": 1691574887,
    "mail": "",
    "ownername": "刘xx", //企业创建人昵称
    "card_url": "https://work.weixin.qq.com/wework_admin/user/h5/corp?",//企业头像
    "nickname": "奥利给干就完了", //切换后的账号昵称
    "avatar": "https://wework.qpic.cn/wwpic3az/259628_Ehj-EqAlSROfFMB_1713965307/0", //切换后的账号头像
    "corp_name": "哼哈二将"//切换后的企业名称
}
```

##### 备注 

- 更多返回错误代码请看首页的错误代码描述

### 1.6 需要输入验证码消息回调

**简要描述：** 

- 需要验证二维码消息

**返回类型：** 
- ` 100004 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "Vid": 1688857826547606, 
    "Qrcode_key": "E58DBF1516D268A519CC0C714FA0B50D", 
    "Corpid": 1970326674001513, 
    "Nick_name": "xxx", 
    "Icon_url": "https://wework.qpic.cn/wwpic/28825_3HEipRLPR4-SUsQ_1691575135/0", 
    "Logo": "https://wework.qpic.cn/wwpic/447429_AZoXgU5fQ5efSYi_1691574887/0"
}

```

## 二. 消息

### 2.1 文本消息接收

**简要描述：** 

- 文本消息接收

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 16777216,
	"kf_id":111111111, //客服id 只有客服消息接收才会有值否则就是0
    "receiver": 1688855587446404, //接收人
    "sender_name": "",
    "is_room": 0,
    "server_id": 7130717,
    "content": "32423423",
    "issync": false,
    "send_time": 1724034152,
    "sender": 7881302555913738,
    "referid": 0,
    "app_info": "3304183318011621608",
    "readuinscount": 0,
    "msg_id": 1011720,
    "msgtype": 2, //消息类型 0 和 2文本消息
    "atList": [

    ]
}

```

### 2.2 gif表情消息接收

**简要描述：** 

- gif 表情消息接收

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 16777216,
    "receiver": 1688855587446404,
    "sender_name": "",
    "is_room": 0,
    "server_id": 7130718,
    "url": "https://wework.qpic.cn/wwpic3az/wwwx_c1f24912f3d283214fd7bf5da3c3b9cd/0",
    "issync": false,
    "send_time": 1724034243,
	"width": 100,
	"height": 100,
    "sender": 7881302555913738,
    "referid": 0,
    "EmotionType": "EMOTION_DYNAMIC",
    "app_info": "7678571470421336167",
    "readuinscount": 0,
    "msg_id": 1011722,
    "msgtype": 104
}

```

### 2.3 小程序消息

**简要描述：** 

- 小程序消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
	"thumbFileId": "306b020102046430620201000204bf366e0a02031e90380204dfc6f46d0204639acc410436333230383031353337305f313531323736303031375f38326163316362353632346136313035393265643566363861353963326666650201000203009c8004000201010201000400",
	"flag": 16777216,
	"receiver": 1688853790599424,
	"sender_name": "",
	"is_room": 0,
	"server_id": 12779516,
	"title": "铁友火车票",
	"appName": "小程序名称",
	"thumbMD5": "82ac1cb5624a610592ed5f68a59c2ffe",
	"pagepath": "pages/home/index.html",
	"send_time": 1671089217,
	"size": 40058,
	"sender": 7881302555913738,
	"referid": 0,
	"appid": "wx45dff5234240ad90",
	"thumbAESKey": "7063756277636571667576656873616A",
	"app_info": "from_msgid_2000643230453161151",
	"readuinscount": 0,
	"msg_id": 1108389,
	"msgtype": 78,
	"username": "gh_c4a2a98a7366@app",
	"weappIconUrl": "http://mmbiz.qpic.cn/mmbiz_png/BMAjdIR5qwChFiaRnxUTiadJicGwMpqchXVdfQQ67tHTzfgd0gkibupgeOKtHtyE9m9SZopHXxcpsSicWxzLKXwk0qA/640?wx_fmt=png&wxfrom=200",
	"desc": ""
}

```

### 2.4 链接消息

**简要描述：** 

- 连接消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
	"authkey": "",
	"flag": 16777216,
	"receiver": 1688853790599424,
	"headImg": "https://mmbiz.qpic.cn/mmbiz_jpg/0oBxv6CfzTpWI5gwRZg3PTSmHNkCI6sfuqge5woWBCEMZKwWiaOyvBseEvq3hPNkc8h5Xe20nAnDwv0ictew7VWQ/300?wxtype=jpeg&wxfrom=0",
	"sender_name": "",
	"is_room": 0,
	"server_id": 12779517,
	"title": "不限户籍！济南新一轮300元现金补贴来了！",
	"url": "http://mp.weixin.qq.com/s?__biz=Mzk0ODE5OTg3OA==&mid=2247486862&idx=1&sn=341e039070001e40b35e3c3a221ec4b1&chksm=c36a0527f41d8c317008aa6ef7699ee501d8c139a02234c85e2bfe3345b9501f4780dd4d0852#rd",
	"send_time": 1671089268,
	"des": "美味君的福利  又双叒叕来啦！",
	"size": 0,
	"sender": 7881302555913738,
	"referid": 0,
	"app_info": "from_msgid_1491271774350930187",
	"readuinscount": 0,
	"state": "1",
	"aeskey": "",
	"msg_id": 1108392,
	"msgtype": 13
}

```

### 2.5 视频号消息

**简要描述：** 

- 视频号消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
	"thumb_url": "http://wxapp.tc.qq.com/251/20350/stodownload?encfilekey=Cvvj5Ix3eexKX1zo1IZZBrQomawdVfSQH1uu2U31EqHhx1QOrTTdw1RUwEmSiaZCAFnM77JFJoYuQpiahtJehYibzBtzKlcpKIsrkmQU4Nlg8BaiaJyqM7yK72HRXGxfY8UfKg0vaEbTk8PNrDRGP4ibwUWOK8gDwn7BXkAjAknahGY0&adaptivelytrans=0&bizid=1023&dotrans=0&hy=SH&idx=1&m=598b0abead239f31dc34b9b24a7698a1&token=x5Y29zUxcibC9gkDvObknpjEKmTiahrwGfE61uZibzqBZTaibGrWlIj3gSoVWHBlK0T0x982aQDck4I&finder_expire_time=1671694153&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAhKkzEi37bwAAAAstQy6ubaLX4KHWvLEZgBPEgKM8Q3FrDPuBzNPgMIKn8ubxw9bwD9zYyTKt3gKI",
	"cover_url": "http://wxapp.tc.qq.com/251/20350/stodownload?encfilekey=Cvvj5Ix3eexKX1zo1IZZBrQomawdVfSQH1uu2U31EqHhx1QOrTTdw1RUwEmSiaZCAFnM77JFJoYuQpiahtJehYibzBtzKlcpKIsrkmQU4Nlg8BaiaJyqM7yK72HRXGxfY8UfKg0vaEbTk8PNrDRGP4ibwUWOK8gDwn7BXkAjAknahGY0&adaptivelytrans=0&bizid=1023&dotrans=0&hy=SH&idx=1&m=598b0abead239f31dc34b9b24a7698a1&token=x5Y29zUxcibC9gkDvObknpjEKmTiahrwGfE61uZibzqBZTaibGrWlIj3gSoVWHBlK0T0x982aQDck4I&finder_expire_time=1671694153&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAhKkzEi37bwAAAAstQy6ubaLX4KHWvLEZgBPEgKM8Q3FrDPuBzNPgMIKn8ubxw9bwD9zYyTKt3gKI",
	"flag": 16777216,
	"receiver": 1688853790599424,
	"extras": "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",
	"sender_name": "",
	"is_room": 0,
	"avatar": "http://wx.qlogo.cn/finderhead/ver_1/y4ZA3geMVqNNoaVOYynb1ib0JFcAqdjZ2y4QgwpVW0Lb5mrq3c2qibBA9SUaw2rB9j9fFTx6ae4p6Wnl4hSmUn0ica9RsQe29yzzmftjNGpHzs/0",
	"server_id": 12779518,
	"url": "https://channels.weixin.qq.com/web/pages/feed?eid=export%2FUzFfAgtgekIEAQAAAAAAhKkzEi37bwAAAAstQy6ubaLX4KHWvLEZgBPEgKM8Q3FrDPuBzNPgMIKn8ubxw9bwD9zYyTKt3gKI",
	"send_time": 1671089353,
	"sender": 7881302555913738,
	"referid": 0,
	"nickname": "云剪辑.",
	"objectId":"14479811159081949639",
	"objectNonceId":"1113158575529036062_0_0_0_0_0",
	"app_info": "from_msgid_721532404981185807",
	"readuinscount": 0,
	"msg_id": 1108395,
	"msgtype": 141,
	"desc": "#余罪2 余儿康复训练"
}

```

### 2.6 名片消息

**简要描述：** 

- 名片消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
	"flag": 16777216,
	"receiver": 1688853790599424,
	"headImg": "http://wx.qlogo.cn/mmhead/X4XEGYefSBTTPqheGS9crfw2OeJW9Yb4IDfWYod02btPnsJtIXrGEQ/0",
	"sender_name": "",
	"is_room": 0,
	"server_id": 12779520,
	"send_time": 1671089416,
	"sender": 7881302555913738,
	"referid": 0,
	"nickname": "huyang",
	"app_info": "from_msgid_997448547889589003",
	"readuinscount": 0,
	"msg_id": 1108401,
	"msgtype": 41,
	"enterpriseName": "微信",
	"corp_id": 0,
	"username": 7881299726922560
}

```

### 2.7 位置消息

**简要描述：** 

- 位置消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
	"detailed_address": "xxx广场1号楼(xxxxxx路东50米)",
	"flag": 16777216,
	"address": "xxxxxxxxxx东路",
	"receiver": 1688853790599424,
	"latitude": 363.664135,
	"sender_name": "",
	"is_room": 0,
	"server_id": 12779521,
	"send_time": 1671089496,
	"sender": 7881302555913738,
	"referid": 0,
	"app_info": "from_msgid_4430967334562546003",
	"readuinscount": 0,
	"msg_id": 1108405,
	"msgtype": 6,
	"longitude": 11.148247
}

```

### 2.8 外部联系人图片消息（个v）

**简要描述：** 

- 外部联系人图片消息（个微发的）

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
	"flag": 16777216,
	"receiver": 1688853790599424,
	"sender_name": "",
	"is_room": 0,
	"server_id": 12779523,
	"file_size": 7982507,
	"send_time": 1671089597,
	"sender": 7881302555913738,
	"referid": 0,
	"file_id": "https://imunion.weixin.qq.com/cgi-bin/mmae-bin/tpdownloadmedia?param=v1_9402dda5b09845e753c80d85b8fdbbd9fec7a526c9da0b0d88ac1b7a74a4c4c1dd9b982527532a112af12c7a8a624b8117d171397ec6d0b94ca02c80bc5633357326cebc5d348c35625b2e6980fe6c9f3ecea79c3f22b5c4ab0c2d6c59d86e403f95bfcb1ffe7311b67cc0d7b31d4dd349ce38628e7c95561b036c6b787d6a04460b7146c54651e4ce73c5c31d5519936614efc9fb510d13f03229fac07d847626ed6dcf4673cd7312763fd73f763f208b2fa9e388e90fba6fc8dbb777bba4b94143d39f09a4e97a807a10195c4e871f8065321f8e9355304a669ac1deafc333399b2d06d072390ac980fcc80c4fa764ae59dd01b16dee0384df3073cdf28744",
	"openim_cdn_authkey": "v1_9402dda5b09845e753c80d85b8fdbbd9fec7a526c9da0b0d88ac1b7a74a4c4c1e5230c9b1626f74f55ffe8c3660d9d7aa62fe7376a0dddbb8f6575d040e9f369",
	"width": 0,
	"app_info": "from_msgid_9211250126199164434",
	"readuinscount": 0,
	"aes_key": "a2078ba39c47bfcfdbd9ec2574c2e134",
	"msg_id": 1108410,
	"msgtype": 101,
	"md5": "714d6b8ee6ca20fa6143338ba9c1c69a",
	"height": 0
}

```

### 2.9 外部联系人文件消息

**简要描述：** 

- 外部联系人文件消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
	"flag": 16777216,
	"receiver": 1688853790599424,
	"file_name": "dump.rdb",
	"sender_name": "",
	"is_room": 0,
	"server_id": 12779529,
	"file_size": 113564,
	"send_time": 1671089980,
	"sender": 7881302555913738,
	"referid": 0,
	"file_id": "https://imunion.weixin.qq.com/cgi-bin/mmae-bin/tpdownloadmedia?param=v1_9402dda5b09845e753c80d85b8fdbbd9fec7a526c9da0b0d88ac1b7a74a4c4c1e8b38d076dc0683fbb01af67a9ebcbeab65ecdf9aa68764d4683bec9cd8e9ffe7453d9de6b915504638ab0dda6c0c10c2ba4b38a1c7366faf736656344e31c47532c5a174bc5b4c249318253a995fa8462b2868f4b63b63fa16eef6120f377252c12817e0a30db83157402b395ff793c679d7d5f1dc3dff24f147d9f3381a3170b4ec59a28235105d1eeb7f57173afffb8aecf0e69c3620e291a6a7e78e6e188709fb6550891770dbc899598f19bd2e688204a46b30e812443fb5deafbab5bf30593f3a0b01dfb920ba86be52370103d29813dbcd0d1aa170efbbc237acbe40f",
	"openim_cdn_authkey": "v1_9402dda5b09845e753c80d85b8fdbbd9fec7a526c9da0b0d88ac1b7a74a4c4c1bf9e0299514bfddca8e6db1d3f11977073255bc80fb5ec7a0512fff9548d9783",
	"app_info": "from_msgid_3239399033279654384",
	"readuinscount": 0,
	"aeskey": "b28c380313f3e8827a2740da0a1decff",
	"msg_id": 1108425,
	"msgtype": 102,
	"md5": "f506a445ee47e2232a087d628043935c"
}
```

### 2.10 外部联系人视频消息

**简要描述：** 

- 外部视频消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "video_duration": 4, 
    "video_height": 0, 
    "openim_cdn_ld_aeskey": "fae8cf6f653b9c92adc6faec93a6e57a", //缩略图aes
    "flag": 16777216, 
    "receiver": 1688853790599424, 
    "openim_cdn_ld_size": 4997, //缩略图大小
    "video_width": 0, 
    "sender_name": "", 
    "is_room": 0, 
    "server_id": 12884979, 
    "file_size": 311641, 
	"openim_cdn_ld_size":2546,
    "send_time": 1680526456, 
    "sender": 7881302555913738, 
    "referid": 0, 
	//fileid是视频下载地址
    "file_id": "https://imunion.weixin.qq.com/cgi-bin/mmae-bin/tpdownloadmedia?param=v1_9402dda5b09845e753c80d8x06f87792ff1795fe3b00fdafcb1f52df", 
    "openim_cdn_authkey": "v1_9402dda5b09845e75xc5785886e719d334044e189991e6b36c850334edf2427756af9b",  //视频和缩略图authkey
    "aes_key": "fae8cf6f653b9c92adc6faec93a6e57a", 
    "app_info": "from_msgid_64299209193941486", 
    "readuinscount": 0, 
    "msg_id": 1119632, 
    "msgtype": 103, 
    "md5": "18826ff24ce328cd4d850edfa3f8c60c", 
    "preview_img_url": "https://imunion.weixin.qq.com/cgi-bin/mmae-bin/tpdownloadmedia?param=v1_9402dda5b09845e753c80dx2809bfcf49fa67e02354a4741a5b1e" //缩略图地址
}

```

### 2.11 内部联系人图片消息（企业微信）

**简要描述：** 

- 内部联系人图片消息（企业微信用户）

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "innerkf_vid": 0,
    "flag": 83886080,
    "extraContent": {

    },
    "mid_img_size": 25297,
    "openim_cdn_ld_height": 290,
    "devInfo": 65537,
    "sender_name": "",
    "is_room": 1,
    "is_hd": 1,
    "openim_cdn_ld_md5": "58e3b6f7ec1aa792f1b273dd821e91d9",
    "app_info": "CIGABBCK263DBhjhs6ywkICAAyAP",
    "room_conversation_id": "10817857838351957",
    "msg_id": 1026382,
    "msgtype": 14,
    "height": 154,
    "receiver": 0,
    "thumb_img_size": 6849,
    "file_name": "1722476983013.png",
    "server_id": 7752688,
    "kf_id": 0,
    "file_size": 41505,
    "issync": false,
    "send_time": 1751870858,
    "sender": 1688854256622049,
    "openim_cdn_ld_width": 290,
    "referid": 0,
    "file_id": "306b020102046430620201000204060b19e102030f55ca020498ec1d3c0204686b6d8a042430316539333437612d376664372d346238332d386362352d3530306564363038343432380203103800020300a2300410efaf887ca055806f139409b13f35b93c0201010201000400",
    "aes_key": "5349db00f66c5000d85aa9ff3a9898ca",
    "width": 154,
    "readuinscount": 0,
    "md5": "efaf887ca055806f139409b13f35b93c"
}

```

### 2.12 内部联系人文件消息

**简要描述：** 

- 文件消息内部

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
	"flag": 16777216,
	"receiver": 7881302555913738,
	"file_name": "dump.rdb",
	"sender_name": "",
	"is_room": 0,
	"server_id": 12779532,
	"file_size": 113564,
	"send_time": 1671094399,
	"sender": 1688853790599424,
	"referid": 0,
	"file_id": "3069020102046230600201000204ea44290002030f55cb0204789f216f0204639ae07e042437336331303039352d343030342d343732342d393062622d303432326361653230363763020100020301bba00410f506a445ee47e2232a087d628043935c0201050201000400",
	"aes_key": "7e22920298d94b038fac516bcf97d3ce",
	"app_info": "CAQQ/cDrnAYYgNKQ0o6AgAMg0+HxqwE=",
	"readuinscount": 0,
	"msg_id": 1108433,
	"msgtype": 15,
	"md5": "f506a445ee47e2232a087d628043935c"
}

```

### 2.12 内部联系人视频消息

**简要描述：** 

- 内部视频消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "video_duration": 23, 
    "video_height": 0, 
    "openim_cdn_ld_aeskey": "", 
    "flag": 16777216, 
    "receiver": 7881302555913738, 
    "openim_cdn_ld_size": 24389, 
    "video_width": 0, 
    "sender_name": "", 
    "is_room": 0, 
    "server_id": 15319792, 
    "file_size": 1592432, 
    "send_time": 1689910819, 
    "sender": 1688855749266556, 
    "referid": 0, 
    "file_id": "30819302010204818b30818802010002045f030c7c02030f55cb0204179f216f020464b9fd4f044c4e45574944315f35663033306337633137396632313666363462396665323330303034336133325f39356636613365342d346332352d346361652d626361332d3861656630383461333539630201000203184c7004107dc7b1a27f9aa62c26b0790cc9cd859b0201040201000400", 
    "openim_cdn_authkey": "", 
    "aes_key": "64393035623864326166303932653431", 
    "app_info": "CAQQovznpQYY/JiM+JWAgAMgutTfsQ4=", 
    "readuinscount": 0, 
    "msg_id": 1018391, 
    "msgtype": 23, 
    "md5": "7dc7b1a27f9aa62c26b0790cc9cd859b", 
    "preview_img_url": ""
}

```

### 2.13 红包消息

**简要描述：** 

- 红包消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
	"icon_url": "http://dldir1.qq.com/qqcontacts/hongbao1x_20160413.png",
	"flag": 553648128,
	"receiver": 1688853790599424,
	"sender_name": "",
	"remark": "恭喜发财，大吉大利",
	"is_room": 0,
	"server_id": 12779524,
	"send_time": 1671089820,
	"sender": 7881302555913738,
	"packet_id": "1800008896202212157398737500018",
	"referid": 0,
	"app_info": "hongbao_7881302555913738_1800008896202212157398737500018",
	"readuinscount": 0,
	"msg_id": 1108413,
	"msgtype": 26,
	"desc": "来自十年一刻的红包，请进入手机版企业微信领取"
}

```

### 2.14 语音消息

**简要描述：** 

- 语音消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
	"flag": 16777216,
	"receiver": 1688853790599424,
	"voice_id": "308180020102047930770201000204bf366e0a02031e90380204a4df892b0204639acf92044c333230383031353337305f36323734363337303637373336383735373736453634363536323645363736465f6438333562623966636330656336366337363162666331363564643961666236020100020219a004000201050201000400",
	"sender_name": "",
	"is_room": 0,
	"server_id": 12779530,
	"voice_time": 4,
	"send_time": 1671090066,
	"sender": 7881302555913738,
	"referid": 0,
	"aes_key": "6274637067736875776E6465626E676F",
	"app_info": "from_msgid_6602574656373321919",
	"voice_size": 6544,
	"readuinscount": 0,
	"msg_id": 1108428,
	"msgtype": 16,
	"md5": "d835bb9fcc0ec66c761bfc165dd9afb6"
}

```

### 2.15 语音通话或者视频通话消息

**简要描述：** 

- 语音或者视频通话消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 536870912, 
    "receiver": 0, 
    "inviteId": 7881302555913738, 
    "msgid": 7227528031095598000, 
    "sender_name": "", 
    "call_user": [ //发起人
        {
            "xid": 7881302555913738, 
            "openid": "o9cq804RcItRx7Uxv4njTZf2-mnM@im.wechat", 
            "name": "十年一刻", 
            "headUrl": "http://wx.qlogo.cn/mmhead/Q3auHgzwzM7BQVtj83EnfQcDMjz2hhrIv6vzOPdXX4mLpIEJUbE2ew/0"
        }
    ], 
    "inviteType": 0, 
    "is_room": 0, 
    "server_id": 8079815, 
    "type": 1, 
    "roomid": 1235025531, 
    "send_time": 1682790003, 
    "sender": 10024, 
    "referid": 0, 
    "app_info": "7227528031095598083", 
    "actType": 1, 
    "readuinscount": 0, 
    "roomkey": 0, 
    "msg_id": 1007924, 
    "msgtype": 503, 
    "call_type": 2,  // 2语音消息 1视频消息
    "timestamp": 1682790003, 
    "content_masg": "十年一刻 邀请你进行语音通话"
}

```

### 2.16 语音或者视频通话，未接听消息

**简要描述：** 

- 语音或者视频通话，未接听消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 536870912, 
    "receiver": 0, 
    "msgid": 7227528061160334000, 
    "sender_name": "", 
    "inviteType": 2,  //2语音 1视频
    "is_room": 0, 
    "server_id": 8079818, 
    "type": 3, //3是未接听  1是正在打电话 
    "roomid": 1235025531, 
    "senderid": 7881302555913738, 
    "send_time": 1682790010, 
    "sender": 10024, 
    "referid": 0, 
    "app_info": "wxvoipmiss_1688858100400545_12350255311682790010", 
    "readuinscount": 0, 
    "msg_id": 1007930, 
    "wording": "来自十年一刻6666649765的未接语音通话", 
    "msgtype": 503, //正在通话 挂断 未接都是503 根据里面的字段状态区分具体是哪一个
    "timestamp": 1682790010
}

```

### 2.17 语音或者视频通话，在其他端接听消息

**简要描述：** 

- 语音或者视频通话，在其他端接听消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 536870912, 
    "receiver": 0, 
    "msgid": 7227530934493491000, //当前语音消息的id 根据这个来匹配这个语音的生命周期每个语音消息都有这个id 通话 接听 或者挂断 接听时常
    "sender_name": "", 
    "is_room": 0, 
    "server_id": 8079825, 
    "type": 2, // 1是正在打电话 2是其他端接收 3是挂断
    "send_time": 1682790679, 
    "sender": 10024, 
    "referid": 0, 
    "app_info": "7227530934493491289", 
    "readuinscount": 0, 
    "msg_id": 1007946, 
    "msgtype": 503, 
    "timestamp": 1682790679, 
    "content_masg": "已在其他终端接听/拒接"
}

```

### 2.18 语音或者视频通话消息，对方取消通话

**简要描述：** 

- 语音或者视频通话消息，对方取消通话

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 553648128, 
    "receiver": 1688858100400545, 
    "sender_name": "", 
    "is_room": 0, 
    "recordwording": "对方已取消", 
    "server_id": 8079817, 
    "send_time": 1682790010, 
    "sender": 7881302555913738, //语音发起人
    "acceptOrRejectPlatform": 131074, 
    "referid": 0, 
    "recordtype": 1, //1代表对方取消 5代表语音结束
    "app_info": "wxvoipconv_1688858100400545_12350255311682790010", 
    "readuinscount": 0, 
    "invitetype": 2, //2语音  1视频
    "msg_id": 1007928, 
    "msgtype": 40, 
    "acceptOrRejectImei": "A9C97D19-77AA-4CCD-A88B-B789F5D8CACE"
}




```

### 2.19 语音视频通话时长消息通知

**简要描述：** 

- 语音视频通话时常消息通知

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 553648128, 
    "receiver": 1688858100400545, 
    "sender_name": "", 
    "is_room": 0, 
    "recordwording": "通话时长00:03", 
    "server_id": 8079827, 
    "send_time": 1682790685, 
    "sender": 7881302555913738, 
    "acceptOrRejectPlatform": 131074, 
    "referid": 0, 
    "recordtype": 5, 
    "app_info": "wxvoipconv_1688858100400545_12349965841682790685", 
    "readuinscount": 0, 
    "invitetype": 2, 
    "msg_id": 1007950, 
    "msgtype": 40, 
    "acceptOrRejectImei": "A9C97D19-77AA-4CCD-A88B-B789F5D8CACE"
}

```

### 2.20 撤回消息

**简要描述：** 

- 撤回消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "revoke_time": 1720166719,//撤回的消息的发送时间
    "flag": 0,
    "receiver": 1688855587446404, //接收人
    "sender_name": "",
    "is_room": 0,
    "revoke_msgtype": 2, //撤回消息的消息类型
    "server_id": 7122663,
    "revoke_ref_appinfo": "2068197696680161941", //撤回消息的appinfo
    "issync": false,
    "send_time": 1720166720,//撤回时间
    "sender": 7881302555913738,//发送人 也是撤回人
    "referid": 0,
    "app_info": "UbOr114nTxOjSjb",
    "readuinscount": 0,
    "msg_id": 1009088,
    "msgtype": 2063
}

```

### 2.21 合并消息

**简要描述：** 

- 合并消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 16777216, 
    "receiver": 1688853790599424, 
    "sender_name": "", 
    "is_room": 0, 
    "server_id": 12885464, 
    "send_time": 1681057244, 
    "sender": 7881302555913738, 
    "referid": 0, 
    "app_info": "from_msgid_4323568062863172402", 
    "readuinscount": 0, 
    "msg_id": 1120902, 
    "msgtype": 4, 
    "msglist": [
        {
            "send_time": 1680871380, 
            "sender": 0, 
            "send_name": "sfdsf", 
            "msgtype": 2, 
            "atList": [ ], 
            "content": "456"
        }, 
        {
            "send_time": 1680871380, 
            "sender": 0, 
            "send_name": "sfdsf", 
            "msgtype": 2, 
            "atList": [ ], 
            "content": "45"
        }, 
        {
            "send_time": 1680871380, 
            "sender": 0, 
            "send_name": "sfdsf", 
            "msgtype": 2, 
            "atList": [ ], 
            "content": "54"
        }, 
        {
            "video_duration": 4, 
            "video_height": 0, 
            "openim_cdn_ld_aeskey": "588c27ec325bca3f906e8b545f281dc8", 
            "openim_cdn_ld_size": 19207, 
            "video_width": 0, 
            "file_size": 753510, 
            "send_time": 1680981000, 
            "sender": 0, 
            "send_name": "sfdsf", 
            "file_id": "https://imunion.weixin.qq.com/cgi-bin/mmae-bin/tpdownloadmedia?param=v1_9402dda5b09845e753c80d85b8fdbbd9fec7a526c9da0b0d88ac1b7a74a4c4c19759fb25e6358f94c27c21c56da8848244d5d7f92a277a2c4b3a3c05719b5c1c789ddeb5bca81b08b7e158ba9ca786499a4bd003f4e8f68e8d59cacbe6a2fdabd8f9338c0fb8f2d4784e4b9ec7f85b6ab984a2ecdf7c678cb1d71065d5e4cf03112bd9b9cf1a382e12f6f067e335dd4f3001780aacf1cac1102515952eddc6a5869b9cb91b2b34c810c700e1783e9e2f69dc35b72af085998e5ba760c848d7bf2f784d5a0d0", 
            "openim_cdn_authkey": "v1_9402dda5b09845e753c80d85b8fdbbd9fec7a526c9da0b0d88ac1b7a74a4c4c1b3af2ae4a2d511d840070be7565291c2a9447818d75fbffd07e052b4ef0a7d73", 
            "aes_key": "2a2e1f60c749c6ad925e1e59aa7945aa", 
            "msgtype": 103, 
            "md5": "3b9399f4489989e2bb886fe35fce4df4", 
            "preview_img_url": "https://imunion.weixin.qq.com/cgi-bin/mmae-bin/tpdownloadmedia?param=v1_9402dda5b09845e753c80d85b8fdbbd9fec7a526c9da0b0d88ac1b7a74a4c4c17f12c4a584c08854f3380188dcaf6be657e9411151907717abe99cdf99facee539dd2aae29c125946ae0574a9d7db878cd560c4f61ec9290aa389a3c17584bc78379d4eaf85445291826662e6e13a3960fa73b5fe3a4f28fd6047651227244a6687b6a2deb205193710fc58ba8e81"
        }, 
        {
            "authkey": "v1_9402dda5b09845e753c80d85b8fdbbd9fec7a526c9da0b0d88ac1b7a74a4c4c1b3af2ae4a2d511d840070be7565291c2a9447818d75fbffd07e052b4ef0a7d73", 
            "send_time": 1681038120, 
            "des": "百度内容", 
            "size": 15444, 
            "headImg": "https://imunion.weixin.qq.com/cgi-bin/mmae-bin/tpdownloadmedia?param=v1_9402dda5ddfdaa9592b7db664ebbecc90e639e9374ccd1fcc5e8aeda95eff1c20d07577787f5b50923c9dace523dc5e72e25a2e6ad6c112d9a1bfed1e9bedc3232344a69f8cbeb1a32c47d20c17c0b7d55e0f5766ee90", 
            "sender": 0, 
            "send_name": "十年一刻", 
            "state": "1", 
            "aeskey": "80f537cf7f8e63dca87e08a5a02d444d", 
            "title": "百度", 
            "msgtype": 13, 
            "url": "https://www.baidu.com/"
        }
    ]
}

```

### 2.22系统消息(发送失败封禁功能等提醒)

**简要描述：** 

- 系统消息返回

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "msg": "因账号违规，当前功能已被限制使用。 查看详情",
    "flag": 16777216,
    "receiver": 1688856286xxxxx,
    "extraContent": {
        "extraMessageLink": {
            "jumpUrl": "https://work.weixin.qq.com/webapp/spam/info?key=%2FBe%2BZ9UmUqQFKjmhn6LLrFnZtUWQ0ZPud1Or1NgA0MySuAHkVQX7b6Cn64b2jnI18CbSyYv8Hu%2FwKgKiyJiRLDvDgvlck8sNOPhF86cUWi3RtRyVlesQ2hhy0CBLNaTqxYB%2B2txNf9Z3oOA3igKPFl%2FabvJGCem6lzlLJ5JrqMu7TiRBL17J6CWrOpuuNLml5d1C9vDUZRNtskQuh3IEyg%3D%3D&type=3&need_login=true", //详情连接，根据此连接去调用封禁查询接口。
            "linkType": 1
        }
    },
    "sender_name": "",
    "is_room": 0,
    "server_id": 15380771,
    "issync": false,
    "send_time": 1745750312,
    "sender": 788130229xxxx,
    "referid": 0,
    "app_info": "spam_block_CIOACBClj8+25zIYyuir+JeAgAMgDw==",
    "readuinscount": 0,
    "msg_id": 1000796,
    "msgtype": 1011
}

```

### 2.23 接龙消息接收

**简要描述：** 

- 接龙消息接收

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "innerkf_vid": 0,
    "flag": 83886080,
    "receiver": 0,
    "extraContent": {

    },
    "content": {
        "solitaireInfo": {
            "author_id": 7881302555913738,//接龙发起人id
            "default_punctuation": ".",
            "example_range": "9-6", //接龙示例内容，从title下标9开始数6个数 代表接龙示例：《 格式示例\n》
            "has_header": true,
            "items": [
                {
                    "content": "",
                    "header": false,
                    "punctuation": ".",
                    "range": "20-11", //接龙发送内容 20-11 表示从下标20开始数11个数 代表这个人的接龙内容 ：十年一刻 666666
                    "timestamp": 1750605934, //接龙发送时间
                    "user_id": 7881302555913738  //接龙发送人
                },
                {
                    "content": "",
                    "header": false,
                    "punctuation": ".",
                    "range": "35-17",//接龙发送内容 35-17 表示从title下标35开始数17个数 代表这个人的接龙内容 ：77777 23423423423
                    "timestamp": 1750606059, //接龙发送时间
                    "user_id": 1688854256622049 //接龙发送人
                }
            ],
            "loss": false,
            "tail_range": "",
            "title_range": "0-8" //标题 0-8 表示从0开始数8个数 是标题内容
        },
        "title": "#接龙\n测试接龙\n例 格式示例\n\n1. 十年一刻 666666\n2. 77777 23423423423" //接龙内容
    },
    "devInfo": 65537,
    "sender_name": "77777",
    "is_room": 1,
    "server_id": 7752377,
    "kf_id": 0,
    "issync": false,
    "send_time": 1750606064,
    "sender": 1688854256622049,
    "referid": 0,
    "app_info": "CIGABBDwweDCBhjhs6ywkICAAyAD",
    "room_conversation_id": "10831012479402677",
    "readuinscount": 0,
    "msg_id": 1025642,
    "msgtype": 213
}

```

### 2.24 视频号直播消息

**简要描述：** 

- 视频号直播消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "thumb_url": "https://finder.video.qq.com/261/20304/stodownload?encfilekey=oibeqyX228riaCwo9STVsGLIia39tYEhYEEKGI4Gox5zM20A8jGDaUDxBy3SUwgTq5Da9HWpFDSIgRicZfqXq5j79iaWAtMsX847x56t6Lfqe0l1Dnf89WuOibY6tUB9WVIkRlbUCaqes4TI8&token=x5Y29zUxcibDMyGL6Na7xYEOQicWBULHl5tre0jffh0lQ1AtIL9Ore7V26I833jia7hy8pMrxY6cRU0bicyiaWQxE4icjusOOcHCnib4Lfy3TqNsgvQpjSnkHIMibR7iapWA53XXF8j1nGqVwWCscRw4ibJ2wUv36gN5wHVdvoJ9ibuS4G7NV2Z1NMF2BQp1UjNWx0ornDj&bizid=1023&hy=SZ&idx=1&m=8fee888ecafc02e8c31c3dbdaa5dce9a&finder_expire_time=1767685885&wxampicformat=503&picformat=200&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAwO0AoQR3AgAAAAstQy6ubaLX4KHWvLEZgBPEyoNceT8YK7WOzNPgMIIz8OIpYmWXb2aByCDBEEbP",
    "innerkf_vid": 0,
    "flag": 16777216,
    "extraContent": {

    },
    "devInfo": 0,
    "extras": "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",
    "sender_name": "",
    "is_room": 0,
    "nickname": "三角洲行动丶海浪呀",
    "app_info": "3826338618172894593",
    "msg_id": 1008137,
    "msgtype": 146,
    "objectId": "14817343751800428574",
    "cover_url": "https://finder.video.qq.com/261/20304/stodownload?encfilekey=oibeqyX228riaCwo9STVsGLIia39tYEhYEEKGI4Gox5zM20A8jGDaUDxBy3SUwgTq5Da9HWpFDSIgRicZfqXq5j79iaWAtMsX847x56t6Lfqe0l1Dnf89WuOibY6tUB9WVIkRlbUCaqes4TI8&token=x5Y29zUxcibDMyGL6Na7xYEOQicWBULHl5tre0jffh0lQ1AtIL9Ore7V26I833jia7hy8pMrxY6cRU0bicyiaWQxE4icjusOOcHCnib4Lfy3TqNsgvQpjSnkHIMibR7iapWA53XXF8j1nGqVwWCscRw4ibJ2wUv36gN5wHVdvoJ9ibuS4G7NV2Z1NMF2BQp1UjNWx0ornDj&bizid=1023&hy=SZ&idx=1&m=8fee888ecafc02e8c31c3dbdaa5dce9a&finder_expire_time=1767685885&wxampicformat=503&picformat=200&finder_eid=export%2FUzFfAgtgekIEAQAAAAAAwO0AoQR3AgAAAAstQy6ubaLX4KHWvLEZgBPEyoNceT8YK7WOzNPgMIIz8OIpYmWXb2aByCDBEEbP",
    "objectNonceId": "10043444170675585450_0_0_0_0_0_03fd0748-df0b-11f0-86dc-212ae746b05c",
    "receiver": 1688857368772759,
    "avatar": "http://wx.qlogo.cn/finderhead/c1Ohodl25RZ9EV2LicMM88ceSOqeNfFm36BUgjsFQjHcHsEg8zicYnMHcaxeXZ25wAficQTUhQQHUU/0",
    "server_id": 13274178,
    "kf_id": 0,
    "url": "https://channels.weixin.qq.com/web/pages/live?eid=export%2FUzFfAgtgekIEAQAAAAAAwO0AoQR3AgAAAAstQy6ubaLX4KHWvLEZgBPEyoNceT8YK7WOzNPgMIIz8OIpYmWXb2aByCDBEEbP",
    "issync": false,
    "send_time": 1766389884,
    "sender": 7881302555913738,
    "referid": 0,
    "readuinscount": 0,
    "desc": "绝密航天典狱长"
}

```

### 2.25 内部联系人大文件消息

**简要描述：** 

- 内部联系人大文件消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "innerkf_vid": 0,
    "flag": 16777216,
    "receiver": 1688857368772759,
    "extraContent": {

    },
    "file_name": "postman_v7.6.0_downyi.com.zip",
    "devInfo": 65537,
    "sender_name": "xxx",
    "is_room": 0,
    "server_id": 13274472,
    "kf_id": 0,
    "file_size": 72127884,
    "issync": false,
    "send_time": 1766651990,
    "sender": 1688856274767936,
    "referid": 0,
    "file_id": "*1*u3W2k/2+7TigF3ieM2JYWY12fJ0m+w80zO84cIONH6ap/yvmyj0JtKad3fMPQ3X7GvDGg2jv4q6CzqKlbkJVHcg6Q0BzWKuCZzNdwsCXYXNDs54NPJ579w7JqJsaRQ0RN8RSOpRUvS0jYWuRYCWToWn5rXx1OAJnNRHrwZsPOWWRkXBENPjqwRfLI8dbJ09aP2H31RfIliESX9PQuwmA2UecUBxCEMq4C7SzgMBhKFojPM4qniFtRbI/RfClC5cN2/eLtyxNaSoPZY/vLVKYeDKOXrOg6jE7X7++kHSJFIxVTo6uxmEqM+1ePWR4vSNRJy1GAKFl0r42hLoOMqN3MRNJu7yY8oLDNw9O5KmCP4y0wg5D0sXhNtaNaiVDlNLHMzcmnF7ScaumVLydgn6owW8driac/MXyqPiI3u6lJpQCGu3fW2L++DcOYSdCeESqPnyHk4OKObizEMgY5A/EKcvUeCQ7hyIz3Y4/0ZEWuKY=",
    "aes_key": "",
    "app_info": "CAEQ1vCzygYYwKDW8peAgAMgBw==",
    "readuinscount": 0,
    "msg_id": 1008782,
    "msgtype": 20,
    "md5": "02FD6642C4333E1E8104EB3F0A61212F"
}

```

### 2.26 新增客服聊天排队通知

**简要描述：** 

- 新增客服聊天排队通知

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "innerkf_vid": 0,
    "flag": 0,
    "receiver": 0,
    "extraContent": {

    },
    "devInfo": 0,
    "sender_name": "",
    "is_room": 0,
    "server_id": 7752275,
    "kf_id": 25664265286998413, //这个客服id收到的排队通知
    "issync": false,
    "send_time": 1749800760,
    "sender": 10151,
    "referid": 0,
    "app_info": "Dhodmav-QbmseNc",
    "readuinscount": 0,
    "msg_id": 1025395,
    "msgtype": 2314 //排队通知类型 收到这个类型查一下排队列表
}

```



## 三. 群

### 3.1 群新增通知

**简要描述：** 

- 群新增通知

**返回类型：** 
- ` 115001 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "room_id": "10927640775842980", 
    "create_user_id": 7881302555913738, 
    "total": 4, 
    "notice_sendervid": 10927640775842980, 
    "flag": 268435585, 
    "create_time": 1687687448, 
    "notice_content": "", 
    "nickname": "", 
    "member_list": [
        {
            "unionid": "", 
            "create_time": 0, 
            "sex": 1, 
            "mobile": "", 
            "acctid": "", 
            "join_scene": 1, 
            "avatar": "https://wework.qpic.cn/wwpic/510408_Qat5EH4YQaidquh_1678715517/0", 
            "english_name": "", 
            "realname": "xxx", 
            "room_notes": "", 
            "jointime": 1687687448, 
            "nickname": "xxxx", 
            "room_nickname": "", 
            "position": "", 
            "uin": 168xxx199881, 
            "invite_user_id": 7881302555913738, 
            "corp_id": 1970324968093920
        }, 
        {
            "unionid": "ozynqskXWQK3wAWWKT_izAOfcIlY", 
            "create_time": 0, 
            "sex": 1, 
            "mobile": "135xxxxx", 
            "acctid": "ShiNianYiKe", 
            "join_scene": 1, 
            "avatar": "https://wework.qpic.cn/wwpic/247382__uUOzm-wQSmgT5w_1687280998/0", 
            "english_name": "", 
            "realname": "xxx", 
            "room_notes": "", 
            "jointime": 1687687448, 
            "nickname": "xx", 
            "room_nickname": "", 
            "position": "", 
            "uin": 1688xxx0400545, 
            "invite_user_id": 7881302555913738, 
            "corp_id": 1970325085040417
        }, 
        {
            "unionid": "", 
            "create_time": 0, 
            "sex": 1, 
            "mobile": "", 
            "acctid": "", 
            "join_scene": 1, 
            "avatar": "https://wework.qpic.cn/wwpic/80928_SFNfb-ufTySFa6W_1652593247/0", 
            "english_name": "", 
            "realname": "xxx", 
            "room_notes": "", 
            "jointime": 1687687448, 
            "nickname": "十年一刻", 
            "room_nickname": "", 
            "position": "", 
            "uin": 16888581xxx422, 
            "invite_user_id": 7881302555913738, 
            "corp_id": 1970325105140387
        }, 
        {
            "unionid": "", 
            "create_time": 0, 
            "sex": 1, 
            "mobile": "", 
            "acctid": "", 
            "join_scene": 1, 
            "avatar": "http://wx.qlogo.cn/mmhead/Q3auHgzwzM7BQVtj83EnfQcDMjz2hhrIvssZl9fv1H6rYUaN8VRibjg/0", 
            "english_name": "", 
            "realname": "", 
            "room_notes": "", 
            "jointime": 1687687448, 
            "nickname": "十年一刻", 
            "room_nickname": "", 
            "position": "", 
            "uin": 788130xx13738, 
            "invite_user_id": 7881302555913738, 
            "corp_id": 1970325134026788
        }
    ], 
    "new_flag": 4, 
    "notice_time": 10927640775842980, 
    "managers": [ ]
}

```

### 3.2 移除群成员通知

**简要描述：** 

- 移除群成员通知

**返回类型：** 
- ` 115002 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 

{
    "flag": 0, 
    "receiver": 0, 
    "sender_name": "", 
    "is_room": 1, 
    "server_id": 15318083, 
    "send_time": 1687688952, 
    "sender": 1688855749266556, 
    "referid": 0, 
    "member_list": [
        1688856554448765, 
        1688854853482590, 
        1688855749266675
    ], 
    "app_info": "CAQQ+K3gpAYYACD3+fXFAQ==", 
    "room_conversation_id": "266246527627906", 
    "readuinscount": 0, 
    "msg_id": 1014584, 
    "msgtype": 1003
}

```

### 3.3 邀请进群通知

**简要描述：** 

- 邀请进群通知

**返回类型：** 
- ` 115003 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 

{
    "flag": 0, 
    "receiver": 1688858100400545, 
    "sender_name": "", 
    "is_room": 1, 
    "server_id": 8082424, 
    "invitee": 7881302555913738, 
    "send_time": 1687687454, 
    "sender": 7881302555913738, 
    "referid": 0, 
    "member_list": [ ], 
    "app_info": "hAfWo6o-TieYSeYynewroom", 
    "room_conversation_id": "10927640775842980", 
    "readuinscount": 0, 
    "msg_id": 1013451, 
    "msgtype": 1002
}

```

### 3.4 自己主动退出群聊

**简要描述：** 

- 自己主动推出群聊通知（比如手机端退出或者其他端退出群聊）

**返回类型：** 
- ` 115004 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "op_user_id": 1688855749266556, 
    "flag": 16777216, 
    "receiver": 0, 
    "sender_name": "", 
    "is_room": 1, 
    "server_id": 15318052, 
    "send_time": 1687688387, 
    "sender": 1688855749266556, 
    "referid": 0, 
    "app_info": "QUIT_CC_CAQQwqngpAYYACCNytanDg==", 
    "room_conversation_id": "10927640775842980", 
    "readuinscount": 0, 
    "msg_id": 1014522, 
    "msgtype": 2055
}

```

### 3.5 群名变更通知

**简要描述：** 

- 群名变更通知

**返回类型：** 
- ` 115005 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 83886080, 
    "receiver": 0, 
    "sender_name": "", 
    "is_room": 1, 
    "server_id": 8082398, 
    "room_name": "666366123456789123456",  //变动群名
    "send_time": 1687684854, 
    "operation_user_name": "", 
    "sender": 1688858100400545, 
    "referid": 0, 
    "app_info": "42C0EC26619C41B7A9AFBA2DE2EBCCC7", 
    "room_conversation_id": "10884332696128689", 
    "readuinscount": 0, 
    "msg_id": 1013395, 
    "msgtype": 1001, 
    "operation_user_id": 1688858100400545
}

```

### 3.6 群主转让 等其他通知

**简要描述：** 

- 群主转让通知

**返回类型：** 
- ` 115006 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "op_user_id": 1688855302603660,
    "flag": 16777216,
    "receiver": 0,
    "extraContent": {

    },
    "sender_name": "",
    "is_room": 1,
    "server_id": 7651920,
    "content": "已经成为新的群主",
    "issync": false,
    "send_time": 1747119529,
    "sender": 1688854256622049,
    "referid": 0,
    "app_info": "1CC983A50ABF48D1BCF7273B1939EE94",
    "room_conversation_id": "10882195777349793",
    "readuinscount": 0,
    "msg_id": 1015064,
    "msgtype": 1022
}
```

### 3.7 群管理员变动通知

**简要描述：** 

- 群管理员变动通知

**返回类型：** 
- ` 115007 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 83886080, 
    "receiver": 0, 
    "sender_name": "", 
    "is_room": 1, 
    "server_id": 12984113, 
    "send_time": 1687685690, 
    "admin_vid": [
        1688853788065293, 
        1688853788065294
    ], 
    "sender": 1688853790599424, 
    "referid": 0, 
    "app_info": "CAQQupTgpAYYgNKQ0o6AgAMg1ZPK/gI=", 
    "room_conversation_id": "10696052955012996", 
    "readuinscount": 0, 
    "msg_id": 1129753, 
    "msgtype": 1043
}

```

### 3.8 群解散通知

**简要描述：** 

- 群解散通知

**返回类型：** 
- ` 115008 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
	"flag": 16777216,
	"receiver": 0,
	"sender_name": "",
	"is_room": 1,
	"server_id": 15318223,
	"send_time": 1687707487,
	"sender": 1688855749266556,
	"referid": 0,
	"app_info": "1C55A1242928401A8C28EF3E455EB0F3@foreign.dismiss",
	"room_conversation_id": "10838966791521282",
	"readuinscount": 0,
	"msg_id": 1014863,
	"msgtype": 1023
}

```

### 3.9 群变动通知

**简要描述：** 

- 群变动通知 群成员新增删除

**返回类型：** 
- ` 115009 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 83886080,
    "receiver": 0,
    "sender_name": "",
    "is_room": 1,
    "server_id": 7129903,
    "issync": false,
    "send_time": 1721210833,
    "sender": 1688855587446404,
    "referid": 0,
    "app_info": "yT9Gw89sQPursHO",
    "room_conversation_id": "10767918571905215", //变动的群id 根据这个id去拿群成员列表 去对比
    "readuinscount": 0,
    "msg_id": 1009812,
    "msgtype": 2118
}

```



## 四. 联系人

### 4.1 联系人添加申请通知

**简要描述：** 

- 联系人添加申请通知

**返回类型：** 
- ` 116001 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
 	"issync": false, //是否是因为调用消息同步接口返回的
    "list":[
        {
            "unionid":"ozynqskXWQK3wAWWKT_izAOfcIlY",
            "create_time":0,
            "sex":1,
            "mobile":"",
            "company_remark":"",
            "acctid":"",
            "avatar":"http://wx.qlogo.cn/mmhead/Q3auHgzwzM7BQVtj83EnfQcDMjz2hhrIvssZl9fv1H6rYUaN8VRibjg/0",
            "english_name":"",
            "remark_phone":[

            ],
            "realname":"",
            "real_remarks":"",
            "user_id":7881302555913738,
            "nickname":"十年一刻222222",
            "position":"",
            "corp_id":1970325134026788,
            "remarks":"",
            "seq":8085933,
            "status":0
        }
    ],
    "seq":8085933
}

```

### 4.2 联系人变动通知

**简要描述：** 

- 联系人变动通知

**返回类型：** 
- ` 116002 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 

{
 	"issync": false,//是否是因为调用消息同步接口返回
    "list": [
        {
            "unionid": "ozynqskXWQK3wAWWKT_izAOfcIlY", 
            "create_time": 0, 
            "add_customer_time": 1678428081, 
            "sex": 1, 
            "mobile": "", 
            "company_remark": "测试企业55555", 
            "acctid": "", 
            "avatar": "http://wx.qlogo.cn/mmhead/Q3auHgzwzM7BQVtj83EnfQcDMjz2hhrIv6vzOPdXX4mLpIEJUbE2ew/0", 
            "source": 104,
			"ItemFlag": 1,
            "english_name": "", 
            "remark_phone": [
                "135xxxx"
            ], 
            "realname": "", 
            "real_remarks": "十年一刻66666497657", 
            "labelid": [
                "14073749556770626", 
                "14073751625749478"
            ], 
            "user_id": 7881302555913738, 
            "nickname": "十年一刻", 
            "position": "", 
            "ItemFlag": 0, 
            "corp_id": 1970325134026788, 
            "source_info": {
                "vid": 0, 
                "wx_friend_name": "", 
                "apply_mode": 1, 
                "mobile": "", 
                "source_roomid": 10889703532203248, 
                "source_type": 104
            }, 
            "remarks": "是你身边", 
            "seq": 8077754, 
            "status": 2057 //2049 代表被删除 8也代表被删除  9代表新增通知 2057就代表 备注 标签 电话等信息变动
        }
    ], 
    "seq": 8077754
}

```

### 4.3 内部联系人标签描述变动通知

**简要描述：** 

- 内部联系人标签描述变动通知

**返回类型：** 
- ` 116003 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "list": [
        {
            "vid": 1688856554448765, 
            "description": "描述修改", 
            "remark": "备注修改", 
            "timestamp": 1683273947
        }
    ], 
    "seq": 15312890
}

```

### 4.4 当前联系人已读消息通知回调

**简要描述：** 

- 当前联系人已读消息通知回调，在手机上已读了消息的回调，用来去除自己开发的应用小红点。

**返回类型：** 
- ` 102006 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 0,
    "receiver": 7881302059975581, //已读的好友id
    "sender_name": "",
    "is_room": 0,
    "server_id": 7549931,
    "issync": false,
    "send_time": 1733673509,
    "sender": 1688854256622049,
    "referid": 0,
    "app_info": "CAEQpYTXugYY4bOssJCAgAMgEg==",
    "readuinscount": 0,
    "msg_id": 1007061,
    "msgtype": 2001
}

```

### 4.5 联系人发消息提示删除通知

**简要描述：** 

- 联系人删除通知

**返回类型：** 
- ` 102001 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "msg": "十年一刻开启了联系人验证，你还不是他（她）的联系人。请先发送联系人验证请求，对方验证通过后，才能聊天。 发送联系人验证",
    "flag": 16777216,
    "receiver": 788130255591xxxx,  //好友id
    "sender_name": "",
    "is_room": 0,
    "server_id": 7649730,
    "issync": false,
    "send_time": 1741528470,
    "sender": 16888542566xxxxx, //发送人id
    "referid": 0,
    "app_info": "CIOACBDynL7Z1zIY4bOssJCAgAMgNA==_need_verify",
    "readuinscount": 0,
    "msg_id": 1009857,
    "msgtype": 1011
}

```

### 4.6 联系人发消息提示拉黑通知

**简要描述：** 

- 联系人拉黑通知

**返回类型：** 
- ` 102002 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "msg": "消息已发出，但被对方拒收了",
    "flag": 16777216,
    "receiver": 788130255591xxxx,
    "sender_name": "",
    "is_room": 0,
    "server_id": 7649739,
    "issync": false,
    "send_time": 1741528681,
    "sender": 168885425662xxxx,
    "referid": 0,
    "app_info": "CIOACBDUkMvZ1zIY4bOssJCAgAMgOg==_blacklist",
    "readuinscount": 0,
    "msg_id": 1009877,
    "msgtype": 1011
}

```

### 4.7 置顶免打扰标记操作通知

**简要描述：** 

- 置顶免打扰标记操作通知
出现这个通知代表其他端操作了置顶 免打扰 或者标记操作，
调用一下[获取置顶免打扰标记列表接口](https://www.showdoc.com.cn/1889765213630885/11558938230701720 "获取置顶免打扰标记列表接口")

**返回类型：** 
- ` 116005 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "innerkf_vid": 0,
    "flag": 0,
    "receiver": 0,
    "extraContent": {

    },
    "devInfo": 0,
    "sender_name": "",
    "is_room": 0,
    "server_id": 12574586,
    "kf_id": 0,
    "issync": false,
    "send_time": 1758881937,
    "sender": 10008,
    "referid": 0,
    "app_info": "HP1K_0QkQSulhN1",
    "readuinscount": 0,
    "msg_id": 1005422,
    "msgtype": 2115
}
```

## 五. 控制台下发任务消息

### 5.1 控制台下发群发客户群任务消息

**简要描述：** 

- 控制台下发群发客户群任务消息

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 0,
    "receiver": 0,
    "msgid": "112767170301410738",
    "sender_name": "",
    "is_room": 1,
    "server_id": 7122805,
    "title": "33通知你群发消息给客户群",
    "content": {
        "listdata": [
            {
                "now_cnt": 0,
                "total_cnt": 0,
                "time_stamp": 1720690464,
                "keywords": [

                ],
                "id": 112767170301410738,
                "seq": 12462094,
                "content": [
                    {
                        "atList": [

                        ],
                        "content": "sadfasd"
                    },
                    {
                        "file_id": "306d02010204663xxxx393962353366346336393335653239636566350203103800020300c6d004000201010201000400",
                        "aes_key": "716366626C6B7261746D6B6378796F6B",
                        "width": 358,
                        "is_hd": 1,
                        "file_size": 50884,
                        "md5": "bdc884a5c545699b53f4c6935e29cef5",
                        "height": 457
                    }
                ],
                "msg_send_vids": [

                ]
            }
        ],
        "is_end": false,
        "seq": 0
    },
    "issync": false,
    "send_time": 1720690467,
    "sender": 10167,
    "referid": 0,
    "app_info": "KHGSC_1688855587446404_112767170301410738_n",
    "readuinscount": 0,
    "msg_id": 1009446,
    "msgtype": 573
}

```

### 5.2 控制台群发客户消息任务下发

**简要描述：** 

- 控制台群发客户消息任务下发

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 0, 
    "receiver": 0, 
    "msgid": "109948395468847538", 
    "sender_name": "", 
    "is_room": 0, 
    "server_id": 12883270, 
    "title": "但是发射点通知你群发消息给客户", 
    "send_time": 1677679375, 
    "sender": 10067, 
    "referid": 0, 
    "app_info": "KHGSC_1688853790599424_109948395468847538", 
    "readuinscount": 0, 
    "msg_id": 1115206, 
    "msgtype": 573
}

```

### 5.3 朋友圈控制台下发

**简要描述：** 

- 朋友圈控制台下发

**返回类型：** 
- ` 102000 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "flag": 536870912, 
    "receiver": 0, 
    "sender_name": "", 
    "is_room": 0, 
    "server_id": 12884848, 
    "type": 2, 
    "title": "但是发射点通知你发表内容到客户的朋友圈", 
    "content": "朋友圈内容啊啊啊", 
    "sid": 7217451568562415000, 
    "send_time": 1680443894, 
    "sender": 10086, 
    "author_vid": 1688853790599424, 
    "referid": 0, 
    "app_info": "7217451568562414624_new", //要发送的id 传new_前的数字 自己分割一下
    "readuinscount": 0, 
    "msg_id": 1119290, 
    "msgtype": 529, 
    "seq": 1680443894083
}

```

## 六. 标签变动回调

### 6.1 个人标签新增回调

**简要描述：** 

- 标签新增回调

**返回类型：** 
- ` 130001 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "labellist": [
        {
            "op": 3, 
            "bDeleted": 0, //0代表新增
            "create_time": 1678114162, 
            "label_groupid": 14073749131792038, 
            "label_type": 2, 
            "source_appid": 0, 
            "business_type": 0, 
            "name": "测试回调标签", 
            "data_type": 1, 
            "id": 14073750799769410, 
            "service_groupid": 0, 
            "new_order": 0, 
            "order": 0
        }
    ], 
    "index": 12882694
}

```

### 6.2 个人标签删除回调通知

**简要描述：** 

- 个人标签删除回调通知

**返回类型：** 
- ` 130001 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "labellist": [
        {
            "op": 2, 
            "bDeleted": 1, //1代表删除
            "create_time": 1678114162, 
            "label_groupid": 14073749131792038, 
            "label_type": 2, 
            "source_appid": 0, 
            "business_type": 0, 
            "name": "测试回调标签", 
            "data_type": 1, 
            "id": 14073750799769410, 
            "service_groupid": 0, 
            "new_order": 0, 
            "order": 0
        }
    ], 
    "index": 12882695
}

```

### 6.3 企业标签删除回调

**简要描述：** 

- 企业标签删除回调

**返回类型：** 
- ` 130002 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "labellist": [
        {
            "op": 2, 
            "bDeleted": 1, //1代表删除
            "create_time": 1678112965, 
            "label_groupid": 14073749131791700, 
            "label_type": 1, 
            "source_appid": 0, 
            "business_type": 0, 
            "name": "新增测试回调", 
            "data_type": 1, 
            "id": 14073748971730364, 
            "service_groupid": 0, 
            "new_order": 0, 
            "order": 0
        }
    ], 
    "index": 9684989
}

```

### 6.4 企业标签新增回调

**简要描述：** 

- 企业标签新增回调

**返回类型：** 
- ` 130002 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "labellist": [
        {
            "op": 2, 
            "bDeleted": 0, //0代表新增
            "create_time": 1678112965, 
            "label_groupid": 14073749131791700, 
            "label_type": 1, 
            "source_appid": 0, 
            "business_type": 0, 
            "name": "新增测试回调", 
            "data_type": 1, 
            "id": 14073748971730364, 
            "service_groupid": 0, 
            "new_order": 0, 
            "order": 0
        }
    ], 
    "index": 9684989
}

```

## 七. 聊天标签

### 7.1 聊天标签变动通知

**简要描述：** 

- 聊天标签变动通知

**返回类型：** 
- ` 102004 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "send_time":1699256104,
    "flag":0,
    "receiver":0,
    "sender":10014,
    "referid":0,
    "sender_name":"",
    "app_info":"xbTJPAKbRsi6qyX",
    "readuinscount":0,
    "is_room":0,
    "msg_id":1005018,
    "server_id":5402218,
    "msgtype":2160
}

```

### 7.2 聊天标签中的联系人变动通知

**简要描述：** 

- 聊天标签中的联系人变动通知

**返回类型：** 

- ` 102005 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "send_time":1699256104,
    "flag":0,
    "receiver":0,
    "sender":10014,
    "referid":0,
    "sender_name":"",
    "app_info":"0FUGGXUfQ0G7nPh",
    "readuinscount":0,
    "is_room":0,
    "msg_id":1005020,
    "server_id":5402219,
    "msgtype":2161
}

```

## 八. 断开回调

### 8.1 二次验证消息通知

**简要描述：** 

- 二次验证消息通知

**返回类型：** 
- ` 100012 `

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

**返回示例**
``` 
{
    "vid":1688856591445501,
    "msg":"出现二次验证，请重新获取二维码扫码验证登录",
    "uuid":"680554c2b347acf9b77b73ac72448afa"
}
```

### 8.2 手机端结束登陆


##### 简要描述

- 手机端结束登录

##### 操作码

- 100005

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `


##### 返回示例 

``` 
{
	"vid":16688xxxxxxxxxxxx,
	"msg":"手机端结束登录"
}
```

### 8.3 当前账号在其他设备登陆


##### 简要描述

- 当前账号在其他设备登录

##### 操作码

- 100008


**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

##### 返回示例 

``` 
{
	"vid":16688xxxxxxxxxxxx,
	"msg":"当前账号在其他设备登录，请重新扫码登录"
}
```

##### 备注 

- 更多返回错误代码请看首页的错误代码描述

### 8.4 初始化接口二维码扫码超时


##### 简要描述

- 初始化接口二维码扫码超时断开 （10分钟内没扫码登录会销毁实例 需要重新初始化获取二维码）

##### 操作码

- 100009

**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `


##### 返回示例 

``` 
{
	"vid":16688xxxxxxxxxxxx,
	"msg":"初始化接口二维码扫码超时断开"
}
```

##### 备注 

- 更多返回错误代码请看首页的错误代码描述

### 8.5 异常断开


##### 简要描述

- 异常断开

##### 操作码

- 100007


**返回uuid：** 
- ` f7503bb5-7d27-408f-ab24-8c4ace7f068e `

##### 返回示例 

``` 
{
	"vid":16688xxxxxxxxxxxx,
	"msg":"登陆缓存失效，请重新登录。"
}
```

##### 备注 

- 更多返回错误代码请看首页的错误代码描述

