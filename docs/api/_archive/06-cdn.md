> ← [返回文档索引](../../INDEX.md) | [架构文档](../../ARCHITECTURE.md)
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

