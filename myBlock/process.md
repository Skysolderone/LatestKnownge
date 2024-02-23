geth init genesis.json
geth --datadir .\ init .\genesis.json 指定目录创建私有链 如使用上述命令即在默认目录创建 会有冲突

geth --datadir .\ 启动节点
geth attach ipc::\\.\pipe\geth.ipc(geth attach \\.\pipe\geth.ipc) 连接节点 windows 下需使用 url linux 则会生成在 datadir 目录生成 ipc 文件 在连接

0x135c8a4e0c7fd952d0403dd41b6cc31583d33d63 wws
0xfb9332dc4ee3d64af9ec1d21ea1bed4e4aa2d358 gg123456
