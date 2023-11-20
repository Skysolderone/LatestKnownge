package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// 需要构建二进制文件  然后进行test
func hello(cmd *cobra.Command, args []string) {
	fmt.Println("hello world")
}
func main() {

	var rootCmd = &cobra.Command{Use: "wws"}
	var helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "'hello world'",
		Run:   hello,
	}
	rootCmd.AddCommand(helloCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
