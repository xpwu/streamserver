# streamserver
基于[go-stream](https://github.com/xpwu/go-stream) 的完整服务，
支持两种应用服务器的推送协议：1、http协议的推送；2、push协议的推送


## http 协议 [demo](https://github.com/xpwu/streamserver/blob/master/test/push.http)    
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