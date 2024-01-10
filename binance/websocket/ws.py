import websocket

def on_open(ws):
    print("连接成功")

def on_message(ws, message):
    print("收到消息：", message)

def on_error(ws, error):
    print("WebSocket错误：", error)

def on_close(ws):
    print("连接关闭")

if __name__ == "__main__":
    # 币安WebSocket api地址，根据实际情况修改
    ws_url = "wss://stream.binance.com:9443/ws/btcusdt@trade"
    websocket.enableTrace(True)
    ws = websocket.WebSocketApp(ws_url,
        on_open=on_open,
        on_message=on_message,
        on_error=on_error,
        on_close=on_close)
    ws.run_forever()