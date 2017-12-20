短网址服务
===

短网址服务

### 启动服务

```
```


### 接口

#### 生成单个短网址

```
GET /v1/shorten?long_url=

{
    "data": "http://u8x.cc/2W5S",
    "err": "",
    "code": 0,
}
```

#### 短链接转长链接

```
GET /v1/expand?short_url=

{
    "code": 0,
    "data": "http://yuez.me",
    "err_msg": ""
}
```

