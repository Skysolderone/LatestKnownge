import requests 
resp=requests.get('https://api4.binance.com/api/v3/ping')
print(resp.json())