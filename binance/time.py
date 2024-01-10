import requests
import time
t=time.time()
print(t)
resp=requests.get("https://api.binance.us/api/v3/time")
print(resp.json())