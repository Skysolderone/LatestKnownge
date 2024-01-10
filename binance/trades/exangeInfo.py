import requests
#交易规范信息
api_url="https://api4.binance.com"

resp=requests.get(api_url+'/api/v3/exchangeInfo')
print(resp.json())
