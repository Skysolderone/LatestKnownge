import urllib.parse
import hashlib
import hmac
import requests
import time
data={
    "timestamp":int(round(time.time()*1000)),
}
def get_binance_signature(data,secrect):
    postdata=urllib.parse.urlencode(data)
    message=postdata.encode()
    byte_key=bytes(secrect,'UTF-8')
    mac=hmac.new(byte_key,message,hashlib.sha256).hexdigest()
    return mac

api_key='pYVS6x9Ssoxo1MxJBkwBfA8Mxi4EnJvQPQLYfq9pksp5F7xJHCCe2HqUhjff5Pj8'
api_url='https://api1.binance.com/fapi/v3/listenKey'
secrect_key='YpiJcLTr5cN1fUH11EJkf8DkkIKWpzYX5VvA3kvBIQDuRcnOhQCBmeyEFj1X0CRZ'

headers={}
headers['X-MBX-APIKEY']=api_key
signature=get_binance_signature(data,secrect_key)
params={
        **data,
        'signature':signature,
    }
req=requests.post(api_url,params=params,headers=headers)
print("{}".format(req.text))