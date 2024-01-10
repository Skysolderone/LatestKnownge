import requests

#K线数据
api_url='https://api4.binance.com/api/v3'

resp=requests.get(api_url+'/klines?symbol=LTCBTC&interval=15m')
print(resp.json())