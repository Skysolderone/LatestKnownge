import requests

#近期成交
api_url="https://api4.binance.com"
resp=requests.get(api_url+'/api/v3/trades?symbol=LTCBTC')
print(resp.json())