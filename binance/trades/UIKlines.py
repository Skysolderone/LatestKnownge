import requests
#uik线数据
api_url='https://api4.binance.com/api/v3'

resp=requests.get(api_url+'/uiKlines?symbol=LTCBTC&interval=1m&limit=1')
print(resp.json())