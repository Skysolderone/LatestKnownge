import requests

#滚动窗口价格变动统计
api_url='https://api4.binance.com/api/v3'
resp=requests.get(api_url+'/ticker?symbol=BTCUSDT')
print(resp.json())