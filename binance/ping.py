import requests 
resp=requests.get('https://api.binance.us/api/v3/ping')
print(resp.json())