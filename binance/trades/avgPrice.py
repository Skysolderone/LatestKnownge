import requests
#当前平均价格
api_url='https://api4.binance.com/api/v3'

resp=requests.get(api_url+'/avgPrice?symbol=LTCBTC')
print(resp.json())