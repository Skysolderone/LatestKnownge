import requests

api_url="https://api4.binance.com"
resp=requests.get(api_url+'/api/v3/ticker/price?symbol=LTCBTC')
print(resp.json())

resp2=requests.get(api_url+'/api/v3/ticker/price')
print(resp2.json())