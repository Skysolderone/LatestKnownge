import requests

#最新价格
api_url='https://api4.binance.com/api/v3'
resp=requests.get(api_url+'/ticker/price?symbol=BTCUSDT')
print(resp.json())