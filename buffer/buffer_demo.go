package buffer

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

//读取大文件  30gb

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	const MaxReadBuffer = 64 * 1024 * 1024
	buf := make([]byte, MaxReadBuffer)
	//one
	for {
		bytesRead, err := file.Read(buf)
		if err == io.EOF {
			// 文件读取完毕
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// 处理读取的数据，例如输出到控制台
		fmt.Print(string(buffer[:bytesRead]))
	}
	//two
	scanner := bufio.NewScanner(file)
	scanner.Buffer(buf, MaxReadBuffer)
	for scanner.Scan() {
		line := scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
