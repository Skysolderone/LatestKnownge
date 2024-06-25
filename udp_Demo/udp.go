package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

// UDP
// 0x1代表文件信息，客户端传送给服务端；0x2代表服务端确认信息，服务端传送给客户端；0x3代表分块数据，客户端传送给服务端；0x4代表客户端重传指令，服务端传送给客户端；0x5代表文件传输完毕，客户端发送给服务端。
func UdpSlice(file string) {
	var data [1040]byte
	n, addr, err := lis.ReadFromUDP(data[:])
	if err != nil {
		fmt.Println("read udp failed err,", err)
		continue
	}

	fmt.Printf("pkgIdx:%d,read:%d \n", pkgIdx, n)

	// TLV 类型或标签，长度，值

	// ======= 1       | 1    |1    |1     |4     |4     |4     |
	// ======= 0xAA    | 0x1  |x    |x     |seq   |len   |seq     |
	// ======= 前导符  ||命令 ||保留 ||保留 ||包序号 ||长度  ||块序号 ||
	cmd := data[0]
	if cmd == 0xAA { // b:1010 1010
		// 命令消息
		cmdVal := data[1]
		switch cmdVal {
		case 0x1: // 客户端文件信息 [文件大小|缓存大小|文件类型|分片数量]
			if n != 32 {
				fmt.Println("0x1格式不对")
				continue
			}
			fmt.Println("收到0x1消息")
			fileSize = utils.Bytes2Int(data[16:20])
			cacheSize := utils.Bytes2Int(data[20:24])
			ft := utils.Bytes2Int(data[24:28])
			fileType = getSufByType(ft)
			pieceCount := utils.Bytes2Int(data[28:])
			fmt.Println("cacheSize,", cacheSize, ",pieceCount,", pieceCount)
			fileCache = make([]byte, cacheSize)
			fileCacheIdx = 0

			pkgIdx += 1
			b2 := bytes.NewBuffer([]byte{})
			binary.Write(b2, binary.LittleEndian, uint8(0xaa))
			binary.Write(b2, binary.LittleEndian, uint8(0x2))
			binary.Write(b2, binary.LittleEndian, uint8(0x0))
			binary.Write(b2, binary.LittleEndian, uint8(0x0))
			binary.Write(b2, binary.LittleEndian, uint32(pkgIdx))
			binary.Write(b2, binary.LittleEndian, uint32(0))
			binary.Write(b2, binary.LittleEndian, uint32(0))

			_, err = lis.WriteToUDP(b2.Bytes(), addr)
			if err != nil {
				fmt.Println("write to udp failed err,", err)
				continue
			}

		case 0x2: // 服务端收到回执 []
		case 0x3: // 客户端文件分块数据 [分块数据]
			pieceData := data[16:n]
			fmt.Printf("test3,%x\n", pieceData)
			copy(fileCache[fileCacheIdx:], pieceData)
			fileCacheIdx = fileCacheIdx + len(pieceData)

			pkgIdx += 1
			b2 := bytes.NewBuffer([]byte{})
			binary.Write(b2, binary.LittleEndian, uint8(0xaa))
			binary.Write(b2, binary.LittleEndian, uint8(0x2))
			binary.Write(b2, binary.LittleEndian, uint8(0x0))
			binary.Write(b2, binary.LittleEndian, uint8(0x0))
			binary.Write(b2, binary.LittleEndian, uint32(pkgIdx))
			binary.Write(b2, binary.LittleEndian, uint32(0))
			binary.Write(b2, binary.LittleEndian, uint32(0))

			_, err = lis.WriteToUDP(b2.Bytes(), addr)
			if err != nil {
				fmt.Println("write to udp failed err,", err)
				continue
			}

		case 0x4: // 服务端要求重传 [块序号]
		case 0x5: // 客户端文件传输完毕
			fmt.Println("file size,", len(fileCache), fileSize)

			tempFileName := fmt.Sprintf("./t%d%s", time.Now().Unix(), fileType)
			wdata := fileCache[:fileSize]
			os.WriteFile(tempFileName, wdata, 0o666)
			fmt.Printf("test1,%x\n", wdata)
			fmt.Printf("test2,%x\n", fileCache)

		}
	}
}

func main() {
	// udp slice
}
