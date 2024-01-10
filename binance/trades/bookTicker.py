import requests
#当前最优挂单
api_url='https://api4.binance.com/api/v3'

resp=requests.get(api_url+'/ticker/bookTicker?symbol=BTCUSDT')
print(resp.json())