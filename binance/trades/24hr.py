import requests

#24hr 价格变化
api_url='https://api4.binance.com/api/v3'
resp=requests.get(api_url+'/ticker/24hr?symbol=BTCUSDT')
print(resp.json())