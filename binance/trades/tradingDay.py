import requests

#交易日行情
api_url='https://api4.binance.com/api/v3'

resp=requests.get(api_url+'/ticker/tradingDay?symbol=BTCUSDT')
print(resp.json())