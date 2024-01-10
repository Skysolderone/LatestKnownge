import requests

#近期成交（归集）
api_url='https://api4.binance.com/api/v3'
resp=requests.get(api_url+'/aggTrades?symbol=LTCBTC&limit=1')
print(resp.json())

