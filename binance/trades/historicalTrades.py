import requests

#历史成交
api_url='https://api4.binance.com/api/v3'
headers={}
headers['X-MBX-APIKEY']='pYVS6x9Ssoxo1MxJBkwBfA8Mxi4EnJvQPQLYfq9pksp5F7xJHCCe2HqUhjff5Pj8'
resp=requests.get(api_url+'/historicalTrades?symbol=LTCUSDT',headers=headers)
print(resp.json())