# TF100设备协议插件

**测试数据**：
```
CCMD:137{"msgType": 20,"devId": 13900001,"msgId": 5,"mainVer": 101,"powVer": [101, 101],"ports": 2,"vsId": 0,"devType": 1,"peType": 0,"enUpd": 1}
```

在平台的网络组件中创建tcp服务，基中的协议选择tf100，端口号自定义，启动服务。

**注册码包**：
```
"devId":\s*([^,]*)
```

创建产品，并配置物模型，创建设备，将设备与产品绑定，启动测试的tcp客户，设备上线，设备数据将会上报到平台。

**物模型配置**

| 字段名   | 描述   | 类型   | 是否只读 |
|----------|--------|--------|---------|
| msgType  | 消息类型 | int  | 只读  |
| msgId    | 消息ID | int  | 只读  |
| ports    | 端口数量 | int  | 只读  |
| errType  | 故障类型 | array | 只读  |