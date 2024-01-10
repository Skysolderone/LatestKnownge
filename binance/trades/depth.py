import requests


#深度信息
api_url='https://api4.binance.com/api/v3'

resp=requests.get(api_url+'/depth?symbol=LTCBTC&limit=1')
print(resp.json())