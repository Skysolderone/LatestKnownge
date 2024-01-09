import requests

resp=requests.get("https://api.binance.us/api/v3/time")
print(resp.json())