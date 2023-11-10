import json
import socket
import requests
rsp=requests.post("http://127.0.0.1:9090/jsonrpc",json={
    "id":0,
    "params":["bobby"],
    "method":"HelloService.Hello"
})
print(rsp.text)