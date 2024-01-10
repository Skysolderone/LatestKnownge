import urllib.parse
import hashlib
import hmac
import base64
import requests
import time
#账户状态
api_url="https://api4.binance.com"

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

api_key='pYVS6x9Ssoxo1MxJBkwBfA8Mxi4EnJvQPQLYfq9pksp5F7xJHCCe2HqUhjff5Pj8'
secrect_key='YpiJcLTr5cN1fUH11EJkf8DkkIKWpzYX5VvA3kvBIQDuRcnOhQCBmeyEFj1X0CRZ'
uri_path="/sapi/v1/account/status"
data={
    "timestamp":int(round(time.time()*1000)),
}
result=binance_request(uri_path,data,api_key,secrect_key)
print("GET {}:{}".format(uri_path,result))