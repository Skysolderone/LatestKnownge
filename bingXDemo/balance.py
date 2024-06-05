
import time
import requests
import hmac
from hashlib import sha256

APIURL = "https://open-api.bingx.com"

APIKEY = "X9rdbwehnjwb3FNTKbCrmAtKUYu6cnp8lltamaLDkbsQiwmprZagNkSzx7mSQiyVmLoaEPLEjPkJGzKUj87QVg"
SECRETKEY = "enBARedV8QiRPoYrlgqOmftOF2pyFoBqtXNLYOx9VPYhRkllql81ctRKALjUCaMzKVIXOG2iFAQ6acYtt1WQ"
# APIKEY = "hO6oQotzTE0S5FRYze2Jx2wGx7eVnJGMolpA1nZyehsoMgCcgKNWQHd4QgTFZuwl4Zt4xMe2PqGBegWXO4A"
# SECRETKEY = "mheO6dR8ovSsxZQCOYEFCtelpuxcWGTfHw7te326y6jOwq5WpvFQ9JNljoTwBXZGv5It07m9RXSPpDQEK2w"

def demo():
    payload = {}
    path = '/openApi/spot/v1/account/balance'
    method = "GET"
    paramsMap = {
    "recvWindow": "60000",
    # "timestamp": "1702624167523"
}
    paramsStr = parseParam(paramsMap)
    return send_request(method, path, paramsStr, payload)

def get_sign(api_secret, payload):
    signature = hmac.new(api_secret.encode("utf-8"), payload.encode("utf-8"), digestmod=sha256).hexdigest()
    print("sign=" + signature)
    return signature


def send_request(method, path, urlpa, payload):
    print(urlpa)
    url = "%s%s?%s&signature=%s" % (APIURL, path, urlpa, get_sign(SECRETKEY, urlpa))
    print(url)
    headers = {
        'X-BX-APIKEY': APIKEY,
    }
    response = requests.request(method, url, headers=headers, data=payload)
    return response.text

def parseParam(paramsMap):
    sortedKeys = sorted(paramsMap)
    paramsStr = "&".join(["%s=%s" % (x, paramsMap[x]) for x in sortedKeys])
    if paramsStr != "": 
     return paramsStr+"&timestamp="+str(int(time.time() * 1000))
    else:
     return paramsStr+"timestamp="+str(int(time.time() * 1000))


if __name__ == '__main__':
    print("demo:", demo())
