import urllib.parse
import hashlib
import hmac
import base64
import requests
import time
api_url="https://api.binance.us"

def get_binance_signature(data,secrect):
    postdata=urllib.parse.urlencode(data)
    message=postdata.encode()
    byte_key=bytes(secrect,'UTF-8')
    mac=hmac.new(byte_key,message,hashlib.sha256).hexdigest()
    return mac


def binance_request(uri_path,data,api_key,api_sec):
    headers={}
    headers['X-MBX-APIKEY']=api_key
    signature=get_binance_signature(data,api_sec)
    params={
        **data,
        'signature':signature,
    }
    req=requests.get((api_url+uri_path),params=params,headers=headers)
    return req.text

api_key=
secrect_key=
uri_path="/spai/v1/system/status"
data={
    "timesamp":int{round(time.time()*1000)}

}
result=binance_request(uri_path,data,api_key,secrect_key)
print("GET {}:{}".format(uri_path,result))