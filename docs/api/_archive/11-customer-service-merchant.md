> ← [返回文档索引](../../INDEX.md) | [架构文档](../../../ARCHITECTURE.md)
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

