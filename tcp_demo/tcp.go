package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//defer conn.Close()
		//不做边界处理  tcp沾包
		// for {
		// 	var data = make([]byte, 5)
		// 	n, err := conn.Read(data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	if n > 0 {
		// 		log.Println("receive:", n, "bytes:", string(data))
		// 	}
		// }
		//分隔符协议 常用  根据特定"\n"作为边界
		reader := bufio.NewReader(conn)
		for {
			data, err := reader.ReadSlice('\n')
			if err != nil {
				if err != io.EOF {
					log.Println(err)
				} else {
					break
				}
			}
			log.Println("receive:", len(data), "bytes:", string(data))
		}
		//长度协议 根据固定长度作为边界
		// 	reader := bufio.NewReader(conn)
		// 	for {
		// 		peek, err := reader.Peek(4)
		// 		if err != nil {
		// 			if err != io.EOF {
		// 				log.Fatal(err)
		// 			} else {
		// 				break
		// 			}
		// 		}
		// 		buffer := bytes.NewBuffer(peek)
		// 		var length int32
		// 		err = binary.Read(buffer, binary.BigEndian, &length)
		// 		if err != nil {
		// 			log.Println(err)
		// 		}
		// 		if int32(reader.Buffered()) < length+4 {
		// 			continue
		// 		}
		// 		data := make([]byte, length+4)
		// 		_, err = reader.Read(data)
		// 		if err != nil {
		// 			continue
		// 		}
		// 		log.Println("receive msg", string(data[4:]))
		// 	}
	}
}
