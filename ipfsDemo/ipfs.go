package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
)

func main() {
	cid := "Qmf5MjFENk2ArQgV7iqyGK2e5mM4QKvjRJCJrFktPbgj5t"
	node, err := rpc.NewPathApi(cid)
	if err != nil {
		fmt.Println("path", err)
		return
	}
	ctx := context.Background()

	p, err := path.NewPath(cid)
	if err != nil {
		log.Println("new ", err)
		return
	}
	res, err := node.Block().Get(ctx, p)
	file, err := os.Create("ipfsfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = io.Copy(file, res)
	if err != nil {
		fmt.Println(err)
		return
	}
}
