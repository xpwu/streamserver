# streamserver
基于[go-stream](https://github.com/xpwu/go-stream) 的完整服务，
支持两种应用服务器的推送协议：1、http协议的推送；2、push协议的推送

## 运行 
xxx  -h 参数可以看到帮助信息    
按照帮助提示运行即可

## http push 协议 [demo](https://github.com/xpwu/streamserver/blob/master/test/push.http)    
```
response:   
{
  st: 0
}   
```
st的取值：
```
  0: Success
  1: HostNameErr // 服务不匹配
  2: TokenNotExist
  3: ServerInternalErr
  4: AckTimeout  
```

## client test
go-stream client 测试的 api 接口
* 跑 client 测试用例可用的配置
```json
{
  "github.com/xpwu/go-log/log:config": {
    "// level": "0:DEBUG; 1:INFO; 2:WARNING; 3:ERROR; 4:FATAL",
    "level": 0
  },
  "github.com/xpwu/go-stream/lencontent:config": {
    "Servers": [
      {
        "Net": {
          "// Listen": "1、xxx.xxx.xxx.xxx:[0-9] 2、:[0-9] 3、pipe:[0-9] 4、unix:|xxx|xxx|xxx|xxx.socket:0",
          "Listen": ":8080",
          "MaxConnections": -1,
          "TLS": false,
          "TlsFile": {
            "// PrivateKeyPEMFile": "support relative path, must PEM encode data",
            "PrivateKeyPEMFile": "",
            "// CertPEMFile": "support relative path, must PEM encode data",
            "CertPEMFile": ""
          }
        },
        "// HeartBeat_s": "unit:s",
        "HeartBeat_s": 30,
        "FrameTimeout_s": 5,
        "Proxy": {
          "Url": "http://127.0.0.1:7001/clienttest/${fhttp_api}",
          "Headers": [
            {
              "Key": "pushtoken",
              "Value": "${pushtoken}"
            }
          ]
        }
      }
    ]
  },
  "github.com/xpwu/go-stream/push:config": {
    "Servers": [
      {
        "Net": {
          "// Listen": "1、xxx.xxx.xxx.xxx:[0-9] 2、:[0-9] 3、pipe:[0-9] 4、unix:|xxx|xxx|xxx|xxx.socket:0",
          "Listen": ":7999",
          "MaxConnections": -1,
          "TLS": false,
          "TlsFile": {
            "// PrivateKeyPEMFile": "support relative path, must PEM encode data",
            "PrivateKeyPEMFile": "",
            "// CertPEMFile": "support relative path, must PEM encode data",
            "CertPEMFile": ""
          }
        },
        "// AckTimeout": "uint:s;0: no ack",
        "AckTimeout": 30
      }
    ]
  },
  "github.com/xpwu/go-stream/websocket:config": {
    "Servers": [
      {
        "Net": {
          "// Listen": "1、xxx.xxx.xxx.xxx:[0-9] 2、:[0-9] 3、pipe:[0-9] 4、unix:|xxx|xxx|xxx|xxx.socket:0",
          "Listen": ":8001",
          "MaxConnections": -1,
          "TLS": false,
          "TlsFile": {
            "// PrivateKeyPEMFile": "support relative path, must PEM encode data",
            "PrivateKeyPEMFile": "",
            "// CertPEMFile": "support relative path, must PEM encode data",
            "CertPEMFile": ""
          }
        },
        "// HeartBeat_s": "unit:s",
        "HeartBeat_s": 30,
        "FrameTimeout_s": 5,
        "MaxBytesPerFrame": 100000,
        "MaxConcurrentPerConnection": 10,
        "Origin": [
          "*"
        ],
        "Proxy": {
          "Url": "http://127.0.0.1:7001/clienttest/${fhttp_api}",
          "Headers": [
            {
              "Key": "Pushurl",
              "Value": "127.0.0.1:7999/${pushtoken}"
            }
          ]
        }
      }
    ]
  },
  "github.com/xpwu/go-tinyserver/http:serverConfig": {
    "Net": {
      "// Listen": "1、xxx.xxx.xxx.xxx:[0-9] 2、:[0-9] 3、pipe:[0-9] 4、unix:|xxx|xxx|xxx|xxx.socket:0",
      "Listen": ":7001",
      "// MaxConnections": "-1:not limit",
      "MaxConnections": -1,
      "TLS": false,
      "TlsFile": {
        "// PrivateKeyPEMFile": "support relative path, must PEM encode data",
        "PrivateKeyPEMFile": "",
        "// CertPEMFile": "support relative path, must PEM encode data",
        "CertPEMFile": ""
      }
    },
    "// HostName": "leftmost match, []: allow all host name",
    "HostName": [],
    "// RootUri": "match_uri = RootUri + api.RegisterUri",
    "RootUri": "/"
  }
}
```