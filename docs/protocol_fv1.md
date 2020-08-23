[TOC]

# FV1 协议

## 日志上报 - /api/fv1/log

### 上报日志信息

Method: POST

请求
```json
{
    "id": "123456",
    "loginfo": {
        "type": "system",
        "time": 2325251,
        "log": "this is a log"
    }
}
```

响应
```json
{
    "result": 0,
    "description": ""
}
```

