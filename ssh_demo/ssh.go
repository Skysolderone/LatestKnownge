package main

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	//basic
	config := &ssh.ClientConfig{
		User: "wws",
		Auth: []ssh.AuthMethod{
			ssh.Password("gg123456"),//需改为服务器用户名跟密码
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //目前demo忽略密钥
	}
	client, err := ssh.Dial("tcp", "192.168.10.128:22", config)
	if err != nil {
		log.Println(err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//模仿终端
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     //禁止回显
		ssh.TTY_OP_ISPEED: 14400, //输入输出速率为14.4k
		ssh.TTY_OP_OSPEED: 14400,
	}
	if err := session.RequestPty("linux", 80, 40, modes); err != nil {
		log.Fatal(err)
	}
	//设置输入输出
	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr
	//启动远程的shell
	if err := session.Shell(); err != nil {
		log.Fatal(err)
	}
	//wait
	err = session.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
