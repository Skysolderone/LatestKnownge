package main

import (
	"encoding/hex"
	"fmt"

	"github.com/multiformats/go-multihash"
)

func main() {
	// 将字符串转换为字节数组
	buf, _ := hex.DecodeString("0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33")
	// 将字节数组用 multihash 编码
	mHashBuf, _ := multihash.EncodeName(buf, "sha1")
	// 打印编码之后的摘要
	fmt.Printf("hex: %s\n", hex.EncodeToString(mHashBuf))
	// 将 binary multihash 转换为 DecodedMultihash 的形式
	mHash, _ := multihash.Decode(mHashBuf)
	// 获取数字摘要
	sha1hex := hex.EncodeToString(mHash.Digest)
	// 打印
	fmt.Printf("obj: %v 0x%x %d %s\n", mHash.Name, mHash.Code, mHash.Length, sha1hex)
}
